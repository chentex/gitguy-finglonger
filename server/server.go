package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/SUSE/gitguy-finglonger/pkg/caasp"
	"github.com/SUSE/gitguy-finglonger/pkg/config"
	"github.com/SUSE/gitguy-finglonger/pkg/github/model"
	"github.com/SUSE/gitguy-finglonger/pkg/security"
	"github.com/SUSE/gitguy-finglonger/version"
)

// ServerInterface definition of the Server
type ServerInterface interface {
	ListenAndServe() error
}

// Server to connect to github
type Server struct {
	HTTPServer *http.Server
	Config     *config.Config
}

// NewServer returns new instance of the Server
func NewServer(s *http.Server, c *config.Config) ServerInterface {
	a := &Server{
		HTTPServer: s,
		Config:     c,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/v0/webhooks", a.webHooksHandler)
	mux.HandleFunc("/v0/version", version.Handler)
	a.HTTPServer.Handler = mux
	return a
}

// ListenAndServe wrapper for stating the server
func (a *Server) ListenAndServe() error {
	return a.HTTPServer.ListenAndServe()
}

// WebHooksHandler handles webhooks events from github
func (a *Server) webHooksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	signature := r.Header.Get("X-Hub-Signature")
	if !security.IsValidSignature(body, signature, a.Config.Github.Secret) {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	webHook := new(model.WebHook)
	json.Unmarshal(body, webHook)
	log.Printf("process action from Repository: %s", webHook.Repository.Name)

	switch webHook.Repository.Name {
	case a.Config.Caasp.Repositories[0], a.Config.Caasp.Repositories[1], a.Config.Caasp.Repositories[2], a.Config.Caasp.Repositories[3]:
		caasp := caasp.NewCaaspProcessor(a.Config)
		caasp.ProcessWebHook(body, w, r)
	}
}
