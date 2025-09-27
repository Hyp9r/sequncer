package sequence

import (
	"context"
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

func (sr *SequenceRepository) Persist(ctx context.Context, sequence *sequence.Sequence) error {
	tx, err := sr.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	query := `INSERT INTO sequences (id, name, open_tracking_enabled, click_tracking_enabled) VALUES ($1, $2, $3, $4)`

	_, err = tx.ExecContext(ctx, query, sequence.ID, sequence.Name, sequence.OpenTrackingEnabled, sequence.ClickTrackingEnabled)
	if err != nil {
		sr.logger.Error().Err(err).Msg("failed to persist sequence")
		return err
	}

	stepQuery := `INSERT INTO sequence_steps (id, sequence_id, subject, content, step_number) VALUES ($1, $2, $3, $4, $5)`
	for _, s := range sequence.Steps {
		_, err := tx.ExecContext(ctx, stepQuery, s.ID, sequence.ID, s.Subject, s.Content, s.StepNumber)
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}
