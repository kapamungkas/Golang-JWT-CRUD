package handlers

import (
	"betest/internal/middlewares"
	"betest/internal/requests"
	services "betest/internal/services"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type authHandler struct {
	s services.AuthService
}

func NewAuthHandler(service services.AuthService) *authHandler {
	return &authHandler{service}
}

func (h authHandler) Login(c *gin.Context) {
	var Login requests.LoginRequest

	err := c.ShouldBind(&Login)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on filled %s, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)

		}
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": errorMessages,
			"data":    "",
		})
		return
	}

	dataLogin := requests.LoginRequest{
		Username: Login.Username,
		Password: Login.Password,
	}

	user, err := h.s.Login(dataLogin)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  400,
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	tokenString, err := middlewares.SignAccessToken(user)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  400,
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	refreshTokenString, err := middlewares.SignRefreshToken(user)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  400,
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	updateRefreshToken := h.s.UpdateRefreshToken(user.ID, refreshTokenString)

	if updateRefreshToken != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  400,
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":        200,
		"message":       "Authectication success",
		"token":         tokenString,
		"refresh_token": refreshTokenString,
	})
}

func (h authHandler) RefreshToken(c *gin.Context) {
	var RefreshToken requests.RefreshToken

	err := c.ShouldBind(&RefreshToken)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on filled %s, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)

		}
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": errorMessages,
			"data":    "",
		})
		return
	}

	claims, err := middlewares.VerifyRefreshToken(RefreshToken.Token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": err,
			"data":    "verify_refresh_token",
		})
		return
	}
	user_id := fmt.Sprintf("%v", claims["UserID"])

	user, err := h.s.CheckRefreshTokenOnDatabase(user_id, RefreshToken.Token)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": err.Error(),
			"data":    "check_on_db",
		})
		return
	}

	tokenString, err := middlewares.SignAccessToken(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": err.Error(),
			"data":    "create_access_token",
		})
		return
	}

	refreshTokenString, err := middlewares.SignRefreshToken(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": err.Error(),
			"data":    "create_refresh_token",
		})
		return
	}

	updateRefreshToken := h.s.UpdateRefreshToken(user.ID, refreshTokenString)

	if updateRefreshToken != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": err.Error(),
			"data":    "update_refresh_token_on_db",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":        200,
		"message":       "Refresh Token success",
		"token":         tokenString,
		"refresh_token": refreshTokenString,
	})

}

func (h authHandler) ResetPassword(c *gin.Context) {
	var ResetPassword requests.ResetPasswordRequest

	err := c.ShouldBind(&ResetPassword)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on filled %s, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)

		}
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": errorMessages,
			"data":    "",
		})
		return
	}

	random := rand.Intn(999999999-111111111) + 111111111
	err = h.s.ResetPassword(ResetPassword.Email, strconv.Itoa(random))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": err.Error(),
			"data":    "reset_password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Reset password success please check your email",
	})
}

func (h authHandler) GeneratePassword(c *gin.Context) {

	get_token, err := c.GetQuery("token")

	if err != true {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Token not set",
			"data":    "generate_password",
		})
		return
	}

	password := rand.Intn(999999999-111111111) + 111111111
	hash_password, error := bcrypt.GenerateFromPassword([]byte(strconv.Itoa(password)), 14)

	error = h.s.GeneratePassword(get_token, strconv.Itoa(password), string(hash_password))

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": error.Error(),
			"data":    "generate_password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Generate password success please check your email",
	})
}
