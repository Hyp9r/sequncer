package sequence

type CreateSequenceRequest struct {
	Name                 string `json:"name"`
	OpenTrackingEnabled  bool   `json:"open_tracking_enabled"`
	ClickTrackingEnabled bool   `json:"click_tracking_enabled"`
}
