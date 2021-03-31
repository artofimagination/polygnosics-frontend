package contents

import (
	"io/ioutil"
)

const (
	TutorialsKey        = "tutorials"
	TutorialsContentKey = "content"

	FAQGroupsKey    = "faq_groups"
	FAQsKey         = "faqs"
	FAQsGroupKey    = "group"
	FAQsQuestionKey = "question"
	FAQsAnswerKey   = "answer"

	FilesKey = "files"

	NewsFeedKey     = "news_feed"
	NewsFeedTextKey = "news_text"
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
			content[TutorialsKey] = append(content[TutorialsKey].([]interface{}), tutorial)
		}
	}

	return nil
}

func (c *ContentController) prepareFAQ(content map[string]interface{}) error {
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
	files, err := c.RESTBackend.GetFiles()
	if err != nil {
		return err
	}
	content[FilesKey] = files
	return nil
}
