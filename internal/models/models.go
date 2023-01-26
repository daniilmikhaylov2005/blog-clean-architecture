package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Post struct {
	ID     int    `json:"id"`
	UserId int    `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type PostComment struct {
	Body string `json:"body"`
}

type SuccessResponse struct {
	Status string `json:"status"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}
