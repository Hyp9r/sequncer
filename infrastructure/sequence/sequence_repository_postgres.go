package sequence

import (
	"database/sql"

	"github.com/Hyp9r/sequncer/domain/sequence"
	"github.com/rs/zerolog"
)

type SequenceRepository struct {
	db     *sql.DB
	logger *zerolog.Logger
}

func NewSequenceRepository(db *sql.DB, logger *zerolog.Logger) *SequenceRepository {
	return &SequenceRepository{
		db:     db,
		logger: logger,
	}
}

func (sr *SequenceRepository) Persist(sequence *sequence.Sequence) error {
	query := `INSERT INTO sequences (name, open_tracking_enabled, click_tracking_enabled) VALUES ($1, $2, $3)`
	_, err := sr.db.Exec(query, sequence.Name, sequence.OpenTrackingEnabled, sequence.ClickTrackingEnabled)
	if err != nil {
		sr.logger.Error().Err(err).Msg("Failed to persist sequence")
		return err
	}
	return nil
}
