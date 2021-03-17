package context

import (
	"polygnosics-frontend/restfrontend"
)

type Context struct {
	RESTFrontend *restfrontend.RESTFrontend
}

func NewContext() (*Context, error) {
	context := &Context{
		RESTFrontend: restfrontend.NewRESTController(),
	}

	return context, nil
}
