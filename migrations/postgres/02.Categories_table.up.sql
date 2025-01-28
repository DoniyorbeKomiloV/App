CREATE TABLE categories(
    id uuid PRIMARY KEY,
    name VARCHAR NOT NULL,
    type VARCHAR NOT NULL,
    picture VARCHAR,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    is_deleted BOOLEAN DEFAULT FALSE
)