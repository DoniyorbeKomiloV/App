package models

type Category struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Picture string `json:"picture"`
}

type CreateCategory struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Picture string `json:"picture"`
}

type UpdateCategory struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Picture string `json:"picture"`
}

type CategoryGetListRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type CategoryGetListResponse struct {
	Count      int         `json:"count"`
	Categories []*Category `json:"categorys"`
}

type CategoryPrimaryKey struct {
	Id string `json:"id"`
}
