package models

type OrderItem struct {
	ItemId  string `json:"item_id"`
	OrderId string `json:"order_id"`
	BookId  string `json:"book_id"`
}

type CreateOrderItem struct {
	OrderId string `json:"order_id"`
	BookId  string `json:"book_id"`
}

type UpdateOrderItem struct {
	ItemId  string `json:"item_id"`
	OrderId string `json:"order_id"`
	BookId  string `json:"book_id"`
}

type OrderItemGetListRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type OrderItemGetListResponse struct {
	Count      int          `json:"count"`
	OrderItems []*OrderItem `json:"order_items"`
}

type OrderItemPrimaryKey struct {
	ItemId string `json:"item_id"`
}
