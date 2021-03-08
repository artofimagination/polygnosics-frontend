package restcontrollers

import (
	"fmt"
	"net/http"
	"os"
	"polygnosics/app/businesslogic"
	"polygnosics/app/restcontrollers/session"
	"polygnosics/web/contents"
	"text/template"

	"github.com/artofimagination/mysql-user-db-go-interface/dbcontrollers"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

type RESTController struct {
	UserDBController  *dbcontrollers.MYSQLController
	ContentController *contents.ContentController
	BackendContext    *businesslogic.Context
}

var ErrFailedToParseForm = "Failed to parse form"

var htmls = []string{
	"/web/templates/about/about.html",
	"/web/templates/about/career.html",
	"/web/templates/about/contact.html",
	"/web/templates/error.html",
	"/web/templates/index.html",
	"/web/templates/confirm.html",
	"/web/templates/terms.html",
	"/web/templates/general-contact.html",
	"/web/templates/general-news.html",
	"/web/templates/user/user-main.html",
	"/web/templates/user/profile.html",
	"/web/templates/user/profile-side-bar.html",
	"/web/templates/user/profile-edit.html",
	"/web/templates/user/profile-edit-avatar.html",
	"/web/templates/user/settings.html",
	"/web/templates/user/mail-inbox.html",
	"/web/templates/user/mail-compose.html",
	"/web/templates/user/mail-read.html",
	"/web/templates/admin/dashboard.html",
	"/web/templates/project/show.html",
	"/web/templates/project/browser.html",
	"/web/templates/project/project-details.html",
	"/web/templates/project/my-projects.html",
	"/web/templates/project/project-edit.html",
	"/web/templates/project/project-wizard.html",
	"/web/templates/project/documentation.html",
	"/web/templates/project/resources.html",
	"/web/templates/auth_signup.html",
	"/web/templates/auth_login.html",
	"/web/templates/products/store.html",
	"/web/templates/products/product-wizard.html",
	"/web/templates/products/product-edit.html",
	"/web/templates/products/my-products.html",
	"/web/templates/products/details.html",
	"/web/templates/products/documentation.html",
	"/web/templates/products/resources.html",
	"/web/templates/components/side-bar.html",
	"/web/templates/components/content-header.html",
	"/web/templates/components/header-info.html",
	"/web/templates/components/main-header.html",
	"/web/templates/components/footer.html",
	"/web/templates/components/news-feed.html",
	"/web/templates/components/index-header.html",
	"/web/templates/components/index-footer.html",
	"/web/templates/resources/news.html",
	"/web/templates/resources/news-content.html",
	"/web/templates/resources/docs.html",
	"/web/templates/resources/tutorials.html",
	"/web/templates/resources/faq.html",
	"/web/templates/resources/instructions.html",
	"/web/templates/resources/examples.html",
	"/web/templates/resources/files.html",
	"/web/templates/stats/project-stats.html",
	"/web/templates/stats/product-stats.html",
	"/web/templates/stats/user-stats.html",
	"/web/templates/stats/system-health.html",
	"/web/templates/stats/accounting.html",
	"/web/templates/stats/ui-stats.html",
	"/web/templates/stats/misuse-metrics.html",
	"/web/templates/stats/product-project-stats.html",
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
	IndexPath    = "/index"
	IndexPage    = "index"
	IndexContact = "general-contact"
	IndexNews    = "general-news"
)

func parseItemID(r *http.Request) (*uuid.UUID, error) {
	if err := r.ParseForm(); err != nil {
		return nil, err
	}
	itemID, err := uuid.Parse(r.FormValue("item-id"))
	if err != nil {
		return nil, err
	}
	return &itemID, nil
}

func NewRESTController(userDB *dbcontrollers.MYSQLController) *RESTController {
	controller := &RESTController{
		UserDBController: userDB,
		ContentController: &contents.ContentController{
			UserDBController: userDB,
		},
		BackendContext: &businesslogic.Context{
			UserDBController: userDB,
		},
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
			c.HandleError(w, "Url does not exist", http.StatusInternalServerError, IndexPath)
			return
		}

		if !isPublicPage {
			sess, err := session.Store.Get(r, "cookie-name")
			if err != nil {
				c.HandleError(w, "Unable to retrieve session cookie.", http.StatusForbidden, IndexPath)
				return
			}

			uuidString, ok := sess.Values["user"].(string)
			if !ok {
				c.HandleError(w, "Unable to decode session cookie.", http.StatusInternalServerError, IndexPath)
				return
			}

			userUUID, err := uuid.Parse(uuidString)
			if err != nil {
				c.HandleError(w, fmt.Sprintf("Failed to get user id. %s", errors.WithStack(err)), http.StatusInternalServerError, IndexPath)
				return
			}

			user, err := c.UserDBController.GetUser(&userUUID)
			if err != nil {
				c.HandleError(w, "Unable to retrieve user info", http.StatusInternalServerError, IndexPath)
				return
			}

			match, err := session.IsAuthenticated(user.ID, sess, r)
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
func (c *RESTController) RenderTemplate(w http.ResponseWriter, tmpl string, p map[string]interface{}) {
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
func (c *RESTController) HandleError(w http.ResponseWriter, errorStr string, statusCode int, backPage string) {
	content := make(map[string]interface{})
	content["parent_page"] = "Error"
	content["status_code"] = statusCode
	content["status_text"] = http.StatusText(statusCode)
	content["message"] = errorStr
	content["back_page"] = backPage
	c.RenderTemplate(w, "error", content)
}
