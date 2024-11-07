package domain

type User struct {
	Id       uint   `json:"-"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
	RoleId   int    `json:"-"`
}
