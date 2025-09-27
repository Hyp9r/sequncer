package sequence

type CreateSequenceCommand struct {
	Name                 string
	OpenTrackingEnabled  bool
	ClickTrackingEnabled bool
	Steps                []CreateStep
}

type CreateStep struct {
	Subject string
	Content string
}

type UpdateStepCommand struct {
	SequenceID string
	StepID     string
	Subject    *string
	Content    *string
}
