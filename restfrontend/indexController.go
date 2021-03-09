package restfrontend

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

func (c *RESTFrontend) Contact(w http.ResponseWriter, r *http.Request) {
	content := c.ContentController.BuildContactContent()
	c.RenderTemplate(w, AboutContact, content)
}

func (c *RESTFrontend) Career(w http.ResponseWriter, r *http.Request) {
	content := make(map[string]interface{})
	c.RenderTemplate(w, AboutCareer, content)
}

func (c *RESTFrontend) About(w http.ResponseWriter, r *http.Request) {
	content := make(map[string]interface{})
	c.RenderTemplate(w, AboutWhoWeAre, content)
}

func (c *RESTFrontend) GeneralContact(w http.ResponseWriter, r *http.Request) {
	content := make(map[string]interface{})
	c.RenderTemplate(w, IndexContact, content)
}

func (c *RESTFrontend) GeneralNews(w http.ResponseWriter, r *http.Request) {
	content := make(map[string]interface{})
	c.RenderTemplate(w, IndexNews, content)
}

func (c *RESTFrontend) IndexHandler(w http.ResponseWriter, r *http.Request) {
	content := make(map[string]interface{})
	// TODO Issue#107: Replace this with proper way of detecting if root has already been created.
	found, err := c.RESTBackend.DetectRootUser()
	if err != nil {
		c.HandleError(w, fmt.Sprintf("Failed to get root user. %s", errors.WithStack(err)), http.StatusInternalServerError, IndexPath)
		return
	}

	name := IndexPage
	if !found {
		if r.Host == "polygnosics.localhost" {
			name = "auth_signup"
		} else {
			c.HandleError(w, "Server is not configured yet", http.StatusInternalServerError, IndexPath)
			return
		}
	}

	c.RenderTemplate(w, name, content)
}
