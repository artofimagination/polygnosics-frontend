package rest

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const (
	TutorialsKey = "tutorials"
	FilesKey     = "files"
	NewsFeedKey  = "news_feed"
	FAQGroupsKey = "faq_groups"
	FAQsKey      = "faqs"
)

const (
	NewsMonthKey = "news_month"
	NewsDayKey   = "news_day"
	NewsTextKey  = "news_text"
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

func (c *Controller) addFAQ(w http.ResponseWriter, r *http.Request) {
	requestData, err := c.decodeRequest(r)
	if err != nil {
		writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusBadRequest)
		return
	}

	c.TestData["faqs"] = append(c.TestData["faqs"].([]interface{}), requestData)
	prettyPrint(c.TestData)
	writeData("OK", w, http.StatusCreated)
}

func (c *Controller) addTutorial(w http.ResponseWriter, r *http.Request) {
	requestData, err := c.decodeRequest(r)
	if err != nil {
		writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusBadRequest)
		return
	}

	c.TestData["tutorials"] = append(c.TestData["tutorials"].([]interface{}), requestData)
	prettyPrint(c.TestData)
	writeData("OK", w, http.StatusCreated)
}

func (c *Controller) addFile(w http.ResponseWriter, r *http.Request) {
	if err := c.ParseMultipartForm(r); err != nil {
		writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusBadRequest)
		return
	}
	data := make(map[string]interface{})
	data["title"] = r.FormValue("title")
	data["short"] = r.FormValue("short")
	data["files"] = make([]interface{}, 0)
	count, err := strconv.Atoi(r.FormValue("count"))
	if err != nil {
		writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusBadRequest)
		return
	}

	for i := 0; i <= 50 && i < count; i++ {
		file := make(map[string]interface{})

		file["type"] = r.FormValue(fmt.Sprintf("type_%d", i))
		if file["type"] == "file" {
			formName := fmt.Sprintf("upload_file_%d", i)
			file["ref"] = r.FormValue(formName)
			fileName := fmt.Sprintf("/user-assets/uploads/new-file_%d.txt", i)
			if err := uploadFile(formName, fileName, r); err != nil {
				writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusInternalServerError)
				return
			}
		} else {
			file["ref"] = r.FormValue(fmt.Sprintf("repo_link_%d", i))
		}
		file["ref_name"] = r.FormValue(fmt.Sprintf("ref_name_%d", i))
		data["files"] = append(data["files"].([]interface{}), file)
	}

	c.TestData["files"] = append(c.TestData["files"].([]interface{}), data)

	prettyPrint(c.TestData)
	writeData("OK", w, http.StatusCreated)
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

func (c *Controller) addNewsFeed(w http.ResponseWriter, r *http.Request) {
	if err := c.ParseForm(r); err != nil {
		writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusBadRequest)
		return
	}

	newsList := make([]interface{}, 0)
	newsItem := make(map[string]interface{})
	dt, err := time.Parse("Mon Jan 02 2006 15:04:05.0000 GMT-0700", "Mon Jan 02 2006 15:04:05.0000 GMT-0700")
	if err != nil {
		writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusInternalServerError)
		return
	}
	year := strconv.Itoa(dt.Year())
	day := strconv.Itoa(dt.Day())

	err = WriteToFile("/user-assets/new-news-entry.txt", r.FormValue("news"))
	if err != nil {
		writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusInternalServerError)
		return
	}
	newsItem[NewsTextKey] = "/user-assets/new-news-entry.txt"
	newsItem[NewsDayKey] = day
	newsItem[NewsMonthKey] = dt.Month().String()[0:3]
	newsList = append(newsList, newsItem)

	c.TestData[NewsFeedKey].(map[string]interface{})[year] = newsList
	prettyPrint(c.TestData)
	writeData("OK", w, http.StatusCreated)
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
