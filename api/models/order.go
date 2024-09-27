package models

type Order struct {
	OrderId string `json:"order_id"`
	UserId  string `json:"user_id"`
}

type CreateOrder struct {
	UserId string `json:"user_id"`
}

type UpdateOrder struct {
	OrderId string `json:"order_id"`
	UserId  string `json:"user_id"`
}

type OrderGetListRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type OrderGetListResponse struct {
	Count  int      `json:"count"`
	Orders []*Order `json:"orders"`
}

type OrderPrimaryKey struct {
	OrderId string `json:"order_id"`
}
