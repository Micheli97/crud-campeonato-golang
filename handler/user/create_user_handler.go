package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateUserHandler cadastra o usuário
func CreateUserHandler(context *gin.Context) {
	context.JSON(http.StatusCreated, "Usuário foi criado com sucesso!")
}
