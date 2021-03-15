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

const (
	ProductPathAdd       = "/add-product"
	ProductPathGet       = "/get-product"
	ProductPathUpdate    = "/update-product"
	ProductPathDelete    = "/delete-product"
	ProductPathGetByUser = "/get-products-by-user"
)

const (
	ProjectPathAdd          = "/add-project"
	ProjectPathGet          = "/get-project"
	ProjectPathUpdate       = "/update-project"
	ProjectPathDelete       = "/delete-project"
	ProjectPathState        = "/get-project-state"
	ProjectPathRequestState = "/request-state-change"
	ProjectPathGetByUser    = "/get-projects-by-user"
)

const (
	CategoriesPathGet = "/get-categories"
)

type User struct {
	ID       string `json:"id" validate:"required"`
	UserName string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Assets   map[string]interface{}
	Settings map[string]interface{}
}

type Product struct {
	ID        string `json:"id" validate:"required"`
	Privilege int    `json:"privilege" validate:"required"`
	Assets    map[string]interface{}
	Details   map[string]interface{}
}

type Project struct {
	ID      string `json:"id" validate:"required"`
	Assets  map[string]interface{}
	Details map[string]interface{}
}

func (c *RESTBackend) Login(email string, password []byte) (*User, error) {
	params := fmt.Sprintf("?email=%s&password=%s", email, string(password))
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
	_, err := forwardRequest(BusinessLogicServerAddress, UserPathUpdate, r)
	if err != nil {
		return err
	}
	return nil
}

func (c *RESTBackend) AddUser(username string, email string, password []byte, group string) error {
	params := make(map[string]interface{})
	params["name"] = username
	params["email"] = email
	params["password"] = password
	params["group"] = group
	_, err := post(BusinessLogicServerAddress, UserPathAdd, params)
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
	_, err := forwardRequest(BusinessLogicServerAddress, UserPathUpdate, r)
	if err != nil {
		return err
	}
	return nil
}

func (c *RESTBackend) GetProduct(id string) (*Product, error) {
	params := fmt.Sprintf("?id=%s", id)
	product := &Product{}
	data, err := get(BusinessLogicServerAddress, ProductPathGet, params)
	if err != nil {
		return nil, err
	}
	productDataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(productDataBytes, product); err != nil {
		return nil, err
	}
	return product, nil
}

func (c *RESTBackend) GetProductsByUserID(userID string) ([]*Product, error) {
	products := make([]*Product, 0)
	params := fmt.Sprintf("?user-id=%s", userID)
	data, err := get(BusinessLogicServerAddress, ProductPathGetByUser, params)
	if err != nil {
		return nil, err
	}
	productDataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(productDataBytes, &products); err != nil {
		return nil, err
	}
	return products, nil
}

func (c *RESTBackend) AddProduct(w http.ResponseWriter, r *http.Request) error {
	_, err := forwardRequest(BusinessLogicServerAddress, ProductPathAdd, r)
	if err != nil {
		return err
	}
	return nil
}

func (c *RESTBackend) UpdateProduct(r *http.Request) error {
	_, err := forwardRequest(BusinessLogicServerAddress, ProductPathUpdate, r)
	if err != nil {
		return err
	}
	return nil
}

func (c *RESTBackend) DeleteProduct(productID string) error {
	params := make(map[string]interface{})
	params["id"] = productID
	_, err := post(BusinessLogicServerAddress, ProductPathDelete, params)
	if err != nil {
		return err
	}
	return nil
}

func (c *RESTBackend) GetCategoriesMap() (map[string]interface{}, error) {
	categories := make(map[string]interface{})
	_, err := get(BusinessLogicServerAddress, CategoriesPathGet, "?nil")
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (c *RESTBackend) CreateProject(r *http.Request) error {
	_, err := forwardRequest(BusinessLogicServerAddress, ProjectPathAdd, r)
	if err != nil {
		return err
	}
	return nil
}

func (c *RESTBackend) GetProject(id string) (*Project, error) {
	params := fmt.Sprintf("?id=%s", id)
	project := &Project{}
	data, err := get(BusinessLogicServerAddress, ProjectPathGet, params)
	if err != nil {
		return nil, err
	}
	projectDataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(projectDataBytes, project); err != nil {
		return nil, err
	}
	return project, nil
}

func (c *RESTBackend) CheckProjectState(id string) (string, error) {
	params := fmt.Sprintf("?id=%s", id)
	_, err := get(BusinessLogicServerAddress, ProjectPathState, params)
	if err != nil {
		return "", err
	}
	return "", nil
}

func (c *RESTBackend) RunProject(userID string, projectID string) error {
	params := fmt.Sprintf("?user-id=%s&project-id=%s&state=run", userID, projectID)
	_, err := get(BusinessLogicServerAddress, ProjectPathRequestState, params)
	if err != nil {
		return err
	}
	return nil
}

func (c *RESTBackend) GetProjectsByUserID(userID string) ([]*Project, error) {
	projects := make([]*Project, 0)
	params := fmt.Sprintf("?user-id=%s", userID)
	data, err := get(BusinessLogicServerAddress, ProjectPathGetByUser, params)
	if err != nil {
		return nil, err
	}
	projectDataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(projectDataBytes, &projects); err != nil {
		return nil, err
	}
	return projects, nil
}

func (c *RESTBackend) DeleteProject(projectID string) error {
	params := make(map[string]interface{})
	params["id"] = projectID
	_, err := post(BusinessLogicServerAddress, ProjectPathDelete, params)
	if err != nil {
		return err
	}
	return nil
}

func (c *RESTBackend) UpdateProject(r *http.Request) error {
	_, err := forwardRequest(BusinessLogicServerAddress, ProductPathAdd, r)
	if err != nil {
		return err
	}
	return nil
}

// TODO Issue#132: Implement this on the stats service side.
// offer := r.FormValue("offer")
// statsFunc, err := c.BackendContext.GetDataChannelProvider(r.FormValue("type"))
// if err != nil {
// 	c.HandleError(w, fmt.Sprintf("Failed to get webrtc data provider. %s", errors.WithStack(err)), http.StatusInternalServerError, UserMainPath)
// 	return
// }

// if err := webrtc.SetupFrontend(w, r, offer, statsFunc); err != nil {
// 	c.HandleError(w, fmt.Sprintf("Failed to start frontend webrtc. %s", errors.WithStack(err)), http.StatusInternalServerError, UserMainPath)
// 	return
// }
func (c *RESTBackend) InitStatsWebRTC(r *http.Request) error {
	_, err := forwardRequest(BusinessLogicServerAddress, ProductPathAdd, r)
	if err != nil {
		return err
	}
	return nil
}
