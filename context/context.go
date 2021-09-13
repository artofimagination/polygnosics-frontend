package context

import "github.com/artofimagination/polygnosics-frontend/rest/frontend"

type Context struct {
	RESTFrontend *frontend.RESTController
}

func NewContext() (*Context, error) {
	context := &Context{
		RESTFrontend: frontend.NewRESTController(),
	}

	return context, nil
}
