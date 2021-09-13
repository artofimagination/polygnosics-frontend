package frontend

import (
	"fmt"
	"net/http"
)

func (c *RESTController) NewNewsResource(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		content := c.ContentController.BuildCreateNews()
		c.RenderTemplate(w, ResourcesCreateNews, content)
	} else {
		if err := c.RESTBackend.AddNewsItem(r); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err.Error())
			return
		}
		c.News(w, r)
	}
}

func (c *RESTController) UpdateNewsResource(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if err := r.ParseForm(); err != nil {
			c.HandleError(w, err.Error(), http.StatusInternalServerError, c.URI(UserMain))
			return
		}
		content, err := c.ContentController.BuildEditNews(r.FormValue("id"))
		if err != nil {
			c.HandleError(w, err.Error(), http.StatusInternalServerError, c.URI(UserMain))
			return
		}
		c.RenderTemplate(w, ResourcesEditNews, content)
	} else {
		if err := c.RESTBackend.UpdateNewsItem(r); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err.Error())
			return
		}
		c.News(w, r)
	}
}

func (c *RESTController) NewFAQResource(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		content, err := c.ContentController.BuildCreateFAQ()
		if err != nil {
			c.HandleError(w, err.Error(), http.StatusInternalServerError, c.URI(UserMain))
			return
		}
		c.RenderTemplate(w, ResourcesCreateFAQ, content)
	} else {
		if err := c.RESTBackend.AddFAQItem(r); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err.Error())
			return
		}
		c.News(w, r)
	}
}

func (c *RESTController) UpdateFAQResource(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if err := r.ParseForm(); err != nil {
			c.HandleError(w, err.Error(), http.StatusInternalServerError, c.URI(UserMain))
			return
		}
		content, err := c.ContentController.BuildEditFAQ(r.FormValue("id"))
		if err != nil {
			c.HandleError(w, err.Error(), http.StatusInternalServerError, c.URI(UserMain))
			return
		}
		c.RenderTemplate(w, ResourcesEditFAQ, content)
	} else {
		if err := c.RESTBackend.UpdateFAQItem(r); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err.Error())
			return
		}
		c.News(w, r)
	}
}

func (c *RESTController) NewFileResource(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		content := c.ContentController.BuildCreateFiles()
		c.RenderTemplate(w, ResourcesCreateFiles, content)
	} else {
		if err := c.RESTBackend.AddFileSection(r); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err.Error())
			return
		}
		c.News(w, r)
	}
}

func (c *RESTController) UpdateFilesResource(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if err := r.ParseForm(); err != nil {
			c.HandleError(w, err.Error(), http.StatusInternalServerError, c.URI(UserMain))
			return
		}
		content, err := c.ContentController.BuildEditFilesSection(r.FormValue("id"))
		if err != nil {
			c.HandleError(w, err.Error(), http.StatusInternalServerError, c.URI(UserMain))
			return
		}
		c.RenderTemplate(w, ResourcesEditFiles, content)
	} else {
		if err := c.RESTBackend.UpdateFilesSection(r); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err.Error())
			return
		}
		c.News(w, r)
	}
}

func (c *RESTController) NewTutorialResource(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		content := c.ContentController.BuildCreateTutorial()
		c.RenderTemplate(w, ResourcesCreateTutorial, content)
	} else {
		if err := c.RESTBackend.AddTutorialItem(r); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err.Error())
			return
		}
		c.News(w, r)
	}
}

func (c *RESTController) UpdateTutorialResource(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if err := r.ParseForm(); err != nil {
			c.HandleError(w, err.Error(), http.StatusInternalServerError, c.URI(UserMain))
			return
		}
		content, err := c.ContentController.BuildEditTutorial(r.FormValue("id"))
		if err != nil {
			c.HandleError(w, err.Error(), http.StatusInternalServerError, c.URI(UserMain))
			return
		}
		c.RenderTemplate(w, ResourcesEditTutorial, content)
	} else {
		if err := c.RESTBackend.UpdateTutorialItem(r); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err.Error())
			return
		}
		c.News(w, r)
	}
}

func (c *RESTController) News(w http.ResponseWriter, r *http.Request) {
	content, err := c.ContentController.BuildNewsContent()
	if err != nil {
		c.HandleError(w, err.Error(), http.StatusInternalServerError, c.URI(UserMain))
		return
	}
	c.RenderTemplate(w, ResourcesNews, content)
}

func (c *RESTController) Docs(w http.ResponseWriter, r *http.Request) {
	content := c.ContentController.BuildDocsContent()
	c.RenderTemplate(w, ResourcesDocs, content)
}

func (c *RESTController) Article(w http.ResponseWriter, r *http.Request) {
	content, err := c.ContentController.BuildArticleContent(r)
	if err != nil {
		c.HandleError(w, err.Error(), http.StatusInternalServerError, c.URI(UserMain))
		return
	}
	c.RenderTemplate(w, ResourcesArticle, content)
}

func (c *RESTController) Tutorials(w http.ResponseWriter, r *http.Request) {
	content, err := c.ContentController.BuildTutorialsContent()
	if err != nil {
		c.HandleError(w, err.Error(), http.StatusInternalServerError, c.URI(UserMain))
		return
	}

	c.RenderTemplate(w, ResourcesTutorials, content)
}

func (c *RESTController) FAQ(w http.ResponseWriter, r *http.Request) {
	content, err := c.ContentController.BuildFAQContent()
	if err != nil {
		c.HandleError(w, err.Error(), http.StatusInternalServerError, c.URI(UserMain))
		return
	}
	c.RenderTemplate(w, ResourcesFAQ, content)
}

func (c *RESTController) Files(w http.ResponseWriter, r *http.Request) {
	content, err := c.ContentController.BuildFilesContent()
	if err != nil {
		c.HandleError(w, err.Error(), http.StatusInternalServerError, c.URI(UserMain))
		return
	}
	c.RenderTemplate(w, ResourcesFiles, content)
}
