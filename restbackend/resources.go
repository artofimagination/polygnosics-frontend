package restbackend

import (
	"net/http"
)

const (
	ResourcesURIGetTutorials = "/get-tutorials"
	ResourcesURIAddTutorial  = "/add-tutorial"
	ResourcesURIGetFAQ       = "/get-faq"
	ResourcesURIAddFAQ       = "/add-faq"
	ResourcesURIGetNewsFeed  = "/get-news-feed"
	ResourcesURIAddNewsFeed  = "/add-news-feed"
	ResourcesURIGetFAQGroups = "/get-faq-groups"
	ResourcesURIGetFiles     = "/get-files"
	ResourcesURIAddFile      = "/add-file"
)

func (c *RESTBackend) GetTutorials() ([]interface{}, error) {
	data, err := get(BusinessLogicServerAddress, ResourcesURIGetTutorials, "")
	if err != nil {
		return nil, err
	}

	return data.([]interface{}), nil
}

func (c *RESTBackend) GetFAQ() ([]interface{}, error) {
	data, err := get(BusinessLogicServerAddress, ResourcesURIGetFAQ, "")
	if err != nil {
		return nil, err
	}

	return data.([]interface{}), nil
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

func (c *RESTBackend) AddNewsItem(r *http.Request) error {
	err := forwardRequest(BusinessLogicServerAddress, r)
	if err != nil {
		return err
	}

	return nil
}

func (c *RESTBackend) AddFileItem(r *http.Request) error {
	err := forwardRequest(BusinessLogicServerAddress, r)
	if err != nil {
		return err
	}

	return nil
}

func (c *RESTBackend) AddTutorialItem(r *http.Request) error {
	err := forwardRequest(BusinessLogicServerAddress, r)
	if err != nil {
		return err
	}

	return nil
}

func (c *RESTBackend) AddFAQItem(r *http.Request) error {
	err := forwardRequest(BusinessLogicServerAddress, r)
	if err != nil {
		return err
	}

	return nil
}
