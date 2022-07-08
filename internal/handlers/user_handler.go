package handlers

import (
	"betest/internal/entities"
	"betest/internal/helpers"
	requests "betest/internal/requests"
	services "betest/internal/services"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userHandler struct {
	services.UserService
}

func NewUserHandler(service services.UserService) *userHandler {
	return &userHandler{service}
}

func (h userHandler) GetAllUser(c *gin.Context) {
	option := map[string]string{}
	if c.Query("order_by") != "" {
		option["order_by"] = c.Query("order_by")
	}

	if c.Query("limit") != "" {
		option["limit"] = c.Query("limit")
	}

	user, err := h.UserService.GetAllUser(option)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	if user == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "Data is empty",
			"data":    user,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Success get all data",
		"data":    user,
	})
}

func (h userHandler) FindUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := h.UserService.FindUserByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Success find data",
		"data":    user,
	})
}

func (h userHandler) CreateUser(c *gin.Context) {

	var User requests.CreateUserRequest

	err := c.ShouldBind(&User)

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

	id := uuid.New()
	password, err := bcrypt.GenerateFromPassword([]byte(User.Password), 14)
	User.Password = string(password)

	file_type := []string{"image/png", "image/jpeg"}

	upload, err_up := helpers.SingleFileUpload(c, "profile_pic", os.Getenv("STORAGE_PATH"), file_type, 1024, true)

	if err_up != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  400,
			"message": err_up.Error(),
			"data":    "",
		})
		return
	}

	dataUser := entities.UserEntity{
		ID:             id.String(),
		Username:       User.Username,
		Password:       User.Password,
		FirstName:      User.FirstName,
		LastName:       User.LastName,
		Email:          User.Email,
		Phone:          User.Phone,
		UserRole:       User.UserRole,
		ProfilePicture: upload,
	}

	user, err := h.UserService.CreateUser(dataUser)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  400,
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Created success",
		"data":    user,
	})
}

func (h userHandler) UpdateUser(c *gin.Context) {

	id := c.Param("id")

	var User requests.UpdateUserRequest

	err := c.ShouldBind(&User)

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

	file_type := []string{"image/png", "image/jpeg"}

	upload, err := helpers.SingleFileUpload(c, "profile_pic", os.Getenv("STORAGE_PATH"), file_type, 1024, false)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  400,
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	dataUser := entities.UserEntity{
		FirstName:      User.FirstName,
		LastName:       User.LastName,
		Email:          User.Email,
		Phone:          User.Phone,
		ProfilePicture: upload,
	}

	user, err := h.UserService.UpdateUser(id, dataUser)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  400,
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Updated success",
		"data":    user,
	})
}

func (h userHandler) DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	_, err := h.UserService.DeleteUser(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Success deleted data",
		// "data":    user,
	})
}

func (h userHandler) ChangePassword(c *gin.Context) {

	id := c.Params.ByName("id")

	var ChangePassword requests.ChangePasswordRequest

	err := c.ShouldBind(&ChangePassword)

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

	if ChangePassword.NewPassword != ChangePassword.ConfirmNewPassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "New password and confirm password not same",
			"data":    "",
		})
		return
	}

	new_password, err := bcrypt.GenerateFromPassword([]byte(ChangePassword.NewPassword), 14)
	ChangePassword.NewPassword = string(new_password)

	ChangePassword = requests.ChangePasswordRequest{
		CurrentPassword: ChangePassword.CurrentPassword,
		NewPassword:     string(new_password),
	}
	err = h.UserService.ChangePassword(ChangePassword, id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  400,
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Change password success",
	})
}

func (h userHandler) MyProfile(c *gin.Context) {
	id := c.GetString("user_id")
	user, err := h.UserService.FindUserByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Success find data",
		"data":    user,
	})
}
