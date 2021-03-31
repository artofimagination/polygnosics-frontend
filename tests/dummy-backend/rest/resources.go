package rest

import (
	"fmt"
	"net/http"
)

const (
	TutorialsKey = "tutorials"
	FilesKey     = "files"
	NewsFeedKey  = "news_feed"
	FAQGroupsKey = "faq_groups"
	FAQsKey      = "faqs"
)

func (c *Controller) getTutorials(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	if err := c.ParseForm(r); err != nil {
		writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusBadRequest)
		return
	}

	data["data"] = c.TestData[TutorialsKey]
	encodeResponse(data, w)
}

func (c *Controller) getFAQs(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	if err := c.ParseForm(r); err != nil {
		writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusBadRequest)
		return
	}

	data["data"] = c.TestData[FAQsKey]
	encodeResponse(data, w)
}

func (c *Controller) getFiles(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	if err := c.ParseForm(r); err != nil {
		writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusBadRequest)
		return
	}

	data["data"] = c.TestData[FilesKey]
	encodeResponse(data, w)
}

func (c *Controller) getNewsFeed(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	if err := c.ParseForm(r); err != nil {
		writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusBadRequest)
		return
	}

	data["data"] = c.TestData[NewsFeedKey]
	encodeResponse(data, w)
}

func (c *Controller) getFAQGroups(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	if err := c.ParseForm(r); err != nil {
		writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusBadRequest)
		return
	}

	data["data"] = c.TestData[FAQGroupsKey]
	encodeResponse(data, w)
}
