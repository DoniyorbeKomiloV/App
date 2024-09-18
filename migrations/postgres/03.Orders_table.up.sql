CREATE TABLE orders(
    order_id uuid PRIMARY KEY,
    user_id uuid REFERENCES users(id)
)
