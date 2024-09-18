CREATE TABLE books(
     id uuid PRIMARY KEY,
     title VARCHAR NOT NULL,
     author VARCHAR NOT NULL,
     publisher VARCHAR NOT NULL,
     category VARCHAR NOT NULL,
     num_pages INT NOT NULL,
     lang VARCHAR NOT NULL,
     is_deleted BOOLEAN DEFAULT FALSE
)