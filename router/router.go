package router

import (
	"github.com/Micheli97/crud-campeonato-golang/handler/user"
	service "github.com/Micheli97/crud-campeonato-golang/model/service/user"
	"github.com/gin-gonic/gin"
	"log"
)

// SetupRouter configura o servidor
func SetupRouter() {
	router := gin.Default()

	// Init dependencies
	service := service.NewUserDomainService()
	userHandler := user.NewUserHandlerInterface(service)

	InitializeRouters(router, userHandler)

	if err := router.Run(); err != nil {
		log.Fatal("Ocorreu um erro ao iniciar o servidor.")
	}

}
