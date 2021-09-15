package ipresolver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/artofimagination/polygnosics-frontend/initialization"
	"github.com/artofimagination/polygnosics-frontend/rest"
	"github.com/artofimagination/polygnosics-frontend/rest/backend"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

type IPResolver struct {
	BackendController *backend.RESTController
	BackendAddress    *rest.Server
}

func NewIPResolver(backend *backend.RESTController, cfg *initialization.Config) *IPResolver {
	backendServer := &rest.Server{
		IP:   cfg.BackendAddress,
		Port: cfg.BackendPort,
		Name: cfg.BackendName,
	}

	prettyPrint(backendServer)

	backend.BackendAddress = backendServer

	ipresolver := &IPResolver{
		BackendAddress: backendServer,
	}

	return ipresolver
}

func prettyPrint(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
		return
	}
	fmt.Println("Failed to pretty print data")
}

// DetectValidAddresses waits 5 seconds to allow the IP resolver server to set the user and resource db addresses
// Only waits if the address was not set through env vars before.
// Returns error if the address is not set either by env. vars or the IP resolver server.
func (c *IPResolver) DetectValidAddresses() error {
	backendSet := false
	for retryCount := 5; retryCount > 0; retryCount-- {
		if c.BackendAddress.Name != "Unknown" {
			backendSet = true
			break
		}
		time.Sleep(1 * time.Second)
		log.Printf("No valid backend address")
	}

	if !backendSet {
		return errors.New("No valid backend address detected")
	}

	return nil
}

func (c *IPResolver) setBackendAddress(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		log.Fatalf("Failed to get backend address. %s\n", errors.WithStack(err))
		return
	}

	c.BackendAddress.IP = data["ip"].(string)
	c.BackendAddress.Port = data["port"].(int)
	c.BackendAddress.Name = data["name"].(string)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "OK")
}

func (c *IPResolver) AddRouting(r *mux.Router) {
	r.HandleFunc("/set-backend-address", c.setBackendAddress)
}
