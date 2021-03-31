package restbackend

const (
	ResourcesURIGetTutorials = "/get-tutorials"
	ResourcesURIGetFAQ       = "/get-faq"
	ResourcesURIGetNewsFeed  = "/get-news-feed"
	ResourcesURIGetFAQGroups = "/get-faq-groups"
	ResourcesURIGetFiles     = "/get-files"
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
