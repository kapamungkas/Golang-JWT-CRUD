package middlewares

import (
	"betest/internal/responses"
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var APPLICATION_NAME = os.Getenv("APPLICATION_NAME")
var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte(os.Getenv("JWT_SIGNATURE_KEY"))
var REFRESH_JWT_SIGNATURE_KEY = []byte(os.Getenv("REFRESH_JWT_SIGNATURE_KEY"))
var REFRESH_LOGIN_EXPIRATION_DURATION = time.Duration(1000) * time.Hour

type UserJWT struct {
	jwt.StandardClaims
	Username string `json:"Username"`
	UserRole string `json:"UserRole"`
	UserID   string `json:"UserID"`
}

func UserMiddleware(role []string) gin.HandlerFunc {

	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")

		if len(authHeader) < 6 {
			respondWithError(c, 401, "API token required")
			return
		}

		tokenString := authHeader[len(BEARER_SCHEMA):]

		claims, err := VerifyAccessToken(tokenString)

		if err != nil {
			respondWithError(c, 401, err.Error())
			return
		}
		auth := false
		for _, v := range role {
			if claims["UserRole"] == v {
				auth = true
			}
		}

		if auth == false {
			respondWithError(c, 401, "You do not have access")
			return
		}

		c.Set("user_id", claims["UserID"])
		c.Set("username", claims["Username"])
		c.Set("user_role", claims["UserRole"])
		c.Next()
	}
}

func SignAccessToken(user responses.LoginResponse) (string, error) {
	claims := UserJWT{
		StandardClaims: jwt.StandardClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
		},
		Username: user.Username,
		UserRole: user.UserRole,
		UserID:   user.ID,
	}

	token := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		claims,
	)

	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyAccessToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			// respondWithError(c, 401, "Signing method invalid")
			return JWT_SIGNATURE_KEY, errors.New("signing method invalid")
		} else if method != JWT_SIGNING_METHOD {
			return JWT_SIGNATURE_KEY, errors.New("signing method invalid")
		}

		return JWT_SIGNATURE_KEY, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}

func SignRefreshToken(user responses.LoginResponse) (string, error) {
	claims := UserJWT{
		StandardClaims: jwt.StandardClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: time.Now().Add(REFRESH_LOGIN_EXPIRATION_DURATION).Unix(),
		},
		Username: user.Username,
		UserRole: user.UserRole,
		UserID:   user.ID,
	}

	token := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		claims,
	)

	signedToken, err := token.SignedString(REFRESH_JWT_SIGNATURE_KEY)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyRefreshToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			// respondWithError(c, 401, "Signing method invalid")
			return REFRESH_JWT_SIGNATURE_KEY, errors.New("signing method invalid")
		} else if method != JWT_SIGNING_METHOD {
			return REFRESH_JWT_SIGNATURE_KEY, errors.New("signing method invalid")
		}

		return REFRESH_JWT_SIGNATURE_KEY, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{
		"status":  400,
		"message": message,
		"data":    "",
	})
}
