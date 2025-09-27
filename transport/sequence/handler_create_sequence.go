package sequence

import (
	"encoding/json"
	"net/http"

	"github.com/Hyp9r/sequncer/domain/sequence"
)

func (c *Controller) createSequenceHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateSequenceRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	sequence := sequence.NewSequence(req.Name, req.OpenTrackingEnabled, req.ClickTrackingEnabled)

	err = c.sequenceService.CreateSequence(sequence)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
