package ipresolver

import (
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

func NewIPResolver(cfg *initialization.Config) *IPResolver {
	backendServer := &rest.Server{
		IP:   cfg.BackendAddress,
		Port: cfg.BackendPort,
		Name: cfg.BackendName,
	}

	ipresolver := &IPResolver{
		BackendAddress: backendServer,
	}

	return ipresolver
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
	requestData, err := r.DecodeRequest()
	if err != nil {
		c.HandleError(w, fmt.Sprintf("Failed to get root user. %s", errors.WithStack(err)), http.StatusInternalServerError, c.URI(IndexPage))
		w.WriteError(fmt.Sprintf("Frontend -> %s", err.Error()), http.StatusBadRequest)
		return
	}

	c.BackendAddress.IP = requestData["ip"].(string)
	c.BackendAddress.Port = requestData["port"].(int)
	c.BackendAddress.Name = requestData["name"].(string)

	w.WriteData("OK", http.StatusCreated)
}

func (c *IPResolver) AddRouting(r *mux.Router) {
	r.HandleFunc("/set-backend-address", c.setBackendAddress)
}
