package restfrontend

import (
	"fmt"
	"net/http"

	"github.com/artofimagination/polygnosics-frontend/restfrontend/session"

	"github.com/pkg/errors"
)

// LoginHandler checks the user email and password.
// On success generates and stores a cookie in the session sotre and adds it to the response
func (c *RESTFrontend) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == GET {
		content := c.ContentController.BuildLoginContent()
		c.RenderTemplate(w, "auth_login", content)
	} else {
		if err := r.ParseForm(); err != nil {
			c.HandleError(w, fmt.Sprintf("Failed to parse form. %s", errors.WithStack(err)), http.StatusInternalServerError, IndexLoginPath)
			return
		}
		email := r.FormValue("email")
		pwd := []byte(r.FormValue("password"))

		user, err := c.RESTBackend.Login(email, pwd)
		if err != nil {
			w.WriteHeader(http.StatusAccepted)
			fmt.Fprintf(w, "Failed to login. %s", errors.WithStack(err))
			return
		}
		c.ContentController.User = user

		// Create session cookie.
		sess, err := session.Store.Get(r, "cookie-name")
		if err != nil {
			c.HandleError(w, fmt.Sprintf("Failed to create cookie. %s", errors.WithStack(err)), http.StatusInternalServerError, IndexLoginPath)
			return
		}
		sess.Options.MaxAge = 60000
		sess.Values["authenticated"] = true
		sess.Values["user"] = c.ContentController.User.ID

		cookieKey, err := session.EncryptUserAndOrigin(c.ContentController.User.ID, r.RemoteAddr)
		if err != nil {
			c.HandleError(w, fmt.Sprintf("Failed to generate cookie data. %s", errors.WithStack(err)), http.StatusInternalServerError, IndexLoginPath)
			return
		}
		sess.Values["cookie_key"] = cookieKey

		if err := sess.Save(r, w); err != nil {
			c.HandleError(w, fmt.Sprintf("Failed to save cookie. %s", errors.WithStack(err)), http.StatusInternalServerError, IndexLoginPath)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Login successful")
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
	c.ContentController.User = nil

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
		pwd := []byte(r.FormValue("password"))
		group := r.FormValue("group")
		if group == "" {
			group = "client"
		}

		if uName == "" || email == "" || r.FormValue("password") == "" {
			c.HandleError(w, "Form values are empty", http.StatusBadRequest, IndexPath)
			return
		}

		if err := c.RESTBackend.AddUser(uName, email, pwd, group); err != nil {
			w.WriteHeader(http.StatusAccepted)
			fmt.Fprint(w, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Registration successful")
	}
}
