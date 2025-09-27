package sequence

import "github.com/google/uuid"

type Sequence struct {
	ID                   string
	Name                 string
	OpenTrackingEnabled  bool
	ClickTrackingEnabled bool
	Steps                []Step
}

type Step struct {
	ID         string
	Subject    string
	Content    string
	StepNumber int
}

func NewSequence(name string, openTrackingEnabled bool, clickTrackingEnabled bool, steps []Step) *Sequence {
	if len(steps) > 0 {
		for i := range steps {
			if steps[i].ID == "" {
				steps[i].ID = uuid.NewString()
			}
			steps[i].StepNumber = i
		}
	}
	return &Sequence{
		ID:                   uuid.NewString(),
		Name:                 name,
		OpenTrackingEnabled:  openTrackingEnabled,
		ClickTrackingEnabled: clickTrackingEnabled,
		Steps:                steps,
	}
}
