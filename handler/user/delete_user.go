package user

import (
	"github.com/Micheli97/crud-campeonato-golang/config/logger"
	"github.com/gin-gonic/gin"

	"net/http"
)

// DeleteUserHandler retorna os dados do usuário
func (userHandler *userHandlerInterface) DeleteUserHandler(context *gin.Context) {
	userID := context.Param("userId")

	err := userHandler.service.DeleteUser(userID)

	user, err := userHandler.service.FindUserByIDServices(userID)

	if user == nil {
		context.JSON(http.StatusBadRequest, "msg: Usuário não encontrado")
		return
	}

	if err != nil {
		if err != nil {
			logger.Error("Erro trying to call DeleteUserHandler services", err)
			context.JSON(err.Code, err)
			return
		}
	}

	context.JSON(http.StatusOK, "Usuário excluído com sucesso!")

}
