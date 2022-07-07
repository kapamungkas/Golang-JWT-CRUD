package responses

type LoginResponse struct {
	ID       string
	Username string
	Password string
	UserRole string
	Email    string
}

type RefreshTokenResponse struct {
	RefreshToken string
}
