package contents

import (
	"net/http"

	"github.com/artofimagination/polygnosics-frontend/rest/backend"
)

const (
	ProductsPageName           = "Products"
	ProductsPageCreateName     = "Product Wizard"
	ProductsPageMyProductsName = "My Products"
	ProductsPageDetailsName    = "Details"
	ProductsPageStoreName      = "Marketplace"
	ProductsPageEditName       = "Edit"
)

const (
	IndexPageSignupName = "Sign up"
	IndexPageLoginName  = "Sign in"
)

const (
	ProjectsPageName           = "Projects"
	ProjectsPageCreateName     = "Project Wizard"
	ProjectsPageMyProjectsName = "My Projects"
	ProjectsPageDetailsName    = "Details"
	ProjectsPageBrowserName    = "Browser"
	ProjectsPageEditName       = "Edit"
	ProjectsPageRunName        = "Run"
)

const (
	UserPageName         = "User"
	UserPageProfileName  = "Profile"
	UserPageMainPageName = "Info board"
	UserPageSettingsName = "Settings"
)

const (
	ResourcesPageName              = "Resources"
	ResourcesPageNewsName          = "News"
	ResourcesPageDocumentationName = "Documentation"
	ResourcesPageTutorialsName     = "Tutorials"
	ResourcesPageFAQsName          = "FAQs"
	ResourcesPageExamplesName      = "Examples"
	ResourcesPageFilesName         = "Files"
	ResourcesPageNewFAQName        = "Create FAQ"
	ResourcesPageNewNewsName       = "Create News"
	ResourcesPageNewTutorialName   = "Create Tutorial"
	ResourcesPageNewFilesName      = "Create Files"
	ResourcesPageEditFAQName       = "Edit FAQ"
	ResourcesPageEditNewsName      = "Edit News"
	ResourcesPageEditTutorialName  = "Edit Tutorial"
	ResourcesPageEditFilesName     = "Edit Files"
)

const (
	StatsPageName     = "Statistics"
	StatsUsers        = "User Statistics"
	StatsItems        = "Overall Product and Project Stats"
	StatsProducts     = "Product Statistics"
	StatsProjects     = "Project Statistics"
	StatsAccounting   = "Accounting Statistics"
	StatsUI           = "UI Statistics"
	StatsSystemHealth = "System Health"
)

var FutureFeature = "future_feature"

// TODO Issue#40: Replace  user/product/project data with redis storage.
type ContentController struct {
	User        *backend.User
	RESTBackend *backend.RESTController
}

func (c *ContentController) BuildProductWizardContent() (map[string]interface{}, error) {
	content := c.GetUserContent(c.User)
	content = c.prepareContentHeader(content, ProductsPageName, ProductsPageCreateName)
	categories, err := c.RESTBackend.GetCategoriesMap()
	if err != nil {
		return nil, err
	}
	content[ProductCategoriesKey] = categories
	return content, err
}

func (c *ContentController) BuildProjectWizardContent(productID string) (map[string]interface{}, error) {
	content := c.GetUserContent(c.User)
	content = c.prepareContentHeader(content, ProjectsPageName, ProjectsPageCreateName)
	product, err := c.GetProductContent(productID)
	if err != nil {
		return nil, err
	}
	content[ProductMapKey] = product
	return content, nil
}

func (c *ContentController) BuildProductEditContent(productID string) (map[string]interface{}, error) {
	content := c.GetUserContent(c.User)
	content = c.prepareContentHeader(content, ProductsPageName, ProductsPageEditName)
	productContent, err := c.GetProductContent(productID)
	if err != nil {
		return nil, err
	}
	content[ProductMapKey] = productContent
	categories, err := c.RESTBackend.GetCategoriesMap()
	if err != nil {
		return nil, err
	}
	content[ProductCategoriesKey] = categories
	return content, err
}

