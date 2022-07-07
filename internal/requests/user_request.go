package requests

type CreateUserRequest struct {
	Username  string `form:"username" binding:"required,min=4"`
	Password  string `form:"password" binding:"required,min=4"`
	FirstName string `form:"first_name" binding:"required"`
	LastName  string `form:"last_name" binding:"required"`
	Email     string `form:"email" binding:"required,email"`
	Phone     string `form:"phone" binding:"required"`
	UserRole  string `form:"user_role" binding:"required"`
}

type UpdateUserRequest struct {
	FirstName string `form:"first_name" binding:"required"`
	LastName  string `form:"last_name" binding:"required"`
	Email     string `form:"email" binding:"required,email"`
	Phone     string `form:"phone" binding:"required"`
}

type ChangePasswordRequest struct {
	CurrentPassword    string `form:"current_password" binding:"required,min=4"`
	NewPassword        string `form:"new_password" binding:"required,min=4"`
	ConfirmNewPassword string `form:"confirm_new_password" binding:"required,min=4"`
}
