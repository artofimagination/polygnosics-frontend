package rest

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
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

	data["data"] = c.TestData[TutorialsKey]
	encodeResponse(data, w)
}

func (c *Controller) getArticle(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	if err := c.ParseForm(r); err != nil {
		writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusBadRequest)
		return
	}

	id := r.FormValue("id")
	for _, v := range c.TestData[TutorialsKey].([]interface{}) {
		if v.(map[string]interface{})["id"] == id {
			data["data"] = v
			break
		}
	}
	encodeResponse(data, w)
}

func (c *Controller) getFAQs(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})

	data["data"] = c.TestData[FAQsKey]
	encodeResponse(data, w)
}

func (c *Controller) getFAQ(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	if err := c.ParseForm(r); err != nil {
		writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusBadRequest)
		return
	}

	id := r.FormValue("id")
	for _, v := range c.TestData[FAQsKey].([]interface{}) {
		if v.(map[string]interface{})["id"] == id {
			data["data"] = v
			break
		}
	}

	encodeResponse(data, w)
}

func (c *Controller) getFilesSection(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	if err := c.ParseForm(r); err != nil {
		writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusBadRequest)
		return
	}

	id := r.FormValue("id")
	for _, v := range c.TestData[FilesKey].([]interface{}) {
		if v.(map[string]interface{})["id"] == id {
			data["data"] = v
			break
		}
	}

	encodeResponse(data, w)
}

func (c *Controller) addFAQ(w http.ResponseWriter, r *http.Request) {
	if err := c.ParseForm(r); err != nil {
		writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusBadRequest)
		return
	}

	if err := WriteToFile("/user-assets/new-faq-answer-entry.txt", r.FormValue("answer")); err != nil {
		writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusInternalServerError)
		return
	}

	if err := WriteToFile("/user-assets/new-faq-question-entry.txt", r.FormValue("question")); err != nil {
		writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusInternalServerError)
		return
	}

	data := make(map[string]interface{})
	data["id"] = uuid.New().String()
	data["group"] = r.FormValue("group")
	data["question"] = "/user-assets/new-faq-question-entry.txt"
	data["answer"] = "/user-assets/new-faq-answer-entry.txt"

	c.TestData["faqs"] = append(c.TestData["faqs"].([]interface{}), data)
	writeData("OK", w, http.StatusCreated)
}

func (c *Controller) updateFAQ(w http.ResponseWriter, r *http.Request) {
	if err := c.ParseForm(r); err != nil {
		writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusBadRequest)
		return
	}

	for _, v := range c.TestData[FAQsKey].([]interface{}) {
		if v.(map[string]interface{})["id"] == r.FormValue("id") {
			if err := WriteToFile(v.(map[string]interface{})["answer"].(string), r.FormValue("answer")); err != nil {
				writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusInternalServerError)
				return
			}

			if err := WriteToFile(v.(map[string]interface{})["question"].(string), r.FormValue("question")); err != nil {
				writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusInternalServerError)
				return
			}
			v.(map[string]interface{})["group"] = r.FormValue("group")
			break
		}
	}

	writeData("OK", w, http.StatusCreated)
}

func (c *Controller) addTutorial(w http.ResponseWriter, r *http.Request) {
	if err := c.ParseMultipartForm(r); err != nil {
		writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusBadRequest)
		return
	}

	data := make(map[string]interface{})
	data["id"] = uuid.New().String()
	data["avatar_type"] = r.FormValue("avatar_type")
	if data["avatar_type"] == "image" {
		data["avatar"] = "/user-assets/uploads/new-tutorial-image.jpg"
		if err := uploadFile("avatar_image", "new-tutorial-image.jpg", r); err != nil {
			writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusInternalServerError)
			return
		}
	} else {
		data["avatar"] = r.FormValue("avatar_video")
	}
	data["title"] = r.FormValue("title")
	data["short"] = r.FormValue("short")
	data["last_updated"] = "today"
	data["content"] = ""
	if r.FormValue("article") != "" {
		data["content"] = "/user-assets/new-tutorial-entry.txt"
		if err := WriteToFile("/user-assets/new-tutorial-entry.txt", r.FormValue("article")); err != nil {
			writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusInternalServerError)
			return
		}
	}

	c.TestData["tutorials"] = append(c.TestData["tutorials"].([]interface{}), data)
	writeData("OK", w, http.StatusCreated)
}

func (c *Controller) getTutorial(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	if err := c.ParseForm(r); err != nil {
		writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusBadRequest)
		return
	}

	id := r.FormValue("id")
	for _, v := range c.TestData[TutorialsKey].([]interface{}) {
		if v.(map[string]interface{})["id"] == id {
			data["data"] = v
			break
		}
	}

	encodeResponse(data, w)
}

