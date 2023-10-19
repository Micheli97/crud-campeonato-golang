package login

import (
	rest_err "github.com/Micheli97/crud-campeonato-golang/config/error"
	"github.com/dgrijalva/jwt-go"
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

func VerifyToken(tokenValue string) (LoginDomainInterface, *rest_err.RestErr) {
	secret := os.Getenv(JWT_SECRET_KEY)

	token, err := jwt.Parse(RemoveBearerPrefix(tokenValue), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}
		return nil, rest_err.NewBadRequestError("Inválid token")
	})

	if err != nil {
		return nil, rest_err.NewUnauthorizedRequestError("Invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, rest_err.NewUnauthorizedRequestError("Invalid token")
	}

	return &loginDomain{
		ID:       claims["id"].(string),
		Name:     claims["name"].(string),
		Email:    claims["email"].(string),
		Password: "",
	}, nil
}

func RemoveBearerPrefix(token string) string {
	if strings.HasPrefix(token, "Bearer") {
		token = strings.TrimPrefix("Bearer", token)
	}

	return token
}
