package domain

type RegisterRequest struct {
	Id       uint   `json:"id" gorm:"primaryKey;size:256"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleId   int    `json:"-"`
}

func NewRegisterRequest(id uint, username string, email string, password string, description string, roleId int) *RegisterRequest {
	return &RegisterRequest{
		Id:       id,
		Username: username,
		Email:    email,
		Password: password,
		RoleId:   roleId,
	}
}
