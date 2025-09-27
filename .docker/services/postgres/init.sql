CREATE TABLE sequences (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    open_tracking_enabled BOOLEAN DEFAULT FALSE,
    click_tracking_enabled BOOLEAN DEFAULT FALSE
);

CREATE TABLE steps (
    id SERIAL PRIMARY KEY,
    sequence_id INTEGER REFERENCES sequences(id) ON DELETE CASCADE,
    step_number INTEGER NOT NULL,
    content TEXT NOT NULL,
    subject VARCHAR(255) NOT NULL
);