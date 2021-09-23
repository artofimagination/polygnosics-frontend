package backend

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/artofimagination/polygnosics-frontend/rest"
)

var StatsServerAddress string = "http://172.18.0.6:8086"

type RESTController struct {
	BackendAddress *rest.Server
}

type ResponseData struct {
	Error string      `json:"error" validation:"required"`
	Data  interface{} `json:"data" validation:"required"`
}

func (c *RESTController) Post(path string, parameters interface{}) error {
	return post(c.BackendAddress.GetAddress(), path, parameters)
}

func (c *RESTController) Get(path string, parameters string) (interface{}, error) {
	return get(c.BackendAddress.GetAddress(), path, parameters)
}

func (c *RESTController) ForwardRequest(r *http.Request) (interface{}, error) {
	return forwardRequest(c.BackendAddress.GetAddress(), r)
}

func forwardRequest(address string, r *http.Request) (interface{}, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	r.Body = ioutil.NopCloser(bytes.NewReader(body))
	proxyReq, err := http.NewRequest(r.Method, fmt.Sprintf("%s%s", address, r.RequestURI), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

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
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	dataMap := &ResponseData{}
	if err := json.Unmarshal(respBody, &dataMap); err != nil {
		return nil, err
	}

	if dataMap.Error != "" {
		return nil, errors.New(dataMap.Error)
	}

	return dataMap.Data, nil
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

	dataMap := &ResponseData{}
	if err := json.Unmarshal(body, &dataMap); err != nil {
		return nil, err
	}

	if dataMap.Error != "" {
		return nil, errors.New(dataMap.Error)
	}

	return dataMap.Data, nil
}

func post(address string, path string, parameters interface{}) error {
	reqBody, err := json.Marshal(parameters)
	if err != nil {
		return err
	}

	resp, err := http.Post(fmt.Sprintf("%s%s", address, path), "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	dataMap := &ResponseData{}
	if err := json.Unmarshal(body, &dataMap); err != nil {
		return err
	}

	if dataMap.Error != "" {
		return errors.New(dataMap.Error)
	}

	return nil
}
