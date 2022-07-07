package repositories

import (
	entities "betest/internal/entities"
	"betest/internal/helpers"
	"betest/internal/requests"
	"betest/internal/responses"
	"database/sql"
	"errors"
	"fmt"
)

type UserRepository interface {
	GetAllUser(option map[string]string) ([]responses.GetUserResponse, error)
	FindUserByID(id string) (responses.GetUserResponse, error)
	CreateUser(dataUser entities.UserEntity) (responses.CreateUserResponse, error)
	UpdateUser(id string, dataUser entities.UserEntity) (responses.UpdateUserResponse, error)
	UpdateUserWithoutImage(id string, dataUser entities.UserEntity) (responses.UpdateUserResponse, error)
	DeleteUser(id string) (bool, error)
	ChangePassword(requestPassword requests.ChangePasswordRequest, id string) error
	FindUserByIDGetPassword(id string) (responses.GetPasswordUserResponse, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db}
}

func (r userRepository) GetAllUser(option map[string]string) ([]responses.GetUserResponse, error) {
	additional_query := ""
	if len(option) != 0 {
		additional_query = helpers.CreateQueryOrderAndLimit(option)
	}

	rows, err := r.db.Query("SELECT id, username, firstname, lastname, email, phone, user_role, profile_picture,created_at,updated_at FROM users where is_deleted = 0 " + additional_query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// An user slice to hold data from returned rows.
	var users []responses.GetUserResponse

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var user responses.GetUserResponse
		if err := rows.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.Email, &user.Phone, &user.UserRole, &user.ProfilePicture, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return users, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return users, err
	}
	return users, nil
}

func (r userRepository) FindUserByID(id string) (responses.GetUserResponse, error) {
	var user = responses.GetUserResponse{}
	err := r.db.
		QueryRow("select id, username, firstname, lastname, email, phone, user_role, profile_picture,created_at,updated_at from users where id = ?", id).Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.Email, &user.Phone, &user.UserRole, &user.ProfilePicture, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r userRepository) CreateUser(dataUser entities.UserEntity) (responses.CreateUserResponse, error) {
	var user = responses.GetUserResponse{}

	err := r.db.
		QueryRow("select id, username, firstname, lastname, email, phone, user_role, profile_picture,created_at,updated_at from users where username = ? or email = ?", dataUser.Username, dataUser.Email).Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.Email, &user.Phone, &user.UserRole, &user.ProfilePicture, &user.CreatedAt, &user.UpdatedAt)

	if user.ID != "" {
		return responses.CreateUserResponse{}, errors.New("username or email already in used")
	}

	_, err = r.db.Exec("insert into users (id,username,password,firstname,lastname,email,phone,user_role,profile_picture) values (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		dataUser.ID, dataUser.Username, dataUser.Password, dataUser.FirstName, dataUser.LastName, dataUser.Email, dataUser.Phone, dataUser.UserRole, dataUser.ProfilePicture)
	if err != nil {
		fmt.Println(err.Error())
		return responses.CreateUserResponse{}, err
	}

	responseUser := responses.CreateUserResponse{
		ID:             dataUser.ID,
		Username:       dataUser.Username,
		FirstName:      dataUser.FirstName,
		LastName:       dataUser.LastName,
		Email:          dataUser.Email,
		Phone:          dataUser.Phone,
		ProfilePicture: dataUser.ProfilePicture,
	}

	return responseUser, nil
}

func (r userRepository) UpdateUser(id string, dataUser entities.UserEntity) (responses.UpdateUserResponse, error) {
	_, err := r.db.Exec("update users set firstname = ?,lastname = ?,email = ?,phone = ?,profile_picture = ? where id = ?", dataUser.FirstName, dataUser.LastName, dataUser.Email, dataUser.Phone, dataUser.ProfilePicture, id)
	if err != nil {
		return responses.UpdateUserResponse{}, err
	}
	responseUser := responses.UpdateUserResponse{
		ID:             id,
		FirstName:      dataUser.FirstName,
		LastName:       dataUser.LastName,
		Email:          dataUser.Email,
		Phone:          dataUser.Phone,
		ProfilePicture: dataUser.ProfilePicture,
	}

	return responseUser, nil
}

func (r userRepository) UpdateUserWithoutImage(id string, dataUser entities.UserEntity) (responses.UpdateUserResponse, error) {
	_, err := r.db.Exec("update users set firstname = ?,lastname = ?,email = ?,phone = ? where id = ?", dataUser.FirstName, dataUser.LastName, dataUser.Email, dataUser.Phone, id)
	if err != nil {
		return responses.UpdateUserResponse{}, err
	}
	responseUser := responses.UpdateUserResponse{
		ID:             id,
		FirstName:      dataUser.FirstName,
		LastName:       dataUser.LastName,
		Email:          dataUser.Email,
		Phone:          dataUser.Phone,
		ProfilePicture: dataUser.ProfilePicture,
	}

	return responseUser, nil
}

func (r userRepository) DeleteUser(id string) (bool, error) {
	_, err := r.db.Exec("update users set is_deleted = 1 where id = ?", id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r userRepository) ChangePassword(requestPassword requests.ChangePasswordRequest, id string) error {
	_, err := r.db.Exec("update users set password = ? where id = ?", requestPassword.NewPassword, id)
	if err != nil {
		return err
	}

	return nil
}

func (r userRepository) FindUserByIDGetPassword(id string) (responses.GetPasswordUserResponse, error) {
	var user = responses.GetPasswordUserResponse{}
	err := r.db.
		QueryRow("select password from users where id = ?", id).Scan(&user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}
