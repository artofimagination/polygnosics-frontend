package contents

import (
	"fmt"

	"github.com/artofimagination/polygnosics-frontend/rest/backend"
)

// Details and assets field keys
const (
	ProjectMapKey               = "project"
	NewProject                  = "new_project"
	RunProject                  = "run_project"
	ShowProject                 = "show_project"
	ProjectDetailsPageKey       = "detail_path"
	ProjectStateBadge           = "state_badge"
	ProjectDeleteIDKey          = "project_id"
	ProjectDeleteTextKey        = "delete_text"
	ProjectDeleteSuccessTextKey = "delete_success_text"
	ProjectDeleteURLKey         = "delete_url"
	ProjectEditPageKey          = "edit_path"
	ProjectVisibilityKey        = "visibility"

	ProjectServerLogging = "server_logging"
	ProjectClientLogging = "client_logging"
)

// Visibility values of a project
const (
	Public    = "Public"
	Protected = "Protected"
	Private   = "Private"
)

const (
	NotRunning  = "Not running"
	Running     = "Running"
	Stopped     = "Stopped"
	Paused      = "Paused"
	Unreachable = "Unreachable"
)

// ValidateVisibility validates the visibility string
func ValidateVisibility(value string) error {
	if value != Public && value != Protected && value != Private {
		return fmt.Errorf("Invalid visibility: %s", value)
	}
	return nil
}

// generateProjectContent fills a string nested map with all project details and assets info
func (c *ContentController) generateProjectContent(projectData *backend.Project) map[string]interface{} {
	content := projectData.Details
	for k, v := range projectData.Assets {
		content[k] = v
	}
	content[ProjectServerLogging] = convertToCheckboxValue(content[ProjectServerLogging].(string))
	content[ProjectClientLogging] = convertToCheckboxValue(content[ProjectClientLogging].(string))
	content[ProjectStateBadge] = getProjectStateContent(content[ProjectStateBadge].(string)).badge

	content[ProjectDetailsPageKey] = fmt.Sprintf("/user-main/my-projects/details?item-id=%s", projectData.ID)
	content[ProjectEditPageKey] = fmt.Sprintf("/user-main/my-projects/edit?item-id=%s", projectData.ID)
	content[RunProject] = fmt.Sprintf("/user-main/my-projects/run?item-id=%s", projectData.ID)
	content[ShowProject] = fmt.Sprintf("/user-main/my-projects/show?item-id=%s", projectData.ID)
	content[ProjectDeleteIDKey] = projectData.ID
	content[ProjectDeleteTextKey] = "The project will not be accessible anymore"
	content[ProjectDeleteSuccessTextKey] = "Your project has been deleted"
	content[ProjectDeleteURLKey] = "/user-main/my-projects/delete"
	return content
}

// GetProjectContent returns the selected project details and assets info.
func (c *ContentController) GetProjectContent(projectID string) (map[string]interface{}, error) {
	project, err := c.RESTBackend.GetProject(projectID)
	if err != nil {
		return nil, err
	}
	return c.generateProjectContent(project), nil
}

// GetUserProjectContent gathers the contents of all projects belonging to the specified user.
func (c *ContentController) GetUserProjectContent(userID string, limit int) ([]map[string]interface{}, error) {
	projects, err := c.RESTBackend.GetProjectsByUserID(userID)
	if err != nil {
		return nil, err
	}

	if limit > len(projects) {
		limit = len(projects)
	}

	projectContent := make([]map[string]interface{}, len(projects))
	if limit != -1 {
		projectContent = make([]map[string]interface{}, limit)
	}

	for i, project := range projects {
		if limit == 0 {
			break
		}
		limit--
		projectContent[i] = c.generateProjectContent(project)
	}

	return projectContent, nil
}

// GetProjectsByCategory organizes project contents by categories. This is just a placeholder solution until proper search is introduced.
func (c *ContentController) GetProjectsByCategory(userID string) (map[string]interface{}, error) {
	projects, err := c.RESTBackend.GetProjectsByUserID(userID)
	if err != nil {
		return nil, err
	}

	categorizedProjects := make(map[string]interface{})
	categories, err := c.RESTBackend.GetCategoriesMap()
	if err != nil {
		return nil, err
	}
	for k, v := range categories {
		if categorizedProjects[k] == nil {
			categorizedProjects[k] = make(map[string]interface{})
		}
		categorizedProjects[k].(map[string]interface{})["name"] = v
		for _, project := range projects {
			for _, categoryKey := range project.Details[ProductCategoriesKey].([]interface{}) {
				if categoryKey.(string) == k {
					_, ok := categorizedProjects[k].(map[string]interface{})["projects"]
					if !ok {
						categorizedProjects[k].(map[string]interface{})["projects"] = make([]map[string]interface{}, 0)
					}
					categorizedProjects[k].(map[string]interface{})["projects"] = append(categorizedProjects[k].(map[string]interface{})["projects"].([]map[string]interface{}), c.generateProjectContent(project))
					break
				}
			}
		}
	}
	return categorizedProjects, nil
}

// GetRecentProjectsContent gathers the content of the latest 4 projects
func (c *ContentController) GetRecentProjectsContent(userID string) ([]map[string]interface{}, error) {
	projects, err := c.RESTBackend.GetProjectsByUserID(userID)
	if err != nil {
		return nil, err
	}

	limit := 4
	if len(projects) < limit {
		limit = len(projects)
	}
	projectContent := make([]map[string]interface{}, limit)
	for i, project := range projects {
		if limit == 0 {
			break
		}
		limit--
		projectContent[i] = c.generateProjectContent(project)
		projectContent[i][ProductOwnerNameKey] = c.User.Settings[UserNameKey]
		projectContent[i][ProductOwnerPageNameKey] = fmt.Sprintf("/user-main/profile?user=%s", c.User.ID)
		projectContent[i][ProjectDetailsPageKey] = fmt.Sprintf("/user-main/project?item-id=%s", project.ID)
	}

	return projectContent, nil
}

func (c *ContentController) GetRecommendedProjectsContent(userID string) ([]map[string]interface{}, error) {
	projects, err := c.RESTBackend.GetProjectsByUserID(userID)
	if err != nil {
		return nil, err
	}

	limit := 4
	if len(projects) < limit {
		limit = len(projects)
	}
	projectContent := make([]map[string]interface{}, limit)
	for i, project := range projects {
		if limit == 0 {
			break
		}
		limit--
		projectContent[i] = c.generateProjectContent(project)
		projectContent[i][ProductOwnerNameKey] = c.User.Settings[UserNameKey]
		projectContent[i][ProductOwnerPageNameKey] = fmt.Sprintf("/user-main/profile?user=%s", c.User.ID)
		projectContent[i][ProjectDetailsPageKey] = fmt.Sprintf("/user-main/project?item-id=%s", project.ID)
	}

	return projectContent, nil
}
