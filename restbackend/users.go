package restbackend

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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

func (c *RESTBackend) Login(email string, password []byte) (*User, error) {
	params := fmt.Sprintf("?email=%s&password=%s", email, password)
	userData, err := get(BusinessLogicServerAddress, UserPathLogin, params)
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

func (c *RESTBackend) GetUserByID(requestedID string) (*User, error) {
	params := fmt.Sprintf("?requestedID=%s", requestedID)
	userData, err := get(BusinessLogicServerAddress, UserPathLogin, params)
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

func (c *RESTBackend) UpdateUserAvatar(r *http.Request) error {
	err := forwardRequest(BusinessLogicServerAddress, r)
	if err != nil {
		return err
	}
	return nil
}

func (c *RESTBackend) AddUser(username string, email string, password []byte, group string) error {
	params := make(map[string]interface{})
	params["username"] = username
	params["email"] = email
	params["password"] = string(password)
	params["group"] = group
	err := post(BusinessLogicServerAddress, UserPathAdd, params)
	if err != nil {
		return err
	}
	return nil
}

func (c *RESTBackend) DetectRootUser() (bool, error) {
	data, err := get(BusinessLogicServerAddress, UserPathDetectRootUser, "?nil")
	if err != nil {
		return false, err
	}

	found, err := strconv.ParseBool(data.(string))
	if err != nil {
		return false, err
	}
	return found, nil
}

func (c *RESTBackend) UpdateUser(r *http.Request) error {
	err := forwardRequest(BusinessLogicServerAddress, r)
	if err != nil {
		return err
	}
	return nil
}
