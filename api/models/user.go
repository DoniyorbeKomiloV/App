package models

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Phone     string `json:"phone"`
	Picture   string `json:"picture"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CardNo    string `json:"card_no"`
}

type CreateUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Phone     string `json:"phone"`
	Picture   string `json:"picture"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CardNo    string `json:"card_no"`
}

type UpdateUser struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Phone     string `json:"phone"`
	Picture   string `json:"picture"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CardNo    string `json:"card_no"`
}

type UserGetListRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type UserGetListResponse struct {
	Count int     `json:"count"`
	Users []*User `json:"users"`
}

type UserPrimaryKey struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
