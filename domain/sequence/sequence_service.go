package sequence

import (
	"context"
	"errors"
)

type SequencePersister interface {
	Persist(ctx context.Context, sequence *Sequence) error
	UpdateStep(ctx context.Context, sequenceID string, step Step) error
	DeleteStep(ctx context.Context, sequenceID, stepID string) error
}

type SequenceReader interface {
	GetByID(ctx context.Context, id string) (*Sequence, error)
}

type SequenceRepository interface {
	SequencePersister
	SequenceReader
}

type SequenceService struct {
	repo SequenceRepository
}

func NewSequenceService(repo SequenceRepository) *SequenceService {
	return &SequenceService{
		repo: repo,
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
	return s.repo.Persist(ctx, sequence)
}

func (s *SequenceService) UpdateStep(ctx context.Context, cmd UpdateStepCommand) error {
	seq, err := s.repo.GetByID(ctx, cmd.SequenceID)
	if err != nil {
		return err
	}

	if err := seq.UpdateStep(cmd.StepID, cmd.Subject, cmd.Content); err != nil {
		return err
	}

	step := seq.findStep(cmd.StepID)
	if step == nil {
		return errors.New("step not found in sequence")
	}

	return s.repo.UpdateStep(ctx, seq.ID, *step)
}

func (s *SequenceService) DeleteStep(ctx context.Context, cmd DeleteStepCommand) error {
	seq, err := s.repo.GetByID(ctx, cmd.SequenceID)
	if err != nil {
		return err
	}

	if err := seq.DeleteStep(cmd.StepID); err != nil {
		return err
	}

	return s.repo.DeleteStep(ctx, seq.ID, cmd.StepID)
}
