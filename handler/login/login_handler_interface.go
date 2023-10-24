package login

import (
	"github.com/gin-gonic/gin"
)

type LoginHandlerInterface interface {
	LoginHandler(context *gin.Context)
}
