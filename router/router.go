package router

import (
	"github.com/Micheli97/crud-campeonato-golang/handler/user"
	"github.com/gin-gonic/gin"
	"log"
)

// SetupRouter configura o servidor
func SetupRouter(userHandler user.UserHandlerInterface) {
	router := gin.Default()

	InitializeRouters(router, userHandler)

	if err := router.Run(); err != nil {
		log.Fatal("Ocorreu um erro ao iniciar o servidor.")
	}

}
