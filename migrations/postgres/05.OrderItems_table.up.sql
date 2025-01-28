CREATE TABLE order_items(
    item_id uuid PRIMARY KEY,
    order_id uuid REFERENCES orders(order_id),
    book_id uuid REFERENCES books(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    is_deleted BOOLEAN DEFAULT FALSE
)