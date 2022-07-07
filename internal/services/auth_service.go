package services

import (
	"betest/internal/helpers"
	"betest/internal/repositories"
	"betest/internal/requests"
	"betest/internal/responses"
	"errors"
	"os"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(dataAuth requests.LoginRequest) (responses.LoginResponse, error)
	UpdateRefreshToken(id string, refreshToken string) error
	CheckRefreshTokenOnDatabase(id string, refreshToken string) (responses.LoginResponse, error)
	ResetPassword(email string, password string) error
	GeneratePassword(token string, password string, gen_password string) error
}

type authService struct {
	repo repositories.AuthRepository
}

func NewAuthService(r repositories.AuthRepository) *authService {
	return &authService{
		r,
	}
}

func (s authService) Login(dataAuth requests.LoginRequest) (responses.LoginResponse, error) {
	doLogin, err := s.repo.Login(dataAuth)

	if err != nil {
		return doLogin, err
	}
	//check password
	err = bcrypt.CompareHashAndPassword([]byte(doLogin.Password), []byte(dataAuth.Password))

	if err != nil {
		return doLogin, errors.New("please check your password")
	}
	return doLogin, nil
}

func (s authService) UpdateRefreshToken(id string, refreshToken string) error {
	err := s.repo.UpdateRefreshToken(id, refreshToken)
	return err
}

func (s authService) CheckRefreshTokenOnDatabase(id string, refreshToken string) (responses.LoginResponse, error) {
	result, err := s.repo.CheckRefreshTokenOnDatabase(id, refreshToken)
	return result, err
}

func (s authService) ResetPassword(email string, token string) error {
	user, err := s.repo.CheckEmailUserOnDatabase(email)
	if err != nil {
		return errors.New("your email is not registred on system")
	}

	err = s.repo.UpdateResetPassword(user.ID, token)
	if err != nil {
		return err
	}
	base_url := os.Getenv("BASE_URL")
	generate_url := base_url + "/api/auth/generate-password?token=" + token
	subject := "Reset Password"
	message := "Please click this url to generate new password : " + generate_url

	err = helpers.SMTPEmail(email, subject, message)
	if err != nil {
		return err
	}
	return nil

}

func (s authService) GeneratePassword(token string, password string, gen_password string) error {
	user, err := s.repo.CheckTokenUser(token)
	if err != nil {
		return errors.New("Your token is invalid")
	}

	err = s.repo.UpdatePasswordUser(user.ID, gen_password)
	if err != nil {
		return errors.New("Failed on update user")
	}

	subject := "Reset Password"
	message := "Your new password is : " + password

	err = helpers.SMTPEmail(user.Email, subject, message)
	if err != nil {
		return err
	}
	return nil

}
