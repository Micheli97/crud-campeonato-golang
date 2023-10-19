package user

import (
	"github.com/Micheli97/crud-campeonato-golang/config/logger"
	"github.com/Micheli97/crud-campeonato-golang/config/validation"
	user2 "github.com/Micheli97/crud-campeonato-golang/domain/user"
	"github.com/Micheli97/crud-campeonato-golang/handler/model/request"
	"github.com/Micheli97/crud-campeonato-golang/service/user"
	"github.com/Micheli97/crud-campeonato-golang/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func NewUserHandlerInterface(serviceInterface user.UserDomainService) UserHandlerInterface {
	return &userHandler{
		service: serviceInterface,
	}
}

var (
	UserDomainInterface user2.UserDomainInterface
)

// CreateUserHandler cadastra o usuário
func (userHandler *userHandler) CreateUserHandler(context *gin.Context) {

	// Objeto com os atributos necessários para fazer a criação do usuário
	var userRequest request.UserRequest

	logger.Info(userRequest.Name)

	if err := context.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Erro trying to validate user info", err,
			zap.String("journey", "createUser"))
		restErr := validation.ValidateUserError(err)
		context.JSON(restErr.Code, restErr)
		return
	}

	registeredUser, err := userHandler.service.FindUserByEmailServices(userRequest.Email)

	if registeredUser != nil {
		context.JSON(http.StatusBadRequest, "msg: Usuário já cadastrado no sistema!")
		return
	}

	domain := user2.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name)

	domainResult, err := userHandler.service.CreateUser(domain)

	if err != nil {
		context.JSON(err.Code, err)
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"))

	context.JSON(http.StatusOK, view.ConverteToDomainResponse(domainResult))
}

// DeleteUserHandler retorna os dados do usuário
func (userHandler *userHandler) DeleteUserHandler(context *gin.Context) {
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

// FindUserByIDHandler retorna os dados do usuário
func (userHandler *userHandler) FindUserByIDHandler(context *gin.Context) {
	logger.Info("Init findUserByID handler", zap.String("journey", "findUserID"))

	userID := context.Param("userId")

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
func (userHandler *userHandler) FindUserByEmailHandler(context *gin.Context) {

	logger.Info("Init findUserByEmail handler", zap.String("journey", "findUserByEmail"))

	userEmail := context.Param("userEmail")

	userDomain, err := userHandler.service.FindUserByEmailServices(userEmail)

	if err != nil {
		context.JSON(err.Code, err)
		return
	}

	context.JSON(http.StatusOK, view.ConverteToDomainResponse(userDomain))

}

// UpdateUserHandler retorna os dados do usuário
func (userHandler *userHandler) UpdateUserHandler(context *gin.Context) {

	var userRequest request.UserUpdateRequest

	if err := context.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Erro trying to validate user info", err,
			zap.String("journey", "createUser"))
		restErr := validation.ValidateUserError(err)
		context.JSON(restErr.Code, restErr)
		return
	}

	userID := context.Param("userId")

	userUpdateDomain := user2.NewUpdateUserDomain(
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
