package contents

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	TutorialsKey        = "tutorials"
	TutorialsContentKey = "content"
	TutorialsRefKey     = "content_ref"

	FAQsKey         = "faqs"
	FAQGroupsKey    = "faq_groups"
	FAQsGroupKey    = "group"
	FAQsQuestionKey = "question"
	FAQsAnswerKey   = "answer"

	EditKey = "edit"
	IDKey   = "id"

	FilesKey = "files"

	NewsFeedKey      = "news_feed"
	NewsFeedTextKey  = "news_text"
	NewsFeedYearKey  = "news_year"
	NewsFeedMonthKey = "news_month"
	NewsFeedDayKey   = "news_day"

	CreateItemKey   = "create_item"
	ResourceContent = "resource_content"
)

// parseContent parses formatted text stored in binary files
func parseContent(path string) (string, error) {
	if path == "" {
		return "", nil
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (c *ContentController) prepareTutorials(content map[string]interface{}) error {
	c.prepareNewTutorial(content)
	content[TutorialsKey] = make([]interface{}, 0)
	tutorials, err := c.RESTBackend.GetTutorials()
	if err != nil {
		return err
	}
	for _, tutorial := range tutorials {
		if val, ok := tutorial.(map[string]interface{})[TutorialsContentKey]; ok {
			text, err := parseContent(val.(string))
			if err != nil {
				return err
			}
			tutorial.(map[string]interface{})[TutorialsContentKey] = text
			id := tutorial.(map[string]interface{})[IDKey].(string)
			tutorial.(map[string]interface{})[TutorialsRefKey] = fmt.Sprintf("/resources/article?id=%s", id)
			tutorial.(map[string]interface{})[EditKey] = fmt.Sprintf("/resources/edit-tutorial-item?id=%s", tutorial.(map[string]interface{})[IDKey])
			content[TutorialsKey] = append(content[TutorialsKey].([]interface{}), tutorial)
		}
	}

	return nil
}

func extractFAQData(faq map[string]interface{}) error {
	if val, ok := faq[FAQsQuestionKey]; ok {
		text, err := parseContent(val.(string))
		if err != nil {
			return err
		}
		faq[FAQsQuestionKey] = text
	}
	if val, ok := faq[FAQsAnswerKey]; ok {
		text, err := parseContent(val.(string))
		if err != nil {
			return err
		}
		faq[FAQsAnswerKey] = text
	}
	return nil
}

func (c *ContentController) prepareFAQ(content map[string]interface{}) error {
	if err := c.prepareNewFAQ(content); err != nil {
		return err
	}
	content[FAQsKey] = make([]interface{}, 0)

	faqs, err := c.RESTBackend.GetFAQs()
	if err != nil {
		return err
	}
	for _, faq := range faqs {
		if err := extractFAQData(faq.(map[string]interface{})); err != nil {
			return err
		}
		faq.(map[string]interface{})[EditKey] = fmt.Sprintf("/resources/edit-faq-item?id=%s", faq.(map[string]interface{})[IDKey])
		content[FAQsKey] = append(content[FAQsKey].([]interface{}), faq)
	}

	return nil
}

// prepareNewsFeed creates the UI content from DB content
// Structure:
//		Content->News content map
//									-> Year key
//												-> Post key
//													-> Month
//													-> Day
//													-> Text
func (c *ContentController) prepareNewsFeed(content map[string]interface{}) error {
	c.prepareCreateNewsFeed(content)
	news, err := c.RESTBackend.GetNewsFeed()
	if err != nil {
		return err
	}

	newsMap := make(map[string]interface{})
	for _, entry := range news {
		year := ""
		if val, ok := entry.(map[string]interface{})[NewsFeedYearKey]; ok {
			year = val.(string)
		} else {
			return errors.New("Years key missing")
		}

		if _, ok := newsMap[year]; !ok {
			newsMap[year] = make([]interface{}, 0)
		}

		if val, ok := entry.(map[string]interface{})[NewsFeedTextKey]; ok {
			text, err := parseContent(val.(string))
			if err != nil {
				return err
			}
			newsEntry := make(map[string]interface{})

			newsEntry[NewsFeedTextKey] = text
			newsEntry[NewsFeedMonthKey] = entry.(map[string]interface{})[NewsFeedMonthKey]
			newsEntry[NewsFeedDayKey] = entry.(map[string]interface{})[NewsFeedDayKey]
			newsEntry[EditKey] = fmt.Sprintf("/resources/edit-news-item?id=%s", entry.(map[string]interface{})[IDKey])
			newsMap[year] = append(newsMap[year].([]interface{}), newsEntry)
		}
	}
	content[NewsFeedKey] = newsMap
	return nil
}

func (c *ContentController) prepareFiles(content map[string]interface{}) error {
	c.prepareNewFiles(content)
	files, err := c.RESTBackend.GetFiles()
	if err != nil {
		return err
	}
	for _, v := range files {
		v.(map[string]interface{})[EditKey] = fmt.Sprintf("/resources/edit-files-item?id=%s", v.(map[string]interface{})[IDKey])
	}
	content[FilesKey] = files
	return nil
}

func (c *ContentController) prepareEditFiles(id string, content map[string]interface{}) error {
	c.prepareNewFiles(content)
	fileSection, err := c.RESTBackend.GetFilesSection(id)
	if err != nil {
		return err
	}

	fileSection[EditKey] = fmt.Sprintf("/resources/edit-files-item?id=%s", fileSection[IDKey])
	for _, v := range fileSection["files"].([]interface{}) {
		if v.(map[string]interface{})["type"] == "github" ||
			v.(map[string]interface{})["type"] == "gitlab" ||
			v.(map[string]interface{})["type"] == "bitbucket" {
			v.(map[string]interface{})["type"] = "link"
		}
	}
	content["files"] = fileSection
	return nil
}

func (c *ContentController) prepareNewFiles(content map[string]interface{}) {
	content[CreateItemKey] = "/resources/create-files-item"
}

func (c *ContentController) prepareEditNews(id string, content map[string]interface{}) error {
	c.prepareCreateNewsFeed(content)
	news, err := c.RESTBackend.GetNewsItem(id)
	if err != nil {
		return err
	}
	text, err := parseContent(news[NewsFeedTextKey].(string))
	if err != nil {
		return err
	}
	news[NewsFeedTextKey] = text
	news[EditKey] = fmt.Sprintf("/resources/edit-news-item?id=%s", news[IDKey])
	content["news"] = news
	return nil
}

func (c *ContentController) prepareNewTutorial(content map[string]interface{}) {
	content[CreateItemKey] = "/resources/create-tutorial-item"
}

func (c *ContentController) prepareEditTutorial(id string, content map[string]interface{}) error {
	c.prepareNewTutorial(content)
	tutorial, err := c.RESTBackend.GetTutorial(id)
	if err != nil {
		return err
	}
	text, err := parseContent(tutorial[TutorialsContentKey].(string))
	if err != nil {
		return err
	}
	tutorial[TutorialsContentKey] = text
	tutorial[EditKey] = fmt.Sprintf("/resources/edit-tutorial-item?id=%s", tutorial[IDKey])
	content["tutorial"] = tutorial
	return nil
}

func (c *ContentController) prepareNewFAQ(content map[string]interface{}) error {
	content[ResourceContent] = ""
	content[CreateItemKey] = "/resources/create-faq-item"
	groups, err := c.RESTBackend.GetFAQGroups()
	if err != nil {
		return err
	}
	content[FAQGroupsKey] = groups
	return nil
}

func (c *ContentController) prepareEditFAQ(id string, content map[string]interface{}) error {
	if err := c.prepareNewFAQ(content); err != nil {
		return err
	}
	faq, err := c.RESTBackend.GetFAQ(id)
	if err != nil {
		return err
	}
	if err := extractFAQData(faq); err != nil {
		return err
	}
	faq[EditKey] = fmt.Sprintf("/resources/edit-faq-item?id=%s", faq[IDKey])
	content["faq"] = faq
	return nil
}

func (c *ContentController) prepareCreateNewsFeed(content map[string]interface{}) {
	content[CreateItemKey] = "/resources/create-news-item"
}

func (c *ContentController) prepareArticle(content map[string]interface{}, r *http.Request) error {
	article, err := c.RESTBackend.GetArticle(r)
	if err != nil {
		return err
	}
	text, err := parseContent(article["content"].(string))
	if err != nil {
		return err
	}
	article["content"] = text
	content["article"] = article
	return nil
}
