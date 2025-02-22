
CREATE TABLE IF NOT EXISTS urls (
    id SERIAL PRIMARY KEY,
    short_url TEXT NOT NULL, 
    long_url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    total_clicks INTEGER DEFAULT 0
);
