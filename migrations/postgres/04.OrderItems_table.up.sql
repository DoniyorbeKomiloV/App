CREATE TABLE order_items(
    item_id uuid PRIMARY KEY,
    order_id uuid REFERENCES orders(order_id),
    book_id uuid REFERENCES books(id)
)