func (c *ContentController) BuildProjectEditContent(projectID string) (map[string]interface{}, error) {
	content := c.GetUserContent(c.User)
	content = c.prepareContentHeader(content, ProjectsPageName, ProjectsPageEditName)
	projectContent, err := c.GetProjectContent(projectID)
	if err != nil {
		return nil, err
	}
	content[ProjectMapKey] = projectContent
	categories, err := c.RESTBackend.GetCategoriesMap()
	if err != nil {
		return nil, err
	}
	content[ProductCategoriesKey] = categories

	return content, err
}

func (c *ContentController) BuildProjectRunContent(projectID string) (map[string]interface{}, error) {
	content := c.GetUserContent(c.User)
	content = c.prepareContentHeader(content, ProjectsPageName, ProjectsPageRunName)
	projectContent, err := c.GetProjectContent(projectID)
	if err != nil {
		return nil, err
	}
	content[ProjectMapKey] = projectContent
	return content, err
}

func (c *ContentController) BuildMyProductsContent() (map[string]interface{}, error) {
	content := c.GetUserContent(c.User)
	content = c.prepareContentHeader(content, ProductsPageName, ProductsPageMyProductsName)
	productsContent, err := c.GetUserProductContent(c.User.ID)
	if err != nil {
		return nil, err
	}
	content[ProductMapKey] = productsContent
	return content, nil
}

func (c *ContentController) BuildProductDetailsContent(productID string) (map[string]interface{}, error) {
	content := c.GetUserContent(c.User)
	content = c.prepareContentHeader(content, ProductsPageName, ProductsPageDetailsName)
	productContent, err := c.GetProductContent(productID)
	if err != nil {
		return nil, err
	}
	content[ProductMapKey] = productContent
	return content, nil
}

func (c *ContentController) BuildMyProjectsContent() (map[string]interface{}, error) {
	content := c.GetUserContent(c.User)
	content = c.prepareContentHeader(content, ProjectsPageName, ProjectsPageMyProjectsName)
	projectsContent, err := c.GetUserProjectContent(c.User.ID, -1)
	if err != nil {
		return nil, err
	}
	content["project"] = projectsContent
	return content, nil
}

func (c *ContentController) BuildProjectDetailsContent(projectID string) (map[string]interface{}, error) {
	content := c.GetUserContent(c.User)
	content = c.prepareContentHeader(content, ProjectsPageName, ProjectsPageDetailsName)
	projectContent, err := c.GetProjectContent(projectID)
	if err != nil {
		return nil, err
	}
	content[ProjectMapKey] = projectContent
	return content, nil
}

func (c *ContentController) BuildProfileContent(id string) (map[string]interface{}, error) {
	user := c.User
	if id != c.User.ID {
		userNew, err := c.RESTBackend.GetUserByID(id)
		if err != nil {
			return nil, err
		}
		user = userNew
	}
	content := c.GetUserContent(user)
	content = c.prepareContentHeader(content, UserPageName, UserPageProfileName)
	return content, nil
}

func (c *ContentController) BuildUserMainContent() (map[string]interface{}, error) {
	content := c.GetUserContent(c.User)
	content = c.prepareContentHeader(content, UserPageName, UserPageMainPageName)
	err := c.prepareNewsFeed(content)
	if err != nil {
		return nil, err
	}
	productsContent, err := c.GetRecentProductsContent(c.User.ID)
	if err != nil {
		return nil, err
	}
	content["product"] = productsContent
	projectsContent, err := c.GetUserProjectContent(c.User.ID, 4)
	if err != nil {
		return nil, err
	}
	content["project"] = projectsContent
	return content, nil
}

func (c *ContentController) BuildErrorContent(errString string) map[string]interface{} {
	content := c.GetUserContent(c.User)
	content["message"] = errString
	return content
}

