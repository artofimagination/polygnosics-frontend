package restfrontend

import (
	"net/http"
)

func (c *RESTFrontend) News(w http.ResponseWriter, r *http.Request) {
	content := c.ContentController.BuildNewsContent()
	c.RenderTemplate(w, ResourcesNews, content)
}

func (c *RESTFrontend) Docs(w http.ResponseWriter, r *http.Request) {
	content := c.ContentController.BuildDocsContent()
	c.RenderTemplate(w, ResourcesDocs, content)
}

func (c *RESTFrontend) Tutorials(w http.ResponseWriter, r *http.Request) {
	content := c.ContentController.BuildTutorialsContent()
	c.RenderTemplate(w, ResourcesTutorials, content)
}

func (c *RESTFrontend) FAQ(w http.ResponseWriter, r *http.Request) {
	content := c.ContentController.BuildFAQContent()
	c.RenderTemplate(w, ResourcesFAQ, content)
}

func (c *RESTFrontend) Files(w http.ResponseWriter, r *http.Request) {
	content := c.ContentController.BuildFilesContent()
	c.RenderTemplate(w, ResourcesFiles, content)
}

func (c *RESTFrontend) Instructions(w http.ResponseWriter, r *http.Request) {
	content := make(map[string]interface{})
	c.RenderTemplate(w, ResourcesInstructions, content)
}

func (c *RESTFrontend) Examples(w http.ResponseWriter, r *http.Request) {
	content := make(map[string]interface{})
	c.RenderTemplate(w, ResourcesExamples, content)
}
