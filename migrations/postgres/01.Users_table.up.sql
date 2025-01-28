CREATE TABLE users(
    id uuid PRIMARY KEY,
    first_name VARCHAR,
    last_name VARCHAR,
    age INT,
    phone VARCHAR,
    picture VARCHAR,
    username VARCHAR NOT NULL UNIQUE,
    password VARCHAR NOT NULL,
    card_no VARCHAR,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    is_deleted BOOLEAN DEFAULT FALSE
)