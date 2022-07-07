package requests

type LoginRequest struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type RefreshToken struct {
	Token string `form:"refresh_token" binding:"required"`
}

type ResetPasswordRequest struct {
	Email string `form:"email" binding:"required,email"`
}
