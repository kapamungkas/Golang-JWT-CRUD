package routes

import (
	"betest/internal/handlers"
	"betest/internal/middlewares"
	"betest/internal/repositories"
	"betest/internal/services"
	"database/sql"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRoute(db *sql.DB) {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Accept", "Authorization"}
	router.Use(cors.New(config))
	api_group := router.Group("/api")
	{
		user_group := api_group.Group("/user")
		{
			user_repo := repositories.NewUserRepository(db)
			user_service := services.NewUserService(user_repo)
			user_handler := handlers.NewUserHandler(user_service)

			user_group.GET("/", middlewares.UserMiddleware([]string{"admin"}), user_handler.GetAllUser)
			user_group.GET("/:id", middlewares.UserMiddleware([]string{"admin"}), user_handler.FindUserByID)
			user_group.POST("/", middlewares.UserMiddleware([]string{"admin"}), user_handler.CreateUser)
			// user_group.POST("/", user_handler.CreateUser)
			user_group.PATCH("/:id", middlewares.UserMiddleware([]string{"admin"}), user_handler.UpdateUser)
			user_group.DELETE("/:id", middlewares.UserMiddleware([]string{"admin"}), user_handler.DeleteUser)
			user_group.PATCH("/change-password/:id", user_handler.ChangePassword)

			user_group.GET("/my-profile", middlewares.UserMiddleware([]string{"admin", "user"}), user_handler.MyProfile)
		}

		auth_group := api_group.Group("/auth")
		{
			auth_repo := repositories.NewAuthRepository(db)
			auth_service := services.NewAuthService(auth_repo)
			auth_handler := handlers.NewAuthHandler(auth_service)

			auth_group.POST("/login", auth_handler.Login)
			auth_group.POST("/refresh-token", auth_handler.RefreshToken)
			auth_group.POST("/reset-password", auth_handler.ResetPassword)
			auth_group.GET("/generate-password", auth_handler.GeneratePassword)

		}
	}

	router.Static("/storages", os.Getenv("STORAGE_PATH"))

	router.Run()
}
