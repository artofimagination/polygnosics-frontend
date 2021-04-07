package contents

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	TutorialsKey        = "tutorials"
	TutorialsContentKey = "content"
	TutorialsRefKey     = "content_ref"
	TutorialsIDKey      = "id"

	FAQsKey         = "faqs"
	FAQGroupsKey    = "faq_groups"
	FAQsGroupKey    = "group"
	FAQsQuestionKey = "question"
	FAQsAnswerKey   = "answer"

	FilesKey = "files"

	NewsFeedKey     = "news_feed"
	NewsFeedTextKey = "news_text"

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
	content[CreateItemKey] = "/resources/create-tutorial-item"
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
			id := tutorial.(map[string]interface{})[TutorialsIDKey].(string)
			tutorial.(map[string]interface{})[TutorialsRefKey] = fmt.Sprintf("/resources/article?id=%s", id)
			content[TutorialsKey] = append(content[TutorialsKey].([]interface{}), tutorial)
		}
	}

	return nil
}

func (c *ContentController) prepareFAQ(content map[string]interface{}) error {
	content[CreateItemKey] = "/resources/create-faq-item"
	content[FAQsKey] = make([]interface{}, 0)
	faqGroups, err := c.RESTBackend.GetFAQGroups()
	if err != nil {
		return err
	}
	content[FAQGroupsKey] = faqGroups

	faqs, err := c.RESTBackend.GetFAQ()
	if err != nil {
		return err
	}
	for _, faq := range faqs {
		if val, ok := faq.(map[string]interface{})[FAQsQuestionKey]; ok {
			text, err := parseContent(val.(string))
			if err != nil {
				return err
			}
			faq.(map[string]interface{})[FAQsQuestionKey] = text
		}
		if val, ok := faq.(map[string]interface{})[FAQsAnswerKey]; ok {
			text, err := parseContent(val.(string))
			if err != nil {
				return err
			}
			faq.(map[string]interface{})[FAQsAnswerKey] = text
		}
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
	content[CreateItemKey] = "/resources/create-news-item"
	news, err := c.RESTBackend.GetNewsFeed()
	if err != nil {
		return err
	}
	for year, newsYearItem := range news {
		for index, newsItem := range newsYearItem.([]interface{}) {
			if val, ok := newsItem.(map[string]interface{})[NewsFeedTextKey]; ok {
				text, err := parseContent(val.(string))
				if err != nil {
					return err
				}
				news[year].([]interface{})[index].(map[string]interface{})[NewsFeedTextKey] = text
			}
		}
	}
	content[NewsFeedKey] = news
	return nil
}

func (c *ContentController) prepareFiles(content map[string]interface{}) error {
	content[CreateItemKey] = "/resources/create-files-item"
	files, err := c.RESTBackend.GetFiles()
	if err != nil {
		return err
	}
	content[FilesKey] = files
	return nil
}

func (c *ContentController) prepareNewFiles(content map[string]interface{}) {
	content[ResourceContent] = ""
	content[CreateItemKey] = "/resources/create-files-item"
}

func (c *ContentController) prepareNewTutorial(content map[string]interface{}) {
	content[ResourceContent] = ""
	content[CreateItemKey] = "/resources/create-tutorial-item"
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

func (c *ContentController) prepareNewNewsFeed(content map[string]interface{}) {
	content[ResourceContent] = ""
	content[CreateItemKey] = "/resources/create-news-item"
}

func (c *ContentController) prepareArticle(content map[string]interface{}, r *http.Request) error {
	content[ResourceContent] = ""
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
