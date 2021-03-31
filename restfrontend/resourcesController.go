package restfrontend

import (
	"net/http"
)

func (c *RESTFrontend) News(w http.ResponseWriter, r *http.Request) {
	content, err := c.ContentController.BuildNewsContent()
	if err != nil {
		c.HandleError(w, err.Error(), http.StatusInternalServerError, UserMainPath)
		return
	}
	c.RenderTemplate(w, ResourcesNews, content)
}

func (c *RESTFrontend) Docs(w http.ResponseWriter, r *http.Request) {
	content := c.ContentController.BuildDocsContent()
	c.RenderTemplate(w, ResourcesDocs, content)
}

func (c *RESTFrontend) Tutorials(w http.ResponseWriter, r *http.Request) {
	content, err := c.ContentController.BuildTutorialsContent()
	if err != nil {
		c.HandleError(w, err.Error(), http.StatusInternalServerError, UserMainPath)
		return
	}

	c.RenderTemplate(w, ResourcesTutorials, content)
}

func (c *RESTFrontend) FAQ(w http.ResponseWriter, r *http.Request) {
	content, err := c.ContentController.BuildFAQContent()
	if err != nil {
		c.HandleError(w, err.Error(), http.StatusInternalServerError, UserMainPath)
		return
	}
	c.RenderTemplate(w, ResourcesFAQ, content)
}

func (c *RESTFrontend) Files(w http.ResponseWriter, r *http.Request) {
	content, err := c.ContentController.BuildFilesContent()
	if err != nil {
		c.HandleError(w, err.Error(), http.StatusInternalServerError, UserMainPath)
		return
	}
	c.RenderTemplate(w, ResourcesFiles, content)
}
