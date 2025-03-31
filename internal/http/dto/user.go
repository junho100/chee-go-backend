package dto

type SignUpRequest struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	ID string `json:"id"`
}

type CreateUserDto struct {
	ID       string
	Email    string
	Password string
}

type CheckIDResponse struct {
	IsExists bool `json:"is_exists"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type LoginRequest struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

type CheckMeRequest struct {
	Token string `json:"token"`
}

type CheckMeResponse struct {
	UserID string `json:"user_id"`
}
