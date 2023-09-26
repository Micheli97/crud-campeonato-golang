package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUserHandler retorna os dados do usuário
func GetUserHandler(context *gin.Context) {

	context.JSON(http.StatusOK, "Usuário")

}
