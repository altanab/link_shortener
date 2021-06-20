CREATE TABLE IF NOT EXISTS links (
    id SERIAL PRIMARY KEY,
    original_link VARCHAR(256) UNIQUE,
    shorten_link VARCHAR(10) UNIQUE
);