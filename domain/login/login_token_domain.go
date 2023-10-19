package login

import (
	rest_err "github.com/Micheli97/crud-campeonato-golang/config/error"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
	"time"
)

var (
	JWT_SECRET_KEY = "JWT_SECRETY_KEY"
)

func (login *loginDomain) GenerateToken() (string, *rest_err.RestErr) {
	secret := os.Getenv(JWT_SECRET_KEY)

	claims := jwt.MapClaims{
		"ID":    login.ID,
		"email": login.Email,
		"name":  login.Name,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", rest_err.NewInternalServerError(err.Error())
	}

	return tokenString, nil
}

func MiddlewareVerifyToken(context *gin.Context) {
	secret := os.Getenv(JWT_SECRET_KEY)
	tokenValue := RemoveBearerPrefix(context.Request.Header.Get("token"))

	token, err := jwt.Parse(RemoveBearerPrefix(tokenValue), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}
		return nil, rest_err.NewBadRequestError("Inválid token")
	})

	if err != nil {
		errRest := rest_err.NewUnauthorizedRequestError("invalid token")
		context.JSON(errRest.Code, errRest)
	}

	_, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		errRest := rest_err.NewUnauthorizedRequestError("invalid token")
		context.JSON(errRest.Code, errRest)
	}
}

func RemoveBearerPrefix(token string) string {
	if strings.HasPrefix(token, "Bearer") {
		token = strings.TrimPrefix("Bearer", token)
	}

	return token
}
