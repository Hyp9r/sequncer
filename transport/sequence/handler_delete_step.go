package sequence

import (
	"net/http"

	"github.com/Hyp9r/sequncer/domain/sequence"
)

func (c *Controller) deleteStepHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	sequenceID := r.PathValue("ID")
	stepID := r.PathValue("stepID")

	cmd := sequence.DeleteStepCommand{
		SequenceID: sequenceID,
		StepID:     stepID,
	}

	if err := c.sequenceService.DeleteStep(ctx, cmd); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
