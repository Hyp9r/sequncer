package sequence

import (
	"errors"

	"github.com/google/uuid"
)

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

func (s *Sequence) UpdateStep(stepID string, subject, content *string) error {
	step := s.findStep(stepID)
	if step == nil {
		return errors.New("step not found")
	}

	if subject != nil {
		step.Subject = *subject
	}
	if content != nil {
		step.Content = *content
	}

	return nil
}

func (s *Sequence) findStep(stepID string) *Step {
	for i := range s.Steps {
		if s.Steps[i].ID == stepID {
			return &s.Steps[i]
		}
	}
	return nil
}
