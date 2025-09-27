package sequence

import (
	"encoding/json"
	"net/http"

	"github.com/Hyp9r/sequncer/domain/sequence"
)

func (c *Controller) updateSequenceStepHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	sequenceID := r.PathValue("ID")
	stepID := r.PathValue("stepID")

	var req UpdateSequenceStepRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	cmd := sequence.UpdateStepCommand{
		SequenceID: sequenceID,
		StepID:     stepID,
		Subject:    req.Subject,
		Content:    req.Content,
	}

	if err := c.sequenceService.UpdateStep(ctx, cmd); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
