CREATE TABLE orders(
    order_id uuid PRIMARY KEY,
    user_id uuid REFERENCES users(id),
    status VARCHAR,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    is_deleted BOOLEAN DEFAULT FALSE
)
