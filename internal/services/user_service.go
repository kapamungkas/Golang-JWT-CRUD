package services

import (
	entities "betest/internal/entities"
	ru "betest/internal/repositories"
	"betest/internal/requests"
	"betest/internal/responses"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetAllUser(option map[string]string) ([]responses.GetUserResponse, error)
	FindUserByID(id string) (responses.GetUserResponse, error)
	CreateUser(dataUser entities.UserEntity) (responses.CreateUserResponse, error)
	UpdateUser(id string, dataUser entities.UserEntity) (responses.UpdateUserResponse, error)
	DeleteUser(id string) (bool, error)
	ChangePassword(requestPassword requests.ChangePasswordRequest, id string) error
}

type userService struct {
	ru ru.UserRepository
}

func NewUserService(r ru.UserRepository) *userService {
	return &userService{
		r,
	}
}

func (s userService) GetAllUser(option map[string]string) ([]responses.GetUserResponse, error) {
	result, err := s.ru.GetAllUser(option)
	return result, err
}

func (s userService) FindUserByID(id string) (responses.GetUserResponse, error) {
	result, err := s.ru.FindUserByID(id)
	return result, err
}

func (s userService) CreateUser(dataUser entities.UserEntity) (responses.CreateUserResponse, error) {
	result, err := s.ru.CreateUser(dataUser)
	return result, err
}

func (s userService) UpdateUser(id string, dataUser entities.UserEntity) (responses.UpdateUserResponse, error) {

	var result responses.UpdateUserResponse
	var err error

	if dataUser.ProfilePicture == "" {
		result, err = s.ru.UpdateUserWithoutImage(id, dataUser)
	} else {
		result, err = s.ru.UpdateUser(id, dataUser)
	}
	return result, err
}

func (s userService) DeleteUser(id string) (bool, error) {

	result, err := s.ru.DeleteUser(id)
	return result, err
}

func (s userService) ChangePassword(requestPassword requests.ChangePasswordRequest, id string) error {
	user, err := s.ru.FindUserByIDGetPassword(id)
	if err != nil {
		return err
	}
	fmt.Println(requestPassword.CurrentPassword)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestPassword.CurrentPassword))

	if err != nil {
		return errors.New("please check your current password")
	}

	err = s.ru.ChangePassword(requestPassword, id)

	if err != nil {
		return errors.New("failed change password")
	}

	return nil
}
