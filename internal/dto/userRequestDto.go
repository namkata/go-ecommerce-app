package dto

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type UserSignup struct {
	UserLogin
	Phone string `json:"phone"`
}

type VerificationCodeInput struct {
	Code int `json:"code"`
}
