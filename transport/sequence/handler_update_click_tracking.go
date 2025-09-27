package sequence

import (
	"encoding/json"
	"net/http"
)

func (c *Controller) updateClickTrackingHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	sequenceID := r.PathValue("ID")

	var req UpdateTrackingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.sequenceService.UpdateClickTracking(ctx, sequenceID, req.Enabled); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
