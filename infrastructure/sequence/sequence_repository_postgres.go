package sequence

import (
	"context"
	"database/sql"
	"errors"

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

var _ sequence.SequenceRepository = (*SequenceRepository)(nil)

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

func (sr *SequenceRepository) GetByID(ctx context.Context, ID string) (*sequence.Sequence, error) {
	rows, err := sr.db.QueryContext(ctx, `
		SELECT
			s.id AS sequence_id,
			s.name,
			s.open_tracking_enabled,
			s.click_tracking_enabled,
			st.id AS step_id,
			st.step_number,
			st.subject,
			st.content
		FROM sequences s
		LEFT JOIN sequence_steps st ON st.sequence_id = s.id
		WHERE s.id = $1
		ORDER BY st.step_number ASC
	`, ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var seq *sequence.Sequence
	steps := []sequence.Step{}

	for rows.Next() {
		var (
			sequenceID string
			name       string
			openTrack  bool
			clickTrack bool
			stepID     sql.NullString
			stepOrder  sql.NullInt32
			subject    sql.NullString
			content    sql.NullString
		)

		if err := rows.Scan(&sequenceID, &name, &openTrack, &clickTrack,
			&stepID, &stepOrder, &subject, &content); err != nil {
			return nil, err
		}

		if seq == nil {
			seq = &sequence.Sequence{
				ID:                   sequenceID,
				Name:                 name,
				OpenTrackingEnabled:  openTrack,
				ClickTrackingEnabled: clickTrack,
			}
		}

		if stepID.Valid {
			steps = append(steps, sequence.Step{
				ID:         stepID.String,
				StepNumber: int(stepOrder.Int32),
				Subject:    subject.String,
				Content:    content.String,
			})
		}
	}

	if seq == nil {
		return nil, errors.New("sequence not found")
	}

	seq.Steps = steps
	return seq, nil
}

func (r *SequenceRepository) UpdateStep(ctx context.Context, sequenceID string, step sequence.Step) error {
	res, err := r.db.ExecContext(ctx,
		`UPDATE sequence_steps
		 SET subject = COALESCE($1, subject),
		     content = COALESCE($2, content)
		 WHERE id = $3 AND sequence_id = $4`,
		step.Subject,
		step.Content,
		step.ID,
		sequenceID,
	)
	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return errors.New("step not found")
	}
	return nil
}

func (sr *SequenceRepository) DeleteStep(ctx context.Context, sequenceID, stepID string) error {
	res, err := sr.db.ExecContext(ctx,
		"DELETE FROM sequence_steps WHERE id=$1 AND sequence_id=$2",
		stepID, sequenceID,
	)
	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return errors.New("step not found")
	}
	return nil
}

func (sr *SequenceRepository) UpdateTracking(ctx context.Context, seq *sequence.Sequence) error {
	res, err := sr.db.ExecContext(ctx,
		`UPDATE sequences
		 SET open_tracking_enabled = $1,
		     click_tracking_enabled = $2
		 WHERE id = $3`,
		seq.OpenTrackingEnabled,
		seq.ClickTrackingEnabled,
		seq.ID,
	)
	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return errors.New("sequence not found")
	}
	return nil
}
