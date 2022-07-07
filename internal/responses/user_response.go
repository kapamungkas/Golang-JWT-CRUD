package responses

type GetUserResponse struct {
	ID             string
	Username       string
	FirstName      string
	LastName       string
	Email          string
	Phone          string
	UserRole       string
	ProfilePicture string
	CreatedAt      string
	UpdatedAt      string
}

type CreateUserResponse struct {
	ID             string
	Username       string
	FirstName      string
	LastName       string
	Email          string
	Phone          string
	ProfilePicture string
}

type UpdateUserResponse struct {
	ID             string
	Username       string
	FirstName      string
	LastName       string
	Email          string
	Phone          string
	ProfilePicture string
}

type GetPasswordUserResponse struct {
	Password string
}
