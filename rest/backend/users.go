package backend

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	UserPathLogin          = "/login"
	UserPathAdd            = "/add-user"
	UserPathDetectRootUser = "/detect-root-user"
	UserPathGetUserByID    = "/get-user-by-id"
	UserPathUpdate         = "/update-user"
)

type User struct {
	ID       string `json:"id" validate:"required"`
	UserName string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Assets   map[string]interface{}
	Settings map[string]interface{}
}

func (c *RESTController) Login(email string, password []byte) (*User, error) {
	params := fmt.Sprintf("?email=%s&password=%s", email, password)
	userData, err := c.Get(UserPathLogin, params)
	if err != nil {
		return nil, err
	}
	user := &User{}
	userDataBytes, err := json.Marshal(userData)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(userDataBytes, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (c *RESTController) GetUserByID(requestedID string) (*User, error) {
	params := fmt.Sprintf("?requestedID=%s", requestedID)
	userData, err := c.Get(UserPathLogin, params)
	if err != nil {
		return nil, err
	}

	user := &User{}
	userDataBytes, err := json.Marshal(userData)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(userDataBytes, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (c *RESTController) UpdateUserAvatar(r *http.Request) error {
	_, err := c.ForwardRequest(r)
	if err != nil {
		return err
	}
	return nil
}

func (c *RESTController) AddUser(username string, email string, password []byte, group string) error {
	params := make(map[string]interface{})
	params["username"] = username
	params["email"] = email
	params["password"] = string(password)
	params["group"] = group
	err := c.Post(UserPathAdd, params)
	if err != nil {
		return err
	}
	return nil
}

func (c *RESTController) DetectRootUser() (bool, error) {
	found, err := c.Get(UserPathDetectRootUser, "?nil")
	if err != nil {
		return false, err
	}

	return found.(bool), nil
}

func (c *RESTController) UpdateUser(r *http.Request) error {
	_, err := c.ForwardRequest(r)
	if err != nil {
		return err
	}
	return nil
}
