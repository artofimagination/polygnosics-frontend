package rest

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

const (
	UsersKey         = "users"
	UserProductsKey  = "user_products"
	UsersProjectKey  = "user_projects"
	UsersUsernameKey = "username"
	UsersEmailKey    = "email"
	UsersIDKey       = "id"
	UsersPasswordKey = "password"
	UserGroupKey     = "group"

	SettingsKey = "settings"
	DetailsKey  = "details"
	AssetsKey   = "assets"
)

const (
	ProductKey                 = "products"
	ProductAvatarKey           = "avatar"
	ProductMainAppKey          = "main_app"
	ProductClientAppKey        = "client_app"
	ProductDescriptionKey      = "description"
	ProductShortDescriptionKey = "short_description"
	ProductNameKey             = "name"
	ProductRequires3DKey       = "requires_3d"
	ProductURLKey              = "url"
	ProductPublicKey           = "is_public"
	ProductPricingKey          = "pricing"
	ProductPriceKey            = "amount"
	ProductTagsKey             = "tags"
	ProductCategoriesKey       = "categories"
)

const (
	ProjectKey = "projects"
)

const (
	CategoriesKey = "categories"
)

func convertCheckboxValueToText(input string) string {
	if input == "" {
		return "unchecked"
	}
	return input
}

func (c *Controller) addUser(w http.ResponseWriter, r *http.Request) {
	requestData, err := c.decodeRequest(r)
	if err != nil {
		writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusBadRequest)
		return
	}

	for _, v := range c.TestData[UsersKey].(map[string]interface{}) {
		if v.(map[string]interface{})[UsersUsernameKey] == requestData[UsersUsernameKey] {
			writeError("User already exists", w, http.StatusAccepted)
			return
		}
	}
	id := uuid.New()
	userData := make(map[string]interface{})
	userData[AssetsKey] = make(map[string]interface{})
	userData[SettingsKey] = make(map[string]interface{})
	userData[SettingsKey].(map[string]interface{})[UserGroupKey] = requestData[UserGroupKey]
	userData[UsersUsernameKey] = requestData[UsersUsernameKey]
	userData[UsersEmailKey] = requestData[UsersEmailKey]
	userData[UsersPasswordKey] = requestData[UsersPasswordKey]
	c.TestData[UsersKey].(map[string]interface{})[id.String()] = userData
	writeData("OK", w, http.StatusCreated)
}

func (c *Controller) login(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	if err := c.ParseForm(r); err != nil {
		writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusBadRequest)
		return
	}

	email := r.FormValue(UsersEmailKey)
	pwd := r.FormValue(UsersPasswordKey)
	data["data"] = make(map[string]interface{})
	for _, v := range c.TestData[UsersKey].(map[string]interface{}) {
		if v.(map[string]interface{})[UsersEmailKey].(string) == email && v.(map[string]interface{})[UsersPasswordKey].(string) == pwd {
			for k, value := range v.(map[string]interface{}) {
				if k != UsersPasswordKey {
					data["data"].(map[string]interface{})[k] = value
				}
			}
			encodeResponse(data, w)
			return
		}
	}
	writeError("Incorrect email or password", w, http.StatusAccepted)
}

func (c *Controller) detectRootUser(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	for _, v := range c.TestData[UsersKey].(map[string]interface{}) {
		if v.(map[string]interface{})[SettingsKey].(map[string]interface{})[UserGroupKey] == "root" {
			data["data"] = "true"
		}
	}

	encodeResponse(data, w)
}

func (c *Controller) getUserByID(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	if err := c.ParseForm(r); err != nil {
		writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusBadRequest)
		return
	}

	id := r.FormValue(UsersIDKey)
	for k, v := range c.TestData[UsersKey].(map[string]interface{}) {
		if k == id {
			data["data"] = v
			break
		}
	}

	encodeResponse(data, w)
}

func (c *Controller) addProduct(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusBadRequest)
		return
	}

	if err := uploadFile(ProductAvatarKey, "avatar.jpg", r); err != nil {
		writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusInternalServerError)
		return
	}

	if err := uploadFile(ProductMainAppKey, "main-app.tar.gz", r); err != nil {
		writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusInternalServerError)
		return
	}

	if err := uploadFile(ProductClientAppKey, "client-app.tar.gz", r); err != nil {
		writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusInternalServerError)
		return
	}

	id := uuid.New()
	product := make(map[string]interface{})
	product[id.String()] = make(map[string]interface{})
	productData := product[id.String()].(map[string]interface{})
	productData[AssetsKey] = make(map[string]interface{})
	assets := productData[AssetsKey].(map[string]interface{})
	productData[DetailsKey] = make(map[string]interface{})
	details := productData[DetailsKey].(map[string]interface{})

	assets[ProductAvatarKey] = fmt.Sprintf("/backend/uploads/avatar.jpg")
	assets[ProductMainAppKey] = fmt.Sprintf("/backend/uploads/main-app.tar.gz")
	assets[ProductClientAppKey] = fmt.Sprintf("/backend/uploads/client-app.tar.gz")

	details[ProductNameKey] = r.FormValue(ProductNameKey)
	details[ProductPriceKey] = r.FormValue(ProductPriceKey)
	details[ProductPriceKey] = r.FormValue(ProductPriceKey)
	details[ProductDescriptionKey] = r.FormValue(ProductDescriptionKey)
	details[ProductShortDescriptionKey] = r.FormValue(ProductShortDescriptionKey)
	details[ProductURLKey] = r.FormValue(ProductURLKey)
	categoryList := make([]string, 0)
	for k := range c.TestData[CategoriesKey].(map[string]interface{}) {
		if r.FormValue(k) == "checked" {
			categoryList = append(categoryList, k)
		}
	}
	details[ProductCategoriesKey] = categoryList
	details[ProductRequires3DKey] = convertCheckboxValueToText(r.FormValue(ProductRequires3DKey))
	details[ProductPublicKey] = convertCheckboxValueToText(r.FormValue(ProductPublicKey))
	details[ProductTagsKey] = r.FormValue(ProductTagsKey)

	c.TestData[ProductKey] = product
	c.TestData[UserProductsKey].(map[string]interface{})[r.FormValue("user")] = id
	writeData("OK", w, http.StatusCreated)
}

func (c *Controller) getProductsByUserID(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	if err := c.ParseForm(r); err != nil {
		writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusBadRequest)
		return
	}
	data["data"] = make([]interface{}, 0)
	id := r.FormValue(UsersIDKey)
	for userKey, productKey := range c.TestData[UserProductsKey].(map[string]interface{}) {
		if userKey == id {
			for k, v := range c.TestData[ProductKey].(map[string]interface{}) {
				if productKey == k {
					data["data"] = append(data["data"].([]interface{}), v)
					break
				}
			}
		}
	}

	encodeResponse(data, w)
}

func (c *Controller) getProjectsByUserID(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	if err := c.ParseForm(r); err != nil {
		writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusBadRequest)
		return
	}
	data["data"] = make([]interface{}, 0)
	id := r.FormValue(UsersIDKey)
	for userKey, projectKey := range c.TestData[UsersProjectKey].(map[string]interface{}) {
		if userKey == id {
			for k, v := range c.TestData[ProjectKey].(map[string]interface{}) {
				if projectKey == k {
					data["data"] = append(data["data"].([]interface{}), v)
					break
				}
			}
		}
	}

	encodeResponse(data, w)
}

func (c *Controller) getCategoriesMap(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["data"] = c.TestData[CategoriesKey]
	encodeResponse(data, w)
}
