package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteUserHandler retorna os dados do usuário
func DeleteUserHandler(context *gin.Context) {

	context.JSON(http.StatusOK, "Usuário")

}
