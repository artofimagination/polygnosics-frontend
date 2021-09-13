package backend

import (
	"fmt"
	"net/http"
)

const (
	ResourcesURIGetTutorials    = "/get-tutorials"
	ResourcesURIAddTutorial     = "/add-tutorial"
	ResourcesURIGetTutorial     = "/get-tutorial"
	ResourcesURIGetFAQs         = "/get-faqs"
	ResourcesURIGetFAQ          = "/get-faq"
	ResourcesURIAddFAQ          = "/add-faq"
	ResourcesURIGetNewsFeed     = "/get-news-feed"
	ResourcesURIAddNewsFeed     = "/add-news-feed"
	ResourcesURIGetNewsItem     = "/get-news-item"
	ResourcesURIGetFAQGroups    = "/get-faq-groups"
	ResourcesURIGetFiles        = "/get-files"
	ResourcesURIGetFilesSection = "/get-files-section"
	ResourcesURIAddFile         = "/add-file"
)

func (c *RESTController) GetTutorials() ([]interface{}, error) {
	data, err := get(BusinessLogicServerAddress, ResourcesURIGetTutorials, "")
	if err != nil {
		return nil, err
	}

	return data.([]interface{}), nil
}

func (c *RESTController) GetArticle(r *http.Request) (map[string]interface{}, error) {
	data, err := forwardRequest(BusinessLogicServerAddress, r)
	if err != nil {
		return nil, err
	}

	return data.(map[string]interface{}), nil
}

func (c *RESTController) GetFAQs() ([]interface{}, error) {
	data, err := get(BusinessLogicServerAddress, ResourcesURIGetFAQs, "")
	if err != nil {
		return nil, err
	}

	return data.([]interface{}), nil
}

func (c *RESTController) GetFAQ(id string) (map[string]interface{}, error) {
	params := fmt.Sprintf("?id=%s", id)
	data, err := get(BusinessLogicServerAddress, ResourcesURIGetFAQ, params)
	if err != nil {
		return nil, err
	}

	return data.(map[string]interface{}), nil
}

func (c *RESTController) GetTutorial(id string) (map[string]interface{}, error) {
	params := fmt.Sprintf("?id=%s", id)
	data, err := get(BusinessLogicServerAddress, ResourcesURIGetTutorial, params)
	if err != nil {
		return nil, err
	}

	return data.(map[string]interface{}), nil
}

func (c *RESTController) GetFAQGroups() ([]interface{}, error) {
	data, err := get(BusinessLogicServerAddress, ResourcesURIGetFAQGroups, "")
	if err != nil {
		return nil, err
	}

	return data.([]interface{}), nil
}

func (c *RESTController) GetFiles() ([]interface{}, error) {
	data, err := get(BusinessLogicServerAddress, ResourcesURIGetFiles, "")
	if err != nil {
		return nil, err
	}

	return data.([]interface{}), nil
}

func (c *RESTController) GetNewsFeed() ([]interface{}, error) {
	data, err := get(BusinessLogicServerAddress, ResourcesURIGetNewsFeed, "")
	if err != nil {
		return nil, err
	}

	return data.([]interface{}), nil
}

func (c *RESTController) GetNewsItem(id string) (map[string]interface{}, error) {
	params := fmt.Sprintf("?id=%s", id)
	data, err := get(BusinessLogicServerAddress, ResourcesURIGetNewsItem, params)
	if err != nil {
		return nil, err
	}

	return data.(map[string]interface{}), nil
}

func (c *RESTController) GetFilesSection(id string) (map[string]interface{}, error) {
	params := fmt.Sprintf("?id=%s", id)
	data, err := get(BusinessLogicServerAddress, ResourcesURIGetFilesSection, params)
	if err != nil {
		return nil, err
	}

	return data.(map[string]interface{}), nil
}

func (c *RESTController) AddNewsItem(r *http.Request) error {
	_, err := forwardRequest(BusinessLogicServerAddress, r)
	if err != nil {
		return err
	}

	return nil
}

func (c *RESTController) AddFileSection(r *http.Request) error {
	_, err := forwardRequest(BusinessLogicServerAddress, r)
	if err != nil {
		return err
	}

	return nil
}

func (c *RESTController) AddTutorialItem(r *http.Request) error {
	_, err := forwardRequest(BusinessLogicServerAddress, r)
	if err != nil {
		return err
	}

	return nil
}

func (c *RESTController) AddFAQItem(r *http.Request) error {
	_, err := forwardRequest(BusinessLogicServerAddress, r)
	if err != nil {
		return err
	}

	return nil
}

func (c *RESTController) UpdateFAQItem(r *http.Request) error {
	_, err := forwardRequest(BusinessLogicServerAddress, r)
	if err != nil {
		return err
	}

	return nil
}

func (c *RESTController) UpdateTutorialItem(r *http.Request) error {
	_, err := forwardRequest(BusinessLogicServerAddress, r)
	if err != nil {
		return err
	}

	return nil
}

func (c *RESTController) UpdateNewsItem(r *http.Request) error {
	_, err := forwardRequest(BusinessLogicServerAddress, r)
	if err != nil {
		return err
	}

	return nil
}

func (c *RESTController) UpdateFilesSection(r *http.Request) error {
	_, err := forwardRequest(BusinessLogicServerAddress, r)
	if err != nil {
		return err
	}

	return nil
}
