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
	UsersIDKey       = "id"
	UsersPasswordKey = "password"
	UserGroupKey     = "group"

	SettingsKey = "settings"
	DetailsKey  = "details"
	AssetsKey   = "assets"
)

const (
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
	CategoryMLKey           = "machine_learning"
	CategoryMLText          = "Machine Learning"
	CategoryCivilEngNameKey = "civil_eng"
	CategoryCivilEngText    = "Civil Engineering"
	CategoryMedicineKey     = "medicine"
	CategoryMedicineText    = "Medicine"
	CategoryChemistryKey    = "chemistry"
	CategoryChemistryText   = "Chemistry"
)

type Controller struct {
	TestData map[string]interface{}
}

func createCategoriesMap() map[string]string {
	categoriesMap := make(map[string]string)
	categoriesMap[CategoryMLKey] = CategoryMLText
	categoriesMap[CategoryCivilEngNameKey] = CategoryCivilEngText
	categoriesMap[CategoryMedicineKey] = CategoryMedicineText
	return categoriesMap
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
	r.HandleFunc("/get-user-by-id", c.getUserByID)
	r.HandleFunc("/login", c.login)
	r.HandleFunc("/add-product", c.addProduct)
	return r
}

func (c *Controller) addUser(w http.ResponseWriter, r *http.Request) {
	requestData, err := decodeRequest(r)
	if err != nil {
		err = errors.Wrap(errors.WithStack(err), "Backend: Failed to decode request")
		writeResponse(fmt.Sprintf("{\"error\":\"%s\"}", err.Error()), w, http.StatusInternalServerError)
		return
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

}

func (c *Controller) login(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	if err := r.ParseForm(); err != nil {
		err = errors.Wrap(errors.WithStack(err), "Backend: Failed to parse form")
		writeResponse(fmt.Sprintf("{\"error\":\"%s\"}", err.Error()), w, http.StatusInternalServerError)
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
		if v.(map[string]interface{})[SettingsKey].(map[string]interface{})[UserGroupKey] == "root" {
			data["data"] = "true"
		}
	}

	encodeResponse(data, w)
}

func (c *Controller) getUserByID(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	if err := r.ParseForm(); err != nil {
		err = errors.Wrap(errors.WithStack(err), "Backend: Failed to parse form")
		writeResponse(fmt.Sprintf("{\"error\":\"%s\"}", err.Error()), w, http.StatusInternalServerError)
		return
	}

	id := r.FormValue(UsersIDKey)
	for k, v := range c.TestData[UsersKey].(map[string]interface{}) {
		if k == id {
			data["data"] = v
		}
	}

	encodeResponse(data, w)
}

func (c *Controller) addProduct(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		err = errors.Wrap(errors.WithStack(err), "Backend: Failed to parse form")
		writeResponse(fmt.Sprintf("{\"error\":\"%s\"}", err.Error()), w, http.StatusInternalServerError)
		return
	}

	log.Println("Add")
	if err := uploadFile(ProductAvatarKey, "avatar.jpg", r); err != nil {
		err = errors.Wrap(errors.WithStack(err), "Backend: Failed to upload avatar")
		writeResponse(fmt.Sprintf("{'error':'%s'}", err.Error()), w, http.StatusInternalServerError)
		return
	}
	log.Println("Add2")

	if err := uploadFile(ProductMainAppKey, "main-app.tar.gz", r); err != nil {
		err = errors.Wrap(errors.WithStack(err), "Backend: Failed to upload avatar")
		writeResponse(fmt.Sprintf("{\"error\":\"%s\"}", err.Error()), w, http.StatusInternalServerError)
		return
	}
	log.Println("Add3")

	if err := uploadFile(ProductClientAppKey, "client-app.tar.gz", r); err != nil {
		err = errors.Wrap(errors.WithStack(err), "Backend: Failed to upload avatar")
		writeResponse(fmt.Sprintf("{\"error\":\"%s\"}", err.Error()), w, http.StatusInternalServerError)
		return
	}
	log.Println("Add4")

	id := uuid.New()
	product := make(map[string]interface{})
	product[id.String()] = make(map[string]interface{})
	productData := product[id.String()].(map[string]interface{})
	productData[AssetsKey] = make(map[string]interface{})
	productData[DetailsKey] = make(map[string]interface{})

	productData[AssetsKey].(map[string]interface{})[ProductAvatarKey] = fmt.Sprintf("/user-assets/uploads/avatar.jpg")
	productData[AssetsKey].(map[string]interface{})[ProductMainAppKey] = fmt.Sprintf("/user-assets/uploads/main-app.tar.gz")
	productData[AssetsKey].(map[string]interface{})[ProductClientAppKey] = fmt.Sprintf("/user-assets/uploads/client-app.tar.gz")

	productData[DetailsKey].(map[string]interface{})[ProductNameKey] = r.FormValue(ProductNameKey)
	productData[DetailsKey].(map[string]interface{})[ProductPriceKey] = r.FormValue(ProductPriceKey)
	productData[DetailsKey].(map[string]interface{})[ProductPriceKey] = r.FormValue(ProductPriceKey)
	productData[DetailsKey].(map[string]interface{})[ProductDescriptionKey] = r.FormValue(ProductDescriptionKey)
	productData[DetailsKey].(map[string]interface{})[ProductShortDescriptionKey] = r.FormValue(ProductShortDescriptionKey)
	productData[DetailsKey].(map[string]interface{})[ProductURLKey] = r.FormValue(ProductURLKey)
	categories := createCategoriesMap()
	categoryList := make([]string, 0)
	for k := range categories {
		if r.FormValue(k) == "checked" {
			categoryList = append(categoryList, k)
		}
	}
	productData[DetailsKey].(map[string]interface{})[ProductCategoriesKey] = categoryList
	productData[DetailsKey].(map[string]interface{})[ProductRequires3DKey] = r.FormValue(ProductRequires3DKey)
	productData[DetailsKey].(map[string]interface{})[ProductPublicKey] = r.FormValue(ProductPublicKey)
	productData[DetailsKey].(map[string]interface{})[ProductTagsKey] = r.FormValue(ProductTagsKey)

	c.TestData["products"] = product
	log.Println("Add5")
	prettyPrint(c.TestData)
	writeResponse("{'data':'ok'}", w, http.StatusOK)
}
