package user

import (
	"github.com/Micheli97/crud-campeonato-golang/config/logger"
	"github.com/Micheli97/crud-campeonato-golang/config/validation"
	"github.com/Micheli97/crud-campeonato-golang/handler/model/request"
	user2 "github.com/Micheli97/crud-campeonato-golang/model/user"
	"github.com/Micheli97/crud-campeonato-golang/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

var (
	UserDomainInterface user2.UserDomainInterface
)

// CreateUserHandler cadastra o usuário
func (userHandler *userHandlerInterface) CreateUserHandler(context *gin.Context) {

	// Objeto com os atributos necessários para fazer a criação do usuário
	var userRequest request.UserRequest

	if err := context.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Erro trying to validate user info", err,
			zap.String("journey", "createUser"))
		restErr := validation.ValidateUserError(err)
		context.JSON(restErr.Code, restErr)
		return
	}

	domain := user2.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name)

	if err := userHandler.service.CreateUser(domain); err != nil {
		context.JSON(err.Code, err)
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"))

	context.JSON(http.StatusOK, view.ConverteToDomainResponse(domain))
}
