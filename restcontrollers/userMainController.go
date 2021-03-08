package restcontrollers

import (
	"fmt"
	"net/http"

	"polygnosics/app/businesslogic"
	"polygnosics/web/contents"

	"github.com/pkg/errors"
)

// UserMainHandler renders the main page after login.
func (c *RESTController) UserMainHandler(w http.ResponseWriter, r *http.Request) {
	content, err := c.ContentController.BuildUserMainContent()
	if err != nil {
		errString := fmt.Sprintf("Failed to get home page content. %s", errors.WithStack(err))
		c.RenderTemplate(w, UserMain, c.ContentController.BuildErrorContent(errString))
		return
	}
	c.RenderTemplate(w, UserMain, content)
}

// UploadAvatarHandler processes avatar upload request.
// Stores the image in the location defined by the asset ID and avatar ID.
// The file is named by the avatar ID and the folder is determined by the asset ID.
func (c *RESTController) UploadAvatarHandler(w http.ResponseWriter, r *http.Request) {
	p := c.ContentController.GetUserContent(c.ContentController.UserData)

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		c.HandleError(w, fmt.Sprintf("Failed to parse form. %s", errors.WithStack(err)), http.StatusInternalServerError, UserMainPath)
	}

	if c.ContentController.UserData == nil {
		c.HandleError(w, "User is not configured", http.StatusInternalServerError, UserMainPath)
	}

	err := c.BackendContext.UploadFile(c.ContentController.UserData.Assets, businesslogic.UserAvatarKey, businesslogic.DefaultUserAvatarPath, r)
	if err != nil {
		c.HandleError(w, fmt.Sprintf("Failed to upload asset. %s", errors.WithStack(err)), http.StatusInternalServerError, UserMainPath)
	}

	if err := c.UserDBController.UpdateUserAssets(c.ContentController.UserData); err != nil {
		c.HandleError(w, fmt.Sprintf("Failed to update asset. %s", errors.WithStack(err)), http.StatusInternalServerError, UserMainPath)
		return
	}
	p[contents.UserMapKey].(map[string]interface{})[businesslogic.UserAvatarKey] =
		c.UserDBController.ModelFunctions.GetFilePath(
			c.ContentController.UserData.Assets,
			businesslogic.UserAvatarKey,
			businesslogic.DefaultUserAvatarPath)
	http.Redirect(w, r, "profile", http.StatusSeeOther)
}

func (c *RESTController) MailInbox(w http.ResponseWriter, r *http.Request) {
	content := c.ContentController.BuildMailInboxContent()
	c.RenderTemplate(w, UserMainMailInbox, content)
}

func (c *RESTController) MailCompose(w http.ResponseWriter, r *http.Request) {
	content := c.ContentController.BuildMailComposeContent()
	c.RenderTemplate(w, UserMainMailCompose, content)
}

func (c *RESTController) MailRead(w http.ResponseWriter, r *http.Request) {
	content := c.ContentController.BuildMailReadContent()
	c.RenderTemplate(w, UserMainMailRead, content)
}

func (c *RESTController) Settings(w http.ResponseWriter, r *http.Request) {
	content := c.ContentController.BuildSettingsContent()
	c.RenderTemplate(w, UserMainSettings, content)
}
