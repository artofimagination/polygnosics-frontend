package context

import (
	"github.com/artofimagination/polygnosics-frontend/initialization"
	"github.com/artofimagination/polygnosics-frontend/rest/backend"
	"github.com/artofimagination/polygnosics-frontend/rest/frontend"
	"github.com/artofimagination/polygnosics-frontend/rest/ipresolver"
)

type Context struct {
	RESTFrontend *frontend.RESTController
	RESTBackend  *backend.RESTController
	IPResolver   *ipresolver.IPResolver
	Config       *initialization.Config
}

func NewContext() (*Context, error) {
	cfg := &initialization.Config{}
	initialization.InitConfig(cfg)

	backend := &backend.RESTController{}
	ipresolver := ipresolver.NewIPResolver(backend, cfg)
	if err := ipresolver.DetectValidAddresses(); err != nil {
		return nil, err
	}

	context := &Context{
		RESTFrontend: frontend.NewRESTController(backend),
		RESTBackend:  backend,
		IPResolver:   ipresolver,
		Config:       cfg,
	}

	return context, nil
}
