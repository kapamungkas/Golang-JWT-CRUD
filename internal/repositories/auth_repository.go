package repositories

import (
	"betest/internal/requests"
	"betest/internal/responses"
	"database/sql"
	"fmt"
)

type AuthRepository interface {
	Login(dataAuth requests.LoginRequest) (responses.LoginResponse, error)
	UpdateRefreshToken(id string, refreshToken string) error
	CheckRefreshTokenOnDatabase(id string, refreshToken string) (responses.LoginResponse, error)
	CheckEmailUserOnDatabase(email string) (responses.LoginResponse, error)
	CheckTokenUser(token string) (responses.LoginResponse, error)
	UpdateResetPassword(id string, token string) error
	UpdatePasswordUser(id string, new_password string) error
}

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *authRepository {
	return &authRepository{db}
}

func (r authRepository) Login(dataAuth requests.LoginRequest) (responses.LoginResponse, error) {
	var result = responses.LoginResponse{}
	err := r.db.
		QueryRow("select id,username,password,user_role from users where username = ?", dataAuth.Username).
		Scan(&result.ID, &result.Username, &result.Password, &result.UserRole)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (r authRepository) UpdateRefreshToken(id string, refreshToken string) error {
	_, err := r.db.Exec("update users set refresh_token = ? where id = ?", refreshToken, id)
	if err != nil {
		return err
	}

	return nil
}

func (r authRepository) CheckRefreshTokenOnDatabase(id string, refreshToken string) (responses.LoginResponse, error) {
	var result = responses.LoginResponse{}

	err := r.db.QueryRow("select id,username,password,user_role from users where id = ? AND refresh_token = ?", id, refreshToken).
		Scan(&result.ID, &result.Username, &result.Password, &result.UserRole)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (r authRepository) CheckEmailUserOnDatabase(email string) (responses.LoginResponse, error) {
	var result = responses.LoginResponse{}

	err := r.db.QueryRow("select id from users where email = ? ", email).
		Scan(&result.ID)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (r authRepository) CheckTokenUser(token string) (responses.LoginResponse, error) {
	var result = responses.LoginResponse{}
	fmt.Println(token)

	err := r.db.QueryRow("select id,email from users where reset_password = ? ", token).
		Scan(&result.ID, &result.Email)

	if err != nil {
		return responses.LoginResponse{}, err
	}

	return result, nil
}

func (r authRepository) UpdateResetPassword(id string, token string) error {
	_, err := r.db.Exec("update users set reset_password = ? where id = ?", token, id)
	if err != nil {
		return err
	}

	return nil
}

func (r authRepository) UpdatePasswordUser(id string, new_password string) error {
	_, err := r.db.Exec("update users set password = ?,reset_password = '' where id = ?", new_password, id)
	if err != nil {
		return err
	}

	return nil
}
