package entities

type UserEntity struct {
	ID             string
	Username       string
	Password       string
	FirstName      string
	LastName       string
	Email          string
	Phone          string
	UserRole       string
	ProfilePicture string
	IsDeleted      int
	CreatedAt      string
	UpdatedAt      string
}
