package restbackend

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

func (c *RESTBackend) GetTutorials() ([]interface{}, error) {
	data, err := get(BusinessLogicServerAddress, ResourcesURIGetTutorials, "")
	if err != nil {
		return nil, err
	}

	return data.([]interface{}), nil
}

func (c *RESTBackend) GetArticle(r *http.Request) (map[string]interface{}, error) {
	data, err := forwardRequest(BusinessLogicServerAddress, r)
	if err != nil {
		return nil, err
	}

	return data.(map[string]interface{}), nil
}

func (c *RESTBackend) GetFAQs() ([]interface{}, error) {
	data, err := get(BusinessLogicServerAddress, ResourcesURIGetFAQs, "")
	if err != nil {
		return nil, err
	}

	return data.([]interface{}), nil
}

func (c *RESTBackend) GetFAQ(id string) (map[string]interface{}, error) {
	params := fmt.Sprintf("?id=%s", id)
	data, err := get(BusinessLogicServerAddress, ResourcesURIGetFAQ, params)
	if err != nil {
		return nil, err
	}

	return data.(map[string]interface{}), nil
}

func (c *RESTBackend) GetTutorial(id string) (map[string]interface{}, error) {
	params := fmt.Sprintf("?id=%s", id)
	data, err := get(BusinessLogicServerAddress, ResourcesURIGetTutorial, params)
	if err != nil {
		return nil, err
	}

	return data.(map[string]interface{}), nil
}

func (c *RESTBackend) GetFAQGroups() ([]interface{}, error) {
	data, err := get(BusinessLogicServerAddress, ResourcesURIGetFAQGroups, "")
	if err != nil {
		return nil, err
	}

	return data.([]interface{}), nil
}

func (c *RESTBackend) GetFiles() ([]interface{}, error) {
	data, err := get(BusinessLogicServerAddress, ResourcesURIGetFiles, "")
	if err != nil {
		return nil, err
	}

	return data.([]interface{}), nil
}

func (c *RESTBackend) GetNewsFeed() (map[string]interface{}, error) {
	data, err := get(BusinessLogicServerAddress, ResourcesURIGetNewsFeed, "")
	if err != nil {
		return nil, err
	}

	return data.(map[string]interface{}), nil
}

func (c *RESTBackend) GetNewsItem(id string) (map[string]interface{}, error) {
	params := fmt.Sprintf("?id=%s", id)
	data, err := get(BusinessLogicServerAddress, ResourcesURIGetNewsItem, params)
	if err != nil {
		return nil, err
	}

	return data.(map[string]interface{}), nil
}

func (c *RESTBackend) GetFilesSection(id string) (map[string]interface{}, error) {
	params := fmt.Sprintf("?id=%s", id)
	data, err := get(BusinessLogicServerAddress, ResourcesURIGetFilesSection, params)
	if err != nil {
		return nil, err
	}

	return data.(map[string]interface{}), nil
}

func (c *RESTBackend) AddNewsItem(r *http.Request) error {
	_, err := forwardRequest(BusinessLogicServerAddress, r)
	if err != nil {
		return err
	}

	return nil
}

func (c *RESTBackend) AddFileItem(r *http.Request) error {
	_, err := forwardRequest(BusinessLogicServerAddress, r)
	if err != nil {
		return err
	}

	return nil
}

func (c *RESTBackend) AddTutorialItem(r *http.Request) error {
	_, err := forwardRequest(BusinessLogicServerAddress, r)
	if err != nil {
		return err
	}

	return nil
}

func (c *RESTBackend) AddFAQItem(r *http.Request) error {
	_, err := forwardRequest(BusinessLogicServerAddress, r)
	if err != nil {
		return err
	}

	return nil
}

func (c *RESTBackend) UpdateFAQItem(r *http.Request) error {
	_, err := forwardRequest(BusinessLogicServerAddress, r)
	if err != nil {
		return err
	}

	return nil
}

func (c *RESTBackend) UpdateTutorialItem(r *http.Request) error {
	_, err := forwardRequest(BusinessLogicServerAddress, r)
	if err != nil {
		return err
	}

	return nil
}

func (c *RESTBackend) UpdateNewsItem(r *http.Request) error {
	_, err := forwardRequest(BusinessLogicServerAddress, r)
	if err != nil {
		return err
	}

	return nil
}

func (c *RESTBackend) UpdateFilesSection(r *http.Request) error {
	_, err := forwardRequest(BusinessLogicServerAddress, r)
	if err != nil {
		return err
	}

	return nil
}
