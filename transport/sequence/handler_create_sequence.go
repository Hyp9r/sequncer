package sequence

import (
	"encoding/json"
	"net/http"

	"github.com/Hyp9r/sequncer/domain/sequence"
)

func (c *Controller) createSequenceHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req CreateSequenceRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	cmd := sequence.CreateSequenceCommand{
		Name:                 req.Name,
		OpenTrackingEnabled:  req.OpenTrackingEnabled,
		ClickTrackingEnabled: req.ClickTrackingEnabled,
		Steps:                make([]sequence.CreateStep, len(req.Steps)),
	}

	for i, s := range req.Steps {
		cmd.Steps[i] = sequence.CreateStep{
			Subject: s.Subject,
			Content: s.Content,
		}
	}

	err = c.sequenceService.Create(ctx, cmd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
