package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

const (
	UsersKey         = "users"
	UsersUsernameKey = "username"
	UsersEmailKey    = "email"
	UsersPasswordKey = "password"
	SettingsKey      = "settings"
	SettingsGroupKey = "group"

	AssetsKey = "assets"
)

type Controller struct {
	TestData map[string]interface{}
}

func NewController() (*Controller, error) {
	data, err := ioutil.ReadFile("/user-assets/testData.json")
	if err != nil {
		return nil, err
	}
	jsonData := make(map[string]interface{})
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		return nil, err
	}

	return &Controller{
		TestData: jsonData,
	}, nil
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hi! I am a dummy server!")
}

func (c *Controller) CreateRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", sayHello)
	r.HandleFunc("/update-test-data", c.updateTestData)
	r.HandleFunc("/detect-root-user", c.detectRootUser)
	r.HandleFunc("/add-user", c.addUser)
	r.HandleFunc("/login", c.login)
	return r
}

func writeResponse(data string, w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
	fmt.Fprint(w, data)
}

func decodeRequest(r *http.Request) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func encodeResponse(data map[string]interface{}, w http.ResponseWriter) {
	b, err := json.Marshal(data)
	if err != nil {
		err = errors.Wrap(errors.WithStack(err), "Failed to encode response")
		writeResponse(fmt.Sprintf("{'error':'%s'}", err.Error()), w, http.StatusInternalServerError)
		return
	}
	log.Println(string(b))
	writeResponse(string(b), w, http.StatusOK)
}

func (c *Controller) updateTestData(w http.ResponseWriter, r *http.Request) {
	requestData, err := decodeRequest(r)
	if err != nil {
		err = errors.Wrap(errors.WithStack(err), "Failed to decode request")
		writeResponse(fmt.Sprintf("{'error':'%s'}", err.Error()), w, http.StatusInternalServerError)
		return
	}

	c.TestData = requestData
}

func (c *Controller) addUser(w http.ResponseWriter, r *http.Request) {
	requestData, err := decodeRequest(r)
	if err != nil {
		err = errors.Wrap(errors.WithStack(err), "Failed to decode request")
		writeResponse(fmt.Sprintf("{'error':'%s'}", err.Error()), w, http.StatusInternalServerError)
		return
	}

	id := uuid.New()
	userData := make(map[string]interface{})
	userData[AssetsKey] = make(map[string]interface{})
	userData[SettingsKey] = make(map[string]interface{})
	userData[SettingsKey].(map[string]interface{})[SettingsGroupKey] = requestData[SettingsGroupKey]
	userData[UsersUsernameKey] = requestData[UsersUsernameKey]
	userData[UsersEmailKey] = requestData[UsersEmailKey]
	userData[UsersPasswordKey] = requestData[UsersPasswordKey]

	c.TestData[UsersKey].(map[string]interface{})[id.String()] = userData

}

func (c *Controller) login(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	if err := r.ParseForm(); err != nil {
		err = errors.Wrap(errors.WithStack(err), "Failed to parse form")
		writeResponse(fmt.Sprintf("{'error':'%s'}", err.Error()), w, http.StatusInternalServerError)
		return
	}

	email := r.FormValue(UsersEmailKey)
	pwd := r.FormValue(UsersPasswordKey)
	for _, v := range c.TestData[UsersKey].(map[string]interface{}) {
		userdata := v.(map[string]interface{})
		if userdata[UsersEmailKey].(string) == email && userdata[UsersPasswordKey].(string) == string(pwd) {
			data["data"] = v
			delete(data["data"].(map[string]interface{}), UsersPasswordKey)
		}
	}

	encodeResponse(data, w)
}

func (c *Controller) detectRootUser(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	for _, v := range c.TestData[UsersKey].(map[string]interface{}) {
		if v.(map[string]interface{})[SettingsKey].(map[string]interface{})[SettingsGroupKey] == "root" {
			data["data"] = "true"
		}
	}

	encodeResponse(data, w)
}
