package models

type Book struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	Category  string `json:"category"`
	NumPages  int    `json:"num_pages"`
	Lang      string `json:"lang"`
}

type CreateBook struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	Category  string `json:"category"`
	NumPages  int    `json:"num_pages"`
	Lang      string `json:"lang"`
}

type UpdateBook struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	Category  string `json:"category"`
	NumPages  int    `json:"num_pages"`
	Lang      string `json:"lang"`
}

type BookGetListRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type BookGetListResponse struct {
	Count int     `json:"count"`
	Books []*Book `json:"books"`
}

type BookPrimaryKey struct {
	Id string `json:"id"`
}
