package sequence

import (
	"context"
	"errors"
)

type sequencePersister interface {
	Persist(ctx context.Context, sequence *Sequence) error
}

type SequenceService struct {
	persister sequencePersister
}

func NewSequenceService(persister sequencePersister) *SequenceService {
	return &SequenceService{
		persister: persister,
	}
}

func (s *SequenceService) Create(ctx context.Context, cmd CreateSequenceCommand) error {
	if cmd.Name == "" {
		return errors.New("sequence name cannot be empty")
	}
	steps := make([]Step, len(cmd.Steps))
	for i, st := range cmd.Steps {
		if st.Subject == "" && st.Content == "" {
			return errors.New("step must have subject or content")
		}
		steps[i] = Step{
			ID:         "",
			Subject:    st.Subject,
			Content:    st.Content,
			StepNumber: i,
		}
	}
	sequence := NewSequence(cmd.Name, cmd.OpenTrackingEnabled, cmd.ClickTrackingEnabled, steps)
	return s.persister.Persist(ctx, sequence)
}
