CREATE TABLE links (
    id SERIAL PRIMARY KEY,
    short_url VARCHAR(255) NOT NULL,
    original_url VARCHAR(255) NOT NULL,
    user_id VARCHAR(255) NOT NULL
)