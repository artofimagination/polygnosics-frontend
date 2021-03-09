package restbackend

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	GET  = "GET"
	POST = "POST"
)

var BusinessLogicServerAddress string = "0.0.0.0:8081"
var StatsServerAddress string = "0.0.0.0:8082"

type RESTBackend struct {
}

func Get(address string, path string, parameters string) ([]byte, error) {
	//request := fmt.Sprintf("%s%s?%s", address, path, parameters)
	resp, err := http.Get(fmt.Sprintf("%s%s?%s", address, path, parameters))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
