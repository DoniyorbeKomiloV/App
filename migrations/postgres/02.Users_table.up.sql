CREATE TABLE users(
    id uuid PRIMARY KEY,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    age INT,
    phone VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    card_no VARCHAR,
    is_deleted BOOLEAN DEFAULT FALSE
)