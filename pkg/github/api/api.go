package api

import (
	"net/http"
	"time"

	"github.com/chentex/github-project-mgr/pkg/config"
	"github.com/chentex/github-project-mgr/pkg/github/model"
)

// Interface definition of the API
type Interface interface {
	MoveIssueInProgress(issue *model.Issue, w http.ResponseWriter)
	ClosedIssue(issue *model.Issue, w http.ResponseWriter)
	LabelActions(issue *model.Issue, w http.ResponseWriter)
	UnlabelActions(issue *model.Issue, w http.ResponseWriter)
	MoveIssueBlocked(issue *model.Issue, w http.ResponseWriter)
	ListenAndServe() error
}

// API to connect to github
type API struct {
	HTTPServer *http.Server
	Config     *config.Config
}

// NewAPI returns new instance of the API
func NewAPI(c *config.Config) Interface {
	s := &http.Server{
		Addr:         c.Server.ServerAddr,
		ReadTimeout:  time.Duration(c.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(c.Server.WriteTimeout) * time.Second,
	}
	a := &API{
		HTTPServer: s,
		Config:     c,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/v0/issue", a.IssueHandler)
	mux.HandleFunc("/v0/version", a.VersionHandler)
	a.HTTPServer.Handler = mux
	return a
}

// ListenAndServe wrapper for stating the server
func (a *API) ListenAndServe() error {
	return a.HTTPServer.ListenAndServe()
}
