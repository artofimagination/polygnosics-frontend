package restfrontend

import (
	"fmt"
	"net/http"

	"github.com/artofimagination/polygnosics-frontend/contents"

	"github.com/pkg/errors"
)

func (c *RESTFrontend) MyProjects(w http.ResponseWriter, r *http.Request) {
	content, err := c.ContentController.BuildMyProjectsContent()
	if err != nil {
		c.HandleError(w, fmt.Sprintf("Failed to get project content. %s", errors.WithStack(err)), http.StatusInternalServerError, UserMainPath)
		return
	}
	c.RenderTemplate(w, MyProjects, content)
}

func (c *RESTFrontend) ProjectDetails(w http.ResponseWriter, r *http.Request) {
	projectID, err := parseItemID(r)
	if err != nil {
		c.HandleError(w, fmt.Sprintf("Failed to parse project id. %s", errors.WithStack(err)), http.StatusInternalServerError, UserMainPath)
		return
	}

	content, err := c.ContentController.BuildProjectDetailsContent(projectID)
	if err != nil {
		c.HandleError(w, fmt.Sprintf("Failed to get project content. %s", errors.WithStack(err)), http.StatusInternalServerError, UserMainPath)
		return
	}

	c.RenderTemplate(w, UserMainProjectDetails, content)
}

func (c *RESTFrontend) CreateProject(w http.ResponseWriter, r *http.Request) {
	productID, err := parseItemID(r)
	if err != nil {
		c.HandleError(w, fmt.Sprintf("Failed to parse product id. %s", errors.WithStack(err)), http.StatusInternalServerError, UserMainPath)
		return
	}

	if r.Method == GET {
		content, err := c.ContentController.BuildProjectWizardContent(productID)
		if err != nil {
			c.HandleError(w, fmt.Sprintf("Failed to get project wizard content content. %s", errors.WithStack(err)), http.StatusInternalServerError, UserMainPath)
			return
		}
		c.RenderTemplate(w, ProjectWizard, content)
	} else {
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			c.HandleError(w, fmt.Sprintf("Failed to parse avatar. %s", errors.WithStack(err)), http.StatusInternalServerError, UserMainPath)
			return
		}

		if err := contents.ValidateVisibility(r.FormValue(contents.ProjectVisibilityKey)); err != nil {
			c.HandleError(w, fmt.Sprintf("Failed to parse visibility. %s", errors.WithStack(err)), http.StatusInternalServerError, UserMainPath)
			return
		}

		if err := c.RESTBackend.CreateProject(r); err != nil {
			c.HandleError(w, fmt.Sprintf("Failed to create project. %s", errors.WithStack(err)), http.StatusInternalServerError, UserMainPath)
			return
		}

		c.MyProjects(w, r)
	}
}

func (c *RESTFrontend) HandleStatusRequest(w http.ResponseWriter, r *http.Request) {
	projectID, err := parseItemID(r)
	if err != nil {
		c.HandleError(w, fmt.Sprintf("Failed to parse project id. %s", errors.WithStack(err)), http.StatusInternalServerError, UserMainPath)
		return
	}

	state, err := c.RESTBackend.CheckProjectState(projectID)
	if err != nil {
		c.HandleError(w, fmt.Sprintf("Failed to access project. %s", errors.WithStack(err)), http.StatusInternalServerError, UserMainPath)
		return
	}

	if state == "unreachable" {
		c.HandleError(w, "Failed to access project", http.StatusNoContent, UserMainPath)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *RESTFrontend) RunProject(w http.ResponseWriter, r *http.Request) {
	projectID, err := parseItemID(r)
	if err != nil {
		c.HandleError(w, fmt.Sprintf("Failed to parse project id. %s", errors.WithStack(err)), http.StatusInternalServerError, UserMainPath)
		return
	}

	if err := c.RESTBackend.RunProject(c.ContentController.User.ID, projectID); err != nil {
		c.HandleError(w, fmt.Sprintf("Failed to run project. %s", errors.WithStack(err)), http.StatusInternalServerError, UserMainPath)
		return
	}

	content, err := c.ContentController.BuildProjectRunContent(projectID)
	if err != nil {
		c.HandleError(w, fmt.Sprintf("Failed to get project content. %s", errors.WithStack(err)), http.StatusInternalServerError, UserMainPath)
		return
	}

	c.RenderTemplate(w, "show", content)
}

func (c *RESTFrontend) ShowProject(w http.ResponseWriter, r *http.Request) {
	projectID, err := parseItemID(r)
	if err != nil {
		c.HandleError(w, fmt.Sprintf("Failed to parse project id. %s", errors.WithStack(err)), http.StatusInternalServerError, UserMainPath)
		return
	}

	content, err := c.ContentController.BuildProjectRunContent(projectID)
	if err != nil {
		c.HandleError(w, fmt.Sprintf("Failed to get project content. %s", errors.WithStack(err)), http.StatusInternalServerError, UserMainPath)
		return
	}

	c.RenderTemplate(w, "show", content)
}

func (c *RESTFrontend) DeleteProject(w http.ResponseWriter, r *http.Request) {
	if r.Method == POST {
		projectID, err := parseItemID(r)
		if err != nil {
			c.HandleError(w, fmt.Sprintf("Failed to parse project id. %s", errors.WithStack(err)), http.StatusInternalServerError, UserMainPath)
			return
		}

		if err := c.RESTBackend.DeleteProject(projectID); err != nil {
			c.HandleError(w, fmt.Sprintf("Failed to delete project. %s", errors.WithStack(err)), http.StatusInternalServerError, UserMainPath)
			return
		}

		c.MyProjects(w, r)
	}
}

func (c *RESTFrontend) EditProject(w http.ResponseWriter, r *http.Request) {
	projectID, err := parseItemID(r)
	if err != nil {
		c.HandleError(w, fmt.Sprintf("Failed to parse project id. %s", errors.WithStack(err)), http.StatusInternalServerError, UserMainPath)
		return
	}

	if r.Method == GET {
		content, err := c.ContentController.BuildProjectEditContent(projectID)
		if err != nil {
			c.HandleError(w, fmt.Sprintf("Failed to build project edit content. %s", errors.WithStack(err)), http.StatusInternalServerError, UserMainPath)
			return
		}

		c.RenderTemplate(w, "project-edit", content)
	} else {
		err := c.RESTBackend.UpdateProject(r)
		if err != nil {
			c.HandleError(w, fmt.Sprintf("Failed to update project. %s", errors.WithStack(err)), http.StatusInternalServerError, UserMainPath)
			return
		}

		content, err := c.ContentController.BuildMyProjectsContent()
		if err != nil {
			c.HandleError(w, fmt.Sprintf("Failed to get my projects content. %s", errors.WithStack(err)), http.StatusInternalServerError, UserMainPath)
			return
		}
		c.RenderTemplate(w, MyProjects, content)
	}
}
