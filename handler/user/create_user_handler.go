package user

import (
	"github.com/Micheli97/crud-campeonato-golang/config/validation"
	"github.com/Micheli97/crud-campeonato-golang/handler/model/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateUserHandler cadastra o usuário
func CreateUserHandler(context *gin.Context) {

	// Objeto com os atributos necessários para fazer a criação do usuário
	var userRequest request.UserRequest

	if err := context.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)
		context.JSON(restErr.Code, restErr)
		return
	}

	context.JSON(http.StatusCreated, "Usuário foi criado com sucesso!")
}
