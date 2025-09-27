package sequence

import "errors"

type sequencePersister interface {
	Persist(sequence *Sequence) error
}

type SequenceService struct {
	persister sequencePersister
}

func NewSequenceService(persister sequencePersister) *SequenceService {
	return &SequenceService{
		persister: persister,
	}
}

func (s *SequenceService) CreateSequence(sequence *Sequence) error {
	if sequence.Name == "" {
		return errors.New("sequence name cannot be empty")
	}
	return s.persister.Persist(sequence)
}
