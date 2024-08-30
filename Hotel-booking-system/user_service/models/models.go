package models

type RegisterUserRequest struct {
	Username string `json:"username"`
	Age      int32  `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GeneralResponse struct {
	Message string `json:"message"`
}

type VerifyRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LogInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetUserRequest struct {
	ID int32 `json:"id"`
}

type LogInResponse struct {
	Status bool `json:"status"`
}

type LastInsertedUser struct{}

type GetUserResponse struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
	Age      int32  `json:"age"`
	Email    string `json:"email"`
	LogOut   bool   `json:"logout"`
}

type UpdateUserRequest struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
	Age      int32  `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
