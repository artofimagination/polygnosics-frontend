package frontend

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

func (c *RESTController) StatsWebRTC(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		c.HandleError(w, fmt.Sprintf("Failed to parse frontend webrtc offer. %s", errors.WithStack(err)), http.StatusInternalServerError, c.URI(UserMain))
		return
	}

	if err := c.RESTBackend.InitStatsWebRTC(r); err != nil {
		c.HandleError(w, fmt.Sprintf("Failed to init webrtc data provider. %s", errors.WithStack(err)), http.StatusInternalServerError, c.URI(UserMain))
		return
	}
}

func (c *RESTController) ProductStats(w http.ResponseWriter, r *http.Request) {
	content := c.ContentController.BuildProductStatsContent()
	c.RenderTemplate(w, StatsProduct, content)
}

func (c *RESTController) ProjectStats(w http.ResponseWriter, r *http.Request) {
	content := c.ContentController.BuildProjectStatsContent()
	c.RenderTemplate(w, StatsProject, content)
}

func (c *RESTController) UserStats(w http.ResponseWriter, r *http.Request) {
	content := c.ContentController.BuildUserStatsContent()
	c.RenderTemplate(w, StatsUser, content)
}

func (c *RESTController) ProductsProjectsStats(w http.ResponseWriter, r *http.Request) {
	content := c.ContentController.BuildItemStatsContent()
	c.RenderTemplate(w, StatsProductProject, content)
}

func (c *RESTController) UIStats(w http.ResponseWriter, r *http.Request) {
	content := c.ContentController.BuildUIStatsContent()
	c.RenderTemplate(w, StatsUI, content)
}

func (c *RESTController) MisuseMetrics(w http.ResponseWriter, r *http.Request) {
	content := make(map[string]interface{})
	c.RenderTemplate(w, StatsMisuseMetrics, content)
}

func (c *RESTController) AccountingStats(w http.ResponseWriter, r *http.Request) {
	content := c.ContentController.BuildAccountingStatsContent()
	c.RenderTemplate(w, StatsAccounting, content)
}

func (c *RESTController) SystemHealthStats(w http.ResponseWriter, r *http.Request) {
	content := c.ContentController.BuildSystemHealthContent()
	c.RenderTemplate(w, StatsSystemHealth, content)
}
