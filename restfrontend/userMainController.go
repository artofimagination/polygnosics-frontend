package restfrontend

import (
	"fmt"
	"net/http"

	"polygnosics-frontend/contents"

	"github.com/pkg/errors"
)

// UserMainHandler renders the main page after login.
func (c *RESTFrontend) UserMainHandler(w http.ResponseWriter, r *http.Request) {
	content, err := c.ContentController.BuildUserMainContent()
	if err != nil {
		c.HandleError(w, fmt.Sprintf("Failed to get home page content. %s", errors.WithStack(err)), http.StatusInternalServerError, IndexPath)
		return
	}

	c.RenderTemplate(w, UserMain, content)
}

// UploadAvatarHandler processes avatar upload request.
// Stores the image in the location defined by the asset ID and avatar ID.
// The file is named by the avatar ID and the folder is determined by the asset ID.
func (c *RESTFrontend) UploadAvatarHandler(w http.ResponseWriter, r *http.Request) {
	content := c.ContentController.GetUserContent(c.ContentController.User)

	if err := c.RESTBackend.UpdateUserAvatar(r); err != nil {
		c.HandleError(w, fmt.Sprintf("Failed to update avatar. %s", errors.WithStack(err)), http.StatusInternalServerError, UserMainPath)
	}

	content[contents.UserMapKey].(map[string]interface{})[contents.UserAvatarKey] = c.ContentController.User.Assets[contents.UserAvatarKey]
	http.Redirect(w, r, "profile", http.StatusSeeOther)
}

func (c *RESTFrontend) MailInbox(w http.ResponseWriter, r *http.Request) {
	content := c.ContentController.BuildMailInboxContent()
	c.RenderTemplate(w, UserMainMailInbox, content)
}

func (c *RESTFrontend) MailCompose(w http.ResponseWriter, r *http.Request) {
	content := c.ContentController.BuildMailComposeContent()
	c.RenderTemplate(w, UserMainMailCompose, content)
}

func (c *RESTFrontend) MailRead(w http.ResponseWriter, r *http.Request) {
	content := c.ContentController.BuildMailReadContent()
	c.RenderTemplate(w, UserMainMailRead, content)
}

func (c *RESTFrontend) Settings(w http.ResponseWriter, r *http.Request) {
	content := c.ContentController.BuildSettingsContent()
	c.RenderTemplate(w, UserMainSettings, content)
}
