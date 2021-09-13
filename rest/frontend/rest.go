package frontend

import (
	"fmt"
	"net/http"
	"os"
	"text/template"

	"github.com/artofimagination/polygnosics-frontend/backend"
	"github.com/artofimagination/polygnosics-frontend/contents"
	"github.com/artofimagination/polygnosics-frontend/frontend/session"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

type RESTController struct {
	ContentController *contents.ContentController
	RESTBackend       *backend.RESTController
}

var ErrFailedToParseForm = "Failed to parse form"

var htmls = []string{
	"/templates/about/about.html",
	"/templates/about/career.html",
	"/templates/about/contact.html",
	"/templates/error.html",
	"/templates/index.html",
	"/templates/terms.html",
	"/templates/general-contact.html",
	"/templates/general-news.html",
	"/templates/user/user-main.html",
	"/templates/user/profile.html",
	"/templates/user/profile-side-bar.html",
	"/templates/user/profile-edit.html",
	"/templates/user/profile-edit-avatar.html",
	"/templates/user/settings.html",
	"/templates/user/mail-inbox.html",
	"/templates/user/mail-compose.html",
	"/templates/user/mail-read.html",
	"/templates/admin/dashboard.html",
	"/templates/project/show.html",
	"/templates/project/browser.html",
	"/templates/project/project-details.html",
	"/templates/project/my-projects.html",
	"/templates/project/project-edit.html",
	"/templates/project/project-wizard.html",
	"/templates/project/documentation.html",
	"/templates/project/resources.html",
	"/templates/auth_signup.html",
	"/templates/auth_login.html",
	"/templates/products/store.html",
	"/templates/products/product-wizard.html",
	"/templates/products/product-edit.html",
	"/templates/products/my-products.html",
	"/templates/products/details.html",
	"/templates/products/documentation.html",
	"/templates/products/resources.html",
	"/templates/components/side-bar.html",
	"/templates/components/content-header.html",
	"/templates/components/header-info.html",
	"/templates/components/main-header.html",
	"/templates/components/footer.html",
	"/templates/components/news-feed.html",
	"/templates/components/index-header.html",
	"/templates/components/index-footer.html",
	"/templates/resources/news.html",
	"/templates/resources/create-news-item.html",
	"/templates/resources/edit-news-item.html",
	"/templates/resources/create-faq-item.html",
	"/templates/resources/create-tutorial-item.html",
	"/templates/resources/edit-tutorial-item.html",
	"/templates/resources/create-files-item.html",
	"/templates/resources/edit-files-item.html",
	"/templates/resources/edit-faq-item.html",
	"/templates/resources/docs.html",
	"/templates/resources/tutorials.html",
	"/templates/resources/article.html",
	"/templates/resources/faq.html",
	"/templates/resources/files.html",
	"/templates/stats/project-stats.html",
	"/templates/stats/product-stats.html",
	"/templates/stats/user-stats.html",
	"/templates/stats/system-health.html",
	"/templates/stats/accounting.html",
	"/templates/stats/ui-stats.html",
	"/templates/stats/misuse-metrics.html",
	"/templates/stats/product-project-stats.html",
}
var paths = []string{}

const (
	UserMain               = "user-main"
	UserMainMyProducts     = "my-products"
	ProjectWizard          = "project-wizard"
	MyProjects             = "my-projects"
	UserMainProjectDetails = "project-details"
	UserMainMailInbox      = "mail-inbox"
	UserMainMailCompose    = "mail-compose"
	UserMainMailRead       = "mail-read"
	UserMainSettings       = "settings"
)

const (
	StatsProject        = "project-stats"
	StatsProduct        = "product-stats"
	StatsUser           = "user-stats"
	StatsProductProject = "product-project-stats"
	StatsUI             = "ui-stats"
	StatsSystemHealth   = "system-health"
	StatsAccounting     = "accounting"
	StatsMisuseMetrics  = "misuse-metrics"
)

const (
	ResourcesNews           = "news"
	ResourcesDocs           = "docs"
	ResourcesTutorials      = "tutorials"
	ResourcesFAQ            = "faq"
	ResourcesFiles          = "files"
	ResourcesArticle        = "article"
	ResourcesCreateNews     = "create-news-item"
	ResourcesCreateFAQ      = "create-faq-item"
	ResourcesCreateTutorial = "create-tutorial-item"
	ResourcesCreateFiles    = "create-files-item"
	ResourcesEditFAQ        = "edit-faq-item"
	ResourcesEditTutorial   = "edit-tutorial-item"
	ResourcesEditNews       = "edit-news-item"
	ResourcesEditFiles      = "edit-files-item"
)

const (
	AboutWhoWeAre = "about"
	AboutCareer   = "career"
	AboutContact  = "contact"
)

const (
	IndexPage      = "index"
	IndexLoginPage = "auth_login"
	IndexContact   = "general-contact"
	IndexNews      = "general-news"
)

func (c *RESTController) URI(html string) string {
	return fmt.Sprintf("/%s", html)
}

func parseItemID(r *http.Request) (string, error) {
	if err := r.ParseForm(); err != nil {
		return "", err
	}
	return r.FormValue("item-id"), nil
}

func NewRESTController() *RESTController {
	backend := &backend.RESTController{}

	controller := &RESTController{
		ContentController: &contents.ContentController{
			RESTBackend: backend,
		},
		RESTBackend: backend,
	}
	return controller
}

// MakeHandler creates the page handler and check the route validity.
func (c *RESTController) MakeHandler(fn func(http.ResponseWriter, *http.Request), router *mux.Router, isPublicPage bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO Issue#71: Figure out the proper settings and fix UI code that breaks because of CSP
		//w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		// TODO Issue#71: Figure out the proper settings and fix UI code that breaks because of CSP
		//w.Header().Set("Content-Security-Policy", "default-src 'self' http://0.0.0.0:10000; script-src 'self';")

		routeMatch := mux.RouteMatch{}
		if matched := router.Match(r, &routeMatch); !matched {
			c.HandleError(w, "Url does not exist", http.StatusInternalServerError, c.URI(IndexPage))
			return
		}

		if !isPublicPage {
			sess, err := session.Store.Get(r, "cookie-name")
			if err != nil {
				c.HandleError(w, "Unable to retrieve session cookie.", http.StatusForbidden, c.URI(IndexPage))
				return
			}

			if c.ContentController.User == nil {
				c.HandleError(w, "Cookie expired", http.StatusForbidden, c.URI(IndexPage))
				return
			}

			match, err := session.IsAuthenticated(c.ContentController.User.ID, sess, r)
			if err != nil {
				errorString := fmt.Sprintf("Unable to check session cookie:\n%s\n", errors.WithStack(err))
				c.HandleError(w, errorString, http.StatusInternalServerError, c.URI(IndexPage))
				return
			}

			if !match {
				c.HandleError(w, "Forbidden access", http.StatusForbidden, c.URI(IndexPage))
				return
			}
		}
		fn(w, r)
	}
}

// RenderTemplate renders html.
func (c *RESTController) RenderTemplate(w http.ResponseWriter, tmpl string, p map[string]interface{}) {
	wd, err := os.Getwd()
	if err != nil {
		c.HandleError(w, err.Error(), http.StatusInternalServerError, c.URI(IndexPage))
	}

	if len(paths) == 0 {
		for i := 0; i < len(htmls); i++ {
			paths = append(paths, wd+htmls[i])
		}
	}

	t := template.Must(template.ParseFiles(paths...))

	err = t.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		c.HandleError(w, err.Error(), http.StatusInternalServerError, c.URI(IndexPage))
	}
}

// HandleError creates page details and renders html template for an error modal.
func (c *RESTController) HandleError(w http.ResponseWriter, errorStr string, statusCode int, backPage string) {
	content := make(map[string]interface{})
	content["parent_page"] = "Error"
	content["status_code"] = statusCode
	content["status_text"] = http.StatusText(statusCode)
	content["message"] = errorStr
	content["back_page"] = backPage
	c.RenderTemplate(w, "error", content)
}
