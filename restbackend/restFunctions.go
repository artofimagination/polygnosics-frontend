package restbackend

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	GET  = "GET"
	POST = "POST"
)

var BusinessLogicServerAddress string = "http://172.18.0.2:8082"
var StatsServerAddress string = "http://172.18.0.2:8083"

type RESTBackend struct {
}

func forwardRequest(address string, path string, r *http.Request) (interface{}, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	r.Body = ioutil.NopCloser(bytes.NewReader(body))
	proxyReq, err := http.NewRequest(r.Method, fmt.Sprintf("%s%s", address, path), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	// proxyReq.Header.Set("Host", r.Host)
	// proxyReq.Header.Set("X-Forwarded-For", r.RemoteAddr)

	for header, values := range r.Header {
		for _, value := range values {
			proxyReq.Header.Add(header, value)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(proxyReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	log.Println(resp)
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Println(respBody)

	dataMap := make(map[string]interface{})
	if err := json.Unmarshal(respBody, &dataMap); err != nil {
		return nil, err
	}

	log.Println(dataMap)

	if val, ok := dataMap["error"]; ok {
		return nil, errors.New(val.(string))
	}

	if val, ok := dataMap["data"]; ok {
		return val, nil
	}

	return nil, errors.New("Invalid response")
}

func get(address string, path string, parameters string) (interface{}, error) {
	resp, err := http.Get(fmt.Sprintf("%s%s%s", address, path, parameters))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	dataMap := make(map[string]interface{})
	if err := json.Unmarshal(body, &dataMap); err != nil {
		return nil, err
	}

	if val, ok := dataMap["error"]; ok {
		return nil, errors.New(val.(string))
	}

	if val, ok := dataMap["data"]; ok {
		return val, nil
	}

	return nil, errors.New("Invalid response")
}

func post(address string, path string, parameters map[string]interface{}) (interface{}, error) {
	reqBody, err := json.Marshal(parameters)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(fmt.Sprintf("%s%s", address, path), "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	dataMap := make(map[string]interface{})
	if err := json.Unmarshal(body, &dataMap); err != nil {
		return nil, err
	}

	if val, ok := dataMap["error"]; ok {
		return nil, errors.New(val.(string))
	}

	if val, ok := dataMap["data"]; ok {
		return val, nil
	}

	return nil, errors.New("Invalid response")
}