func (c *Controller) updateTutorial(w http.ResponseWriter, r *http.Request) {
	if err := c.ParseMultipartForm(r); err != nil {
		writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusBadRequest)
		return
	}

	for _, v := range c.TestData[TutorialsKey].([]interface{}) {
		if v.(map[string]interface{})["id"] == r.FormValue("id") {
			if v.(map[string]interface{})["content"] == "" && r.FormValue("article") != "" {
				v.(map[string]interface{})["content"] = "/user-assets/update-tutorial-entry.txt"
			}
			if err := WriteToFile(v.(map[string]interface{})["content"].(string), r.FormValue("article")); err != nil {
				writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusInternalServerError)
				return
			}
			v.(map[string]interface{})["avatar_type"] = r.FormValue("avatar_type")
			if r.FormValue("avatar_type") == "image" {
				v.(map[string]interface{})["avatar"] = "/user-assets/uploads/new-tutorial-image.jpg"
				if err := uploadFile("avatar_image", "new-tutorial-image.jpg", r); err != nil {
					writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusInternalServerError)
					return
				}
			} else {
				v.(map[string]interface{})["avatar"] = r.FormValue("avatar_video")
			}
			v.(map[string]interface{})["title"] = r.FormValue("title")
			v.(map[string]interface{})["short"] = r.FormValue("short")
			v.(map[string]interface{})["last_updated"] = "today"
			break
		}
	}

	writeData("OK", w, http.StatusCreated)
}

func (c *Controller) addFile(w http.ResponseWriter, r *http.Request) {
	if err := c.ParseMultipartForm(r); err != nil {
		writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusBadRequest)
		return
	}
	data := make(map[string]interface{})
	data["id"] = uuid.New().String()
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
			fileName := fmt.Sprintf("new-file_%d.txt", i)
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

	writeData("OK", w, http.StatusCreated)
}

func (c *Controller) updateFilesItem(w http.ResponseWriter, r *http.Request) {
	if err := c.ParseMultipartForm(r); err != nil {
		writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusBadRequest)
		return
	}

	count, err := strconv.Atoi(r.FormValue("count"))
	if err != nil {
		writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusBadRequest)
		return
	}

	for _, v := range c.TestData[FilesKey].([]interface{}) {
		if v.(map[string]interface{})["id"] == r.FormValue("id") {
			v.(map[string]interface{})["title"] = r.FormValue("title")
			v.(map[string]interface{})["short"] = r.FormValue("short")
			v.(map[string]interface{})["files"] = make([]interface{}, 0)

			for i := 0; i <= 50 && i < count; i++ {
				file := make(map[string]interface{})

				file["type"] = r.FormValue(fmt.Sprintf("type_%d", i))
				if file["type"] == "file" {
					formName := fmt.Sprintf("upload_file_%d", i)
					file["ref"] = r.FormValue(formName)
					fileName := fmt.Sprintf("new-file_%d.txt", i)
					if err := uploadFile(formName, fileName, r); err != nil {
						writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusInternalServerError)
						return
					}
				} else {
					file["ref"] = r.FormValue(fmt.Sprintf("repo_link_%d", i))
				}
				file["ref_name"] = r.FormValue(fmt.Sprintf("ref_name_%d", i))
				v.(map[string]interface{})["files"] = append(v.(map[string]interface{})["files"].([]interface{}), file)
			}
			break
		}
	}

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
	newsItem["id"] = uuid.New().String()
	newsItem[NewsTextKey] = "/user-assets/new-news-entry.txt"
	newsItem[NewsDayKey] = day
	newsItem[NewsMonthKey] = dt.Month().String()[0:3]
	newsList = append(newsList, newsItem)

	c.TestData[NewsFeedKey].(map[string]interface{})[year] = newsList
	writeData("OK", w, http.StatusCreated)
}

func (c *Controller) getNewsItem(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	if err := c.ParseForm(r); err != nil {
		writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusBadRequest)
		return
	}

	id := r.FormValue("id")
	for _, v := range c.TestData[NewsFeedKey].(map[string]interface{}) {
		for _, news := range v.([]interface{}) {
			if news.(map[string]interface{})["id"] == id {
				data["data"] = news
				break
			}
		}
	}

	encodeResponse(data, w)
}

func (c *Controller) updateNewsItem(w http.ResponseWriter, r *http.Request) {
	if err := c.ParseForm(r); err != nil {
		writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusBadRequest)
		return
	}

	for _, v := range c.TestData[NewsFeedKey].(map[string]interface{}) {
		for _, news := range v.([]interface{}) {
			if news.(map[string]interface{})["id"] == r.FormValue("id") {
				if err := WriteToFile(news.(map[string]interface{})["news_text"].(string), r.FormValue("news")); err != nil {
					writeError(fmt.Sprintf("Backend: %s", err.Error()), w, http.StatusInternalServerError)
					return
				}
				break
			}
		}
	}

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
