package sequence

import (
	"net/http"

	"github.com/Hyp9r/sequncer/domain/sequence"
)

type Controller struct {
	router          *http.ServeMux
	sequenceService *sequence.SequenceService
}

func NewController(router *http.ServeMux, sequenceService *sequence.SequenceService) *Controller {
	ctrl := &Controller{
		router:          router,
		sequenceService: sequenceService,
	}
	ctrl.registerRoutes()
	return ctrl
}

func (c *Controller) registerRoutes() {
	c.router.HandleFunc("POST /sequences", c.createSequenceHandler)
	c.router.HandleFunc("PATCH /sequences/{ID}/steps/{stepID}", c.updateSequenceStepHandler)
	c.router.HandleFunc("DELETE /sequences/{ID}/steps/{stepID}", c.deleteStepHandler)
	c.router.HandleFunc("PATCH /sequences/{ID}/tracking/open", c.updateOpenTrackingHandler)
	c.router.HandleFunc("PATCH /sequences/{ID}/tracking/click", c.updateClickTrackingHandler)
}
