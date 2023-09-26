package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// UpdateUserHandler retorna os dados do usuário
func UpdateUserHandler(context *gin.Context) {
	context.JSON(http.StatusOK, "Usuário")

}
