package restfrontend

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"polygnosics-frontend/contents"
	"polygnosics-frontend/restbackend"
	"polygnosics-frontend/restfrontend/session"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

type RESTFrontend struct {
	ContentController *contents.ContentController
	RESTBackend       *restbackend.RESTBackend
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
	"/templates/resources/news-content.html",
	"/templates/resources/docs.html",
	"/templates/resources/tutorials.html",
	"/templates/resources/faq.html",
	"/templates/resources/instructions.html",
	"/templates/resources/examples.html",
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
	GET  = "GET"
	POST = "POST"
)

const (
	UserMainPath           = "/user-main"
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
	ResourcesNews         = "news"
	ResourcesDocs         = "docs"
	ResourcesTutorials    = "tutorials"
	ResourcesFAQ          = "faq"
	ResourcesExamples     = "examples"
	ResourcesInstructions = "instructions"
	ResourcesFiles        = "files"
)

const (
	AboutWhoWeAre = "about"
	AboutCareer   = "career"
	AboutContact  = "contact"
)

const (
	IndexPath      = "/index"
	IndexLoginPath = "/auth_login"
	IndexPage      = "index"
	IndexContact   = "general-contact"
	IndexNews      = "general-news"
)

func parseItemID(r *http.Request) (string, error) {
	if err := r.ParseForm(); err != nil {
		return "", err
	}
	return r.FormValue("item-id"), nil
}

func NewRESTController() *RESTFrontend {
	backend := &restbackend.RESTBackend{}

	controller := &RESTFrontend{
		ContentController: &contents.ContentController{
			RESTBackend: backend,
		},
		RESTBackend: backend,
	}
	return controller
}

func prettyPrint(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
		return
	}
	fmt.Println("Failed to pretty print data")
}

// MakeHandler creates the page handler and check the route validity.
func (c *RESTFrontend) MakeHandler(fn func(http.ResponseWriter, *http.Request), router *mux.Router, isPublicPage bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO Issue#71: Figure out the proper settings and fix UI code that breaks because of CSP
		//w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		// TODO Issue#71: Figure out the proper settings and fix UI code that breaks because of CSP
		//w.Header().Set("Content-Security-Policy", "default-src 'self' http://0.0.0.0:10000; script-src 'self';")

		routeMatch := mux.RouteMatch{}
		if matched := router.Match(r, &routeMatch); !matched {
			c.HandleError(w, "Url does not exist", http.StatusInternalServerError, IndexPath)
			return
		}

		if !isPublicPage {
			sess, err := session.Store.Get(r, "cookie-name")
			if err != nil {
				c.HandleError(w, "Unable to retrieve session cookie.", http.StatusForbidden, IndexPath)
				return
			}

			if c.ContentController.User == nil {
				c.HandleError(w, "Cookie expired", http.StatusForbidden, IndexPath)
				return
			}

			match, err := session.IsAuthenticated(c.ContentController.User.ID, sess, r)
			if err != nil {
				errorString := fmt.Sprintf("Unable to check session cookie:\n%s\n", errors.WithStack(err))
				c.HandleError(w, errorString, http.StatusInternalServerError, IndexPath)
				return
			}

			if !match {
				c.HandleError(w, "Forbidden access", http.StatusForbidden, IndexPath)
				return
			}
		}
		fn(w, r)
	}
}

// RenderTemplate renders html.
func (c *RESTFrontend) RenderTemplate(w http.ResponseWriter, tmpl string, p map[string]interface{}) {
	wd, err := os.Getwd()
	if err != nil {
		c.HandleError(w, err.Error(), http.StatusInternalServerError, IndexPath)
	}

	if len(paths) == 0 {
		for i := 0; i < len(htmls); i++ {
			paths = append(paths, wd+htmls[i])
		}
	}

	t := template.Must(template.ParseFiles(paths...))

	err = t.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		c.HandleError(w, err.Error(), http.StatusInternalServerError, IndexPath)
	}
}

// HandleError creates page details and renders html template for an error modal.
func (c *RESTFrontend) HandleError(w http.ResponseWriter, errorStr string, statusCode int, backPage string) {
	content := make(map[string]interface{})
	content["parent_page"] = "Error"
	content["status_code"] = statusCode
	content["status_text"] = http.StatusText(statusCode)
	content["message"] = errorStr
	content["back_page"] = backPage
	c.RenderTemplate(w, "error", content)
}
