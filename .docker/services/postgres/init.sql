CREATE TABLE sequences (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    open_tracking_enabled BOOLEAN DEFAULT FALSE,
    click_tracking_enabled BOOLEAN DEFAULT FALSE
);