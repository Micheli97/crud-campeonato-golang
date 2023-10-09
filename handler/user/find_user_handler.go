package user

import (
	"github.com/Micheli97/crud-campeonato-golang/config/logger"
	"github.com/Micheli97/crud-campeonato-golang/view"
	"github.com/gin-gonic/gin"

	"go.uber.org/zap"
	"net/http"
)

// FindUserByIDHandler retorna os dados do usuário
func (userHandler *userHandlerInterface) FindUserByIDHandler(context *gin.Context) {
	logger.Info("Init findUserByID handler", zap.String("journey", "findUserID"))

	userID := context.Param("userId")

	//if _, err := uuid.Parse(userID); err != nil {
	//	errorMessage := rest_err.NewBadRequestError("UserID is not a valid id")
	//	context.JSON(errorMessage.Code, errorMessage)
	//	return
	//}

	userDomain, err := userHandler.service.FindUserByIDServices(userID)

	if err != nil {
		logger.Error("Erro trying to call findUserByID services", err)
		context.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserByID handler excuted successfully", zap.String("journey", "findUserByID"))

	context.JSON(http.StatusOK, view.ConverteToDomainResponse(userDomain))

}

// FindUserByEmailHandler retorna os dados do usuário
func (userHandler *userHandlerInterface) FindUserByEmailHandler(context *gin.Context) {

	logger.Info("Init findUserByEmail handler", zap.String("journey", "findUserByEmail"))

	userEmail := context.Param("userEmail")

	//if _, err := mail.ParseAddress(userEmail); err != nil {
	//	errorMessage := rest_err.NewBadRequestError("Email is not a valid email")
	//	context.JSON(errorMessage.Code, errorMessage)
	//}

	userDomain, err := userHandler.service.FindUserByEmailServices(userEmail)

	if err != nil {
		context.JSON(err.Code, err)
		return
	}

	context.JSON(http.StatusOK, view.ConverteToDomainResponse(userDomain))

}
