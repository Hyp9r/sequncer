CREATE TABLE sequences (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    open_tracking_enabled BOOLEAN DEFAULT FALSE,
    click_tracking_enabled BOOLEAN DEFAULT FALSE
);

CREATE TABLE sequence_steps (
    id UUID PRIMARY KEY,
    sequence_id UUID NOT NULL REFERENCES sequences(id) ON DELETE CASCADE,
    step_number INTEGER NOT NULL,
    content TEXT NOT NULL,
    subject VARCHAR(255) NOT NULL,
    CONSTRAINT unique_sequence_step_order UNIQUE(sequence_id, step_number)
);