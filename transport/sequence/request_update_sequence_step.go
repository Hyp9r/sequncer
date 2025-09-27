package sequence

type UpdateSequenceStepRequest struct {
	Subject *string `json:"subject,omitempty"`
	Content *string `json:"content,omitempty"`
}
