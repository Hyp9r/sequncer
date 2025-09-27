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
}
