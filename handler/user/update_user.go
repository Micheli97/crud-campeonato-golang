package user

import (
	"github.com/Micheli97/crud-campeonato-golang/config/logger"
	"github.com/Micheli97/crud-campeonato-golang/config/validation"
	"github.com/Micheli97/crud-campeonato-golang/handler/model/request"
	"github.com/Micheli97/crud-campeonato-golang/model/user"
	"github.com/gin-gonic/gin"

	"go.uber.org/zap"
	"net/http"
)

// UpdateUserHandler retorna os dados do usuário
func (userHandler *userHandlerInterface) UpdateUserHandler(context *gin.Context) {

	var userRequest request.UserUpdateRequest

	if err := context.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Erro trying to validate user info", err,
			zap.String("journey", "createUser"))
		restErr := validation.ValidateUserError(err)
		context.JSON(restErr.Code, restErr)
		return
	}

	userID := context.Param("userId")

	userUpdateDomain := user.NewUpdateUserDomain(
		userRequest.Email,
		userRequest.Name,
	)

	err := userHandler.service.UpdateUser(userID, userUpdateDomain)
	if err != nil {
		logger.Error("Erro trying to call UpdateUserHandler services", err)
		context.JSON(err.Code, err)
		return
	}

	context.JSON(http.StatusOK, "msg: Os campos Nome e Email foram atualizados com sucesso!")

}
