package restfrontend

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// ProfileHandler renders the profile page template.
func (c *RESTFrontend) ProfileHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		c.RenderTemplate(w, UserMain, c.ContentController.BuildErrorContent(ErrFailedToParseForm))
		return
	}

	content, err := c.ContentController.BuildProfileContent(r.FormValue("user"))
	if err != nil {
		errString := fmt.Sprintf("Failed to get profile page content. %s", errors.WithStack(err))
		c.RenderTemplate(w, UserMain, c.ContentController.BuildErrorContent(errString))
		return
	}

	c.RenderTemplate(w, "profile", content)
}

func (c *RESTFrontend) ProfileEdit(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		content, err := c.ContentController.BuildProfileContent(c.ContentController.User.ID)
		if err != nil {
			errString := fmt.Sprintf("Failed to get profile page content. %s", errors.WithStack(err))
			c.RenderTemplate(w, UserMain, c.ContentController.BuildErrorContent(errString))
			return
		}

		c.RenderTemplate(w, "profile-edit", content)
	} else {
		if err := r.ParseForm(); err != nil {
			c.RenderTemplate(w, UserMain, c.ContentController.BuildErrorContent(ErrFailedToParseForm))
			return
		}

		c.ContentController.StoreUserInfo(r)

		content, err := c.ContentController.BuildProfileContent(c.ContentController.User.ID)
		if err != nil {
			errString := fmt.Sprintf("Failed to get profile page content. %s", errors.WithStack(err))
			c.RenderTemplate(w, UserMain, c.ContentController.BuildErrorContent(errString))
			return
		}

		c.RenderTemplate(w, "profile", content)
	}

}
