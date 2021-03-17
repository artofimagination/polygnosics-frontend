package restbackend

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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

func forwardRequest(address string, _ string, r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	r.Body = ioutil.NopCloser(bytes.NewReader(body))
	proxyReq, err := http.NewRequest(r.Method, fmt.Sprintf("%s%s", address, r.RequestURI), bytes.NewReader(body))
	if err != nil {
		return err
	}

	for header, values := range r.Header {
		for _, value := range values {
			proxyReq.Header.Add(header, value)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(proxyReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	dataMap := make(map[string]interface{})
	if err := json.Unmarshal(respBody, &dataMap); err != nil {
		return err
	}

	if val, ok := dataMap["error"]; ok {
		return errors.New(val.(string))
	}

	return errors.New("Invalid response")
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
