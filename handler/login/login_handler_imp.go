package login

import (
	"github.com/Micheli97/crud-campeonato-golang/config/logger"
	"github.com/Micheli97/crud-campeonato-golang/config/validation"
	login2 "github.com/Micheli97/crud-campeonato-golang/domain/login"
	"github.com/Micheli97/crud-campeonato-golang/handler/model/request"
	"github.com/Micheli97/crud-campeonato-golang/service/login"
	"github.com/Micheli97/crud-campeonato-golang/view"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (loginHandler *loginHandler) LoginHandler(context *gin.Context) {

	var loginRequest request.LoginRequest

	if err := context.BindJSON(&loginRequest); err != nil {
		logger.Error("Erro ao efetuar login", err)
		restErr := validation.ValidateUserError(err)
		context.JSON(restErr.Code, restErr)
		return
	}

	loginDomain := login2.NewLoginDomain(
		loginRequest.Email,
		"",
		loginRequest.Password,
	)

	loginUser, err := loginHandler.loginService.LoginService(loginDomain)

	if err != nil {
		logger.Error("Erro trying to call LoginUser services", err)
		context.JSON(err.Code, err)
		return
	}

	context.JSON(http.StatusOK, view.ConverteToLoginDomainResponse(loginUser))

}

func LoginHandler(loginService login.LoginServiceInterface) LoginHandlerInterface {
	return &loginHandler{
		loginService: loginService,
	}
}
