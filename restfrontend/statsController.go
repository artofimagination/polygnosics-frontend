package restfrontend

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

func (c *RESTFrontend) StatsWebRTC(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		c.HandleError(w, fmt.Sprintf("Failed to parse frontend webrtc offer. %s", errors.WithStack(err)), http.StatusInternalServerError, c.URI(UserMain))
		return
	}

	if err := c.RESTBackend.InitStatsWebRTC(r); err != nil {
		c.HandleError(w, fmt.Sprintf("Failed to init webrtc data provider. %s", errors.WithStack(err)), http.StatusInternalServerError, c.URI(UserMain))
		return
	}
}

func (c *RESTFrontend) ProductStats(w http.ResponseWriter, r *http.Request) {
	content := c.ContentController.BuildProductStatsContent()
	c.RenderTemplate(w, StatsProduct, content)
}

func (c *RESTFrontend) ProjectStats(w http.ResponseWriter, r *http.Request) {
	content := c.ContentController.BuildProjectStatsContent()
	c.RenderTemplate(w, StatsProject, content)
}

func (c *RESTFrontend) UserStats(w http.ResponseWriter, r *http.Request) {
	content := c.ContentController.BuildUserStatsContent()
	c.RenderTemplate(w, StatsUser, content)
}

func (c *RESTFrontend) ProductsProjectsStats(w http.ResponseWriter, r *http.Request) {
	content := c.ContentController.BuildItemStatsContent()
	c.RenderTemplate(w, StatsProductProject, content)
}

func (c *RESTFrontend) UIStats(w http.ResponseWriter, r *http.Request) {
	content := c.ContentController.BuildUIStatsContent()
	c.RenderTemplate(w, StatsUI, content)
}

func (c *RESTFrontend) MisuseMetrics(w http.ResponseWriter, r *http.Request) {
	content := make(map[string]interface{})
	c.RenderTemplate(w, StatsMisuseMetrics, content)
}

func (c *RESTFrontend) AccountingStats(w http.ResponseWriter, r *http.Request) {
	content := c.ContentController.BuildAccountingStatsContent()
	c.RenderTemplate(w, StatsAccounting, content)
}

func (c *RESTFrontend) SystemHealthStats(w http.ResponseWriter, r *http.Request) {
	content := c.ContentController.BuildSystemHealthContent()
	c.RenderTemplate(w, StatsSystemHealth, content)
}
