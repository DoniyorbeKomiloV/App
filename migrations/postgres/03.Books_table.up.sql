CREATE TABLE books(
    id uuid PRIMARY KEY,
    title VARCHAR NOT NULL,
    author VARCHAR NOT NULL,
    publisher VARCHAR NOT NULL,
    category uuid REFERENCES categories(id),
    num_pages INT NOT NULL,
    picture VARCHAR NOT NULL,
    lang VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    is_deleted BOOLEAN DEFAULT FALSE
)