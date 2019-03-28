package caasp

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/SUSE/gitguy-finglonger/pkg/config"
	"github.com/SUSE/gitguy-finglonger/pkg/github/model"
)

// Caasp team workflow for webhooks
type Caasp interface {
	ProcessWebHook(b []byte, w http.ResponseWriter, r *http.Request)
}

// NewCaaspProcessor returns a new caasp processor
func NewCaaspProcessor(c *config.Config) Caasp {
	return &Processor{
		config: c,
	}
}

// Processor processes the webhooks send to this worker
type Processor struct {
	config *config.Config
}

// ProcessWebHook resolve all possible actions sent through the webhook
func (p *Processor) ProcessWebHook(b []byte, w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	event := r.Header.Get("X-GitHub-Event")

	switch event {
	case "issue":
		issue := new(model.Issue)
		err := json.Unmarshal(b, issue)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("ERROR: %s", err)
			return
		}
		log.Printf("process action: %s", issue.Action)
		for _, r := range p.config.Caasp.ProjectRules.MoveCards {

		}

	}
}
