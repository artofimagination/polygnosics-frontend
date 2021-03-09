package webrtc

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/pion/webrtc/v3"
)

// Decode decodes the input from base64
// It can optionally unzip the input after decoding
func decode(in string, obj interface{}) {
	b, err := base64.StdEncoding.DecodeString(in)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(b, obj)
	if err != nil {
		panic(err)
	}
}

// Encode encodes the input in base64
// It can optionally zip the input before encoding
func encode(obj interface{}) string {
	b, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(b)
}

func SetupFrontend(w http.ResponseWriter, r *http.Request, offerStr string, dataProvider func() ([]byte, error)) error {
	// Prepare the configuration
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}

	// Create a new RTCPeerConnection
	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		return err
	}

	// Set the handler for ICE connection state
	// This will notify you when the peer has connected/disconnected
	peerConnection.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
		log.Printf("ICE Connection State has changed: %s\n", connectionState.String())
	})

	// Register data channel creation handling
	peerConnection.OnDataChannel(func(d *webrtc.DataChannel) {
		log.Printf("New DataChannel %s %d\n", d.Label(), d.ID())

		// Register channel opening handling
		d.OnOpen(func() {
			log.Printf("Data channel '%s'-'%d' open.\n", d.Label(), d.ID())

			for range time.NewTicker(2 * time.Second).C {
				jsonData, err := dataProvider()
				if err != nil {
					log.Println(err)
					peerConnection.Close()
					break
				}
				if err := d.SendText(string(jsonData)); err != nil {
					log.Println(err)
					peerConnection.Close()
					break
				}
			}
		})

		// Register text message handling
		d.OnMessage(func(msg webrtc.DataChannelMessage) {
			log.Printf("Message from DataChannel '%s': '%s'\n", d.Label(), string(msg.Data))
		})
	})

	// Wait for the offer to be pasted
	offer := webrtc.SessionDescription{}
	decode(offerStr, &offer)

	// Set the remote SessionDescription
	err = peerConnection.SetRemoteDescription(offer)
	if err != nil {
		peerConnection.Close()
		return err
	}

	// Create an answer
	answer, err := peerConnection.CreateAnswer(nil)
	if err != nil {
		peerConnection.Close()
		return err
	}

	gatherComplete := webrtc.GatheringCompletePromise(peerConnection)

	// Sets the LocalDescription, and starts our UDP listeners
	err = peerConnection.SetLocalDescription(answer)
	if err != nil {
		peerConnection.Close()
		return err
	}

	<-gatherComplete

	_, err = w.Write([]byte(encode(*peerConnection.LocalDescription())))
	if err != nil {
		peerConnection.Close()
		return err
	}
	return nil
}