func (c *ContentController) BuildNewsContent() (map[string]interface{}, error) {
	content := c.GetUserContent(c.User)
	content = c.prepareContentHeader(content, ResourcesPageName, ResourcesPageNewsName)
	err := c.prepareNewsFeed(content)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (c *ContentController) BuildUserStatsContent() map[string]interface{} {
	content := c.GetUserContent(c.User)
	content = c.prepareContentHeader(content, StatsPageName, StatsUsers)
	return content
}

func (c *ContentController) BuildProductStatsContent() map[string]interface{} {
	content := c.GetUserContent(c.User)
	content = c.prepareContentHeader(content, StatsPageName, StatsProducts)
	return content
}

func (c *ContentController) BuildSignupContent() map[string]interface{} {
	content := make(map[string]interface{})
	content = c.prepareContentHeader(content, IndexPageSignupName, "")
	return content
}

func (c *ContentController) BuildLoginContent() map[string]interface{} {
	content := make(map[string]interface{})
	content = c.prepareContentHeader(content, IndexPageLoginName, "")
	return content
}

func (c *ContentController) BuildProjectStatsContent() map[string]interface{} {
	content := c.GetUserContent(c.User)
	content = c.prepareContentHeader(content, StatsPageName, StatsProjects)
	return content
}

func (c *ContentController) BuildAccountingStatsContent() map[string]interface{} {
	content := c.GetUserContent(c.User)
	content = c.prepareContentHeader(content, StatsPageName, StatsAccounting)
	return content
}

func (c *ContentController) BuildSystemHealthContent() map[string]interface{} {
	content := c.GetUserContent(c.User)
	content = c.prepareContentHeader(content, StatsPageName, StatsSystemHealth)
	return content
}

func (c *ContentController) BuildUIStatsContent() map[string]interface{} {
	content := c.GetUserContent(c.User)
	content = c.prepareContentHeader(content, StatsPageName, StatsUI)
	return content
}

func (c *ContentController) BuildItemStatsContent() map[string]interface{} {
	content := c.GetUserContent(c.User)
	content = c.prepareContentHeader(content, StatsPageName, StatsItems)
	return content
}

func (c *ContentController) BuildMailInboxContent() map[string]interface{} {
	content := c.GetUserContent(c.User)
	return content
}

func (c *ContentController) BuildMailComposeContent() map[string]interface{} {
	content := c.GetUserContent(c.User)
	return content
}

func (c *ContentController) BuildDocsContent() map[string]interface{} {
	content := c.GetUserContent(c.User)
	content = c.prepareContentHeader(content, ResourcesPageName, ResourcesPageDocumentationName)
	return content
}

func (c *ContentController) BuildArticleContent(r *http.Request) (map[string]interface{}, error) {
	content := c.GetUserContent(c.User)
	content = c.prepareContentHeader(content, ResourcesPageName, ResourcesPageDocumentationName)
	if err := c.prepareArticle(content, r); err != nil {
		return nil, err
	}
	return content, nil
}

func (c *ContentController) BuildCreateNews() map[string]interface{} {
	content := c.GetUserContent(c.User)
	c.prepareCreateNewsFeed(content)
	content = c.prepareContentHeader(content, ResourcesPageName, ResourcesPageNewNewsName)
	return content
}

func (c *ContentController) BuildEditNews(id string) (map[string]interface{}, error) {
	content := c.GetUserContent(c.User)
	if err := c.prepareEditNews(id, content); err != nil {
		return nil, err
	}
	content = c.prepareContentHeader(content, ResourcesPageName, ResourcesPageEditNewsName)
	return content, nil
}

func (c *ContentController) BuildCreateTutorial() map[string]interface{} {
	content := c.GetUserContent(c.User)
	c.prepareNewTutorial(content)
	content = c.prepareContentHeader(content, ResourcesPageName, ResourcesPageNewTutorialName)

	return content
}

func (c *ContentController) BuildEditTutorial(id string) (map[string]interface{}, error) {
	content := c.GetUserContent(c.User)
	if err := c.prepareEditTutorial(id, content); err != nil {
		return nil, err
	}
	content = c.prepareContentHeader(content, ResourcesPageName, ResourcesPageEditTutorialName)
	return content, nil
}

func (c *ContentController) BuildCreateFAQ() (map[string]interface{}, error) {
	content := c.GetUserContent(c.User)
	if err := c.prepareNewFAQ(content); err != nil {
		return nil, err
	}
	content = c.prepareContentHeader(content, ResourcesPageName, ResourcesPageNewFAQName)
	return content, nil
}

func (c *ContentController) BuildEditFAQ(id string) (map[string]interface{}, error) {
	content := c.GetUserContent(c.User)
	if err := c.prepareEditFAQ(id, content); err != nil {
		return nil, err
	}
	content = c.prepareContentHeader(content, ResourcesPageName, ResourcesPageEditFAQName)
	return content, nil
}

func (c *ContentController) BuildCreateFiles() map[string]interface{} {
	content := c.GetUserContent(c.User)
	c.prepareNewFiles(content)
	content = c.prepareContentHeader(content, ResourcesPageName, ResourcesPageNewFilesName)
	return content
}

func (c *ContentController) BuildEditFilesSection(id string) (map[string]interface{}, error) {
	content := c.GetUserContent(c.User)
	if err := c.prepareEditFiles(id, content); err != nil {
		return nil, err
	}
	content = c.prepareContentHeader(content, ResourcesPageName, ResourcesPageEditFilesName)
	return content, nil
}

func (c *ContentController) BuildFilesContent() (map[string]interface{}, error) {
	content := c.GetUserContent(c.User)
	content = c.prepareContentHeader(content, ResourcesPageName, ResourcesPageFilesName)
	if err := c.prepareFiles(content); err != nil {
		return nil, err
	}
	return content, nil
}

func (c *ContentController) BuildSettingsContent() map[string]interface{} {
	content := c.GetUserContent(c.User)
	content = c.prepareContentHeader(content, UserPageName, UserPageSettingsName)
	return content
}

func (c *ContentController) BuildTutorialsContent() (map[string]interface{}, error) {
	content := c.GetUserContent(c.User)
	content = c.prepareContentHeader(content, ResourcesPageName, ResourcesPageTutorialsName)
	if err := c.prepareTutorials(content); err != nil {
		return nil, err
	}
	return content, nil
}

func (c *ContentController) BuildFAQContent() (map[string]interface{}, error) {
	content := c.GetUserContent(c.User)
	content = c.prepareContentHeader(content, ResourcesPageName, ResourcesPageFAQsName)
	if err := c.prepareFAQ(content); err != nil {
		return nil, err
	}
	return content, nil
}

func (c *ContentController) BuildMailReadContent() map[string]interface{} {
	content := c.GetUserContent(c.User)
	return content
}

func (c *ContentController) BuildContactContent() map[string]interface{} {
	content := c.GetUserContent(c.User)
	return content
}

func (c *ContentController) BuildStoreContent() (map[string]interface{}, error) {
	content := c.GetUserContent(c.User)
	content = c.prepareContentHeader(content, ProductsPageName, ProductsPageStoreName)
	categorizedProducts, err := c.GetProductsByCategory(c.User.ID)
	if err != nil {
		return nil, err
	}
	recent, err := c.GetRecentProductsContent(c.User.ID)
	if err != nil {
		return nil, err
	}
	recommended, err := c.GetRecommendedProductsContent(c.User.ID)
	if err != nil {
		return nil, err
	}
	content[ProductCategoriesKey] = categorizedProducts
	content["recent"] = recent
	content["recommended"] = recommended
	return content, nil
}

func (c *ContentController) BuildProjectBrowserContent() (map[string]interface{}, error) {
	content := c.GetUserContent(c.User)
	content = c.prepareContentHeader(content, ProductsPageName, ProjectsPageBrowserName)
	categorizedProducts, err := c.GetProjectsByCategory(c.User.ID)
	if err != nil {
		return nil, err
	}
	recent, err := c.GetRecentProjectsContent(c.User.ID)
	if err != nil {
		return nil, err
	}
	recommended, err := c.GetRecommendedProjectsContent(c.User.ID)
	if err != nil {
		return nil, err
	}

	content[ProductCategoriesKey] = categorizedProducts
	content["recent"] = recent
	content["recommended"] = recommended
	return content, nil
}
