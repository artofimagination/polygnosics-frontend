package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

type Controller struct {
	TestData    map[string]interface{}
	RequestData map[string]interface{}
}

func prettyPrint(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
		return
	}
	fmt.Println("Failed to pretty print data")
}

func writeError(message string, w http.ResponseWriter, statusCode int) {
	writeResponse(fmt.Sprintf("{\"error\":\"%s\"}", message), w, statusCode)
}

func writeData(data string, w http.ResponseWriter, statusCode int) {
	writeResponse(fmt.Sprintf("{\"data\":\"%s\"}", data), w, statusCode)
}

func writeResponse(data string, w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
	fmt.Fprint(w, data)
}

func (c *Controller) getRequestData(w http.ResponseWriter, r *http.Request) {
	encodeResponse(c.RequestData, w)
}

func (c Controller) ParseForm(r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	c.setRequestData(r)
	return nil
}

func (c Controller) ParseMultipartForm(r *http.Request) error {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		return err
	}

	c.setRequestData(r)
	return nil
}

func (c *Controller) setRequestData(r *http.Request) {
	for k := range c.RequestData {
		delete(c.RequestData, k)
	}

	c.RequestData["uri"] = r.RequestURI
}

func (c *Controller) decodeRequest(r *http.Request) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	c.RequestData = data
	return data, nil
}

func encodeResponse(data map[string]interface{}, w http.ResponseWriter) {
	b, err := json.Marshal(data)
	if err != nil {
		writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusInternalServerError)
		return
	}
	writeResponse(string(b), w, http.StatusOK)
}

func (c *Controller) updateTestData(w http.ResponseWriter, r *http.Request) {
	requestData, err := c.decodeRequest(r)
	if err != nil {
		writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusInternalServerError)
		return
	}

	c.TestData = requestData
}

func uploadFile(key string, fileName string, r *http.Request) error {
	file, handler, err := r.FormFile(key)
	if err == http.ErrMissingFile {
		return nil
	}

	if err != nil {
		return err
	}

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	defer file.Close()

	// Create file
	dst, err := os.Create(fmt.Sprintf("/user-assets/uploads/%s", fileName))
	if err != nil {
		return err
	}

	// Copy the uploaded file to the created file on the file system.
	if _, err := io.Copy(dst, file); err != nil {
		if err2 := dst.Close(); err2 != nil {
			err = errors.Wrap(errors.WithStack(err), err2.Error())
		}
		return err
	}
	dst.Close()

	return nil
}

func WriteToFile(filename string, data string) error {
	if filename == "" {
		return nil
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
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

	requestData := make(map[string]interface{})
	return &Controller{
		TestData:    jsonData,
		RequestData: requestData,
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
	userMain := r.PathPrefix("/user-main").Subrouter()
	userMain.HandleFunc("/product-wizard", c.addProduct)
	r.HandleFunc("/get-products-by-user", c.getProductsByUserID)
	r.HandleFunc("/get-projects-by-user", c.getProjectsByUserID)

	r.HandleFunc("/get-categories", c.getCategoriesMap)
	r.HandleFunc("/get-tutorials", c.getTutorials)
	r.HandleFunc("/get-tutorial", c.getTutorial)
	r.HandleFunc("/get-files", c.getFiles)
	r.HandleFunc("/get-files-section", c.getFilesSection)
	r.HandleFunc("/get-news-feed", c.getNewsFeed)
	r.HandleFunc("/get-news-item", c.getNewsItem)
	r.HandleFunc("/get-faqs", c.getFAQs)
	r.HandleFunc("/get-faq", c.getFAQ)
	r.HandleFunc("/get-faq-groups", c.getFAQGroups)
	resources := r.PathPrefix("/resources").Subrouter()
	resources.HandleFunc("/create-news-item", c.addNewsFeed)
	resources.HandleFunc("/edit-news-item", c.updateNewsItem)
	resources.HandleFunc("/create-files-item", c.addFile)
	resources.HandleFunc("/edit-files-item", c.updateFilesItem)
	resources.HandleFunc("/create-tutorial-item", c.addTutorial)
	resources.HandleFunc("/edit-tutorial-item", c.updateTutorial)
	resources.HandleFunc("/create-faq-item", c.addFAQ)
	resources.HandleFunc("/edit-faq-item", c.updateFAQ)
	resources.HandleFunc("/article", c.getArticle)

	r.HandleFunc("/get-request-data", c.getRequestData)

	return r
}
