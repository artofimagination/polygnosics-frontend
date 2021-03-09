package restfrontend

import (
	"fmt"
	"net/http"

	"polygnosics-frontend/restfrontend/session"

	"github.com/pkg/errors"
)

// LoginHandler checks the user email and password.
// On success generates and stores a cookie in the session sotre and adds it to the response
func (c *RESTFrontend) LoginHandler(w http.ResponseWriter, r *http.Request) {
	content := c.ContentController.BuildLoginContent()
	if r.Method == GET {
		c.RenderTemplate(w, "auth_login", content)
	} else {
		name := "confirm"
		p := make(map[string]interface{})

		if err := r.ParseForm(); err != nil {
			p["message"] = ErrFailedToParseForm
			c.RenderTemplate(w, name, p)
			return
		}
		email := r.FormValue("email")
		pwd := []byte(r.FormValue("psw"))

		user, err := c.RESTBackend.Login(email, pwd)
		if err != nil {
			c.HandleError(w, fmt.Sprintf("Failed to get cookie. %s", errors.WithStack(err)), http.StatusInternalServerError, IndexLoginPath)
		}
		c.ContentController.User = user

		// Create session cookie.
		sess, err := session.Store.Get(r, "cookie-name")
		if err != nil {
			p["message"] = fmt.Sprintf("Failed to create cookie. %s", errors.WithStack(err))
			c.RenderTemplate(w, name, p)
			return
		}
		sess.Options.MaxAge = 60000
		sess.Values["authenticated"] = true
		sess.Values["user"] = user.ID

		cookieKey, err := session.EncryptUserAndOrigin(user.ID, r.RemoteAddr)
		if err != nil {
			p["message"] = fmt.Sprintf("Failed to generate cookie data. %s", errors.WithStack(err))
			c.RenderTemplate(w, name, p)
			return
		}
		sess.Values["cookie_key"] = cookieKey

		if err := sess.Save(r, w); err != nil {
			p["message"] = fmt.Sprintf("Failed to save cookie. %s", errors.WithStack(err))
			c.RenderTemplate(w, name, p)
			return
		}

		http.Redirect(w, r, UserMainPath, http.StatusSeeOther)
	}
}

func (c *RESTFrontend) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := session.Store.Get(r, "cookie-name")
	if err != nil {
		c.HandleError(w, fmt.Sprintf("Failed to get cookie. %s", errors.WithStack(err)), http.StatusInternalServerError, IndexPath)
		return
	}

	// Revoke users authentication
	session.Values["authenticated"] = false
	if err := session.Save(r, w); err != nil {
		c.HandleError(w, fmt.Sprintf("Failed to save cookie. %s", errors.WithStack(err)), http.StatusInternalServerError, IndexPath)
		return
	}

	http.Redirect(w, r, IndexPath, http.StatusSeeOther)
}

func (c *RESTFrontend) SignupHandler(w http.ResponseWriter, r *http.Request) {
	content := c.ContentController.BuildSignupContent()
	if r.Method == GET {
		c.RenderTemplate(w, "auth_signup", content)
	} else {
		if err := r.ParseForm(); err != nil {
			c.HandleError(w, fmt.Sprintf("Failed to parse form. %s", errors.WithStack(err)), http.StatusInternalServerError, IndexPath)
			return
		}
		uName := r.FormValue("username")
		email := r.FormValue("email")
		pwd := []byte(r.FormValue("psw"))
		group := r.FormValue("developer")
		if group == "" {
			group = "client"
		}

		if err := c.RESTBackend.AddUser(uName, email, pwd, group); err != nil {
			c.HandleError(w, fmt.Sprintf("Failed to add user. %s", errors.WithStack(err)), http.StatusInternalServerError, IndexPath)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Registration successful")
	}
}
