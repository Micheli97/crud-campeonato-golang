package router

import (
	"github.com/gin-gonic/gin"
	"log"
)

// SetupRouter configura o servidor
func SetupRouter() {
	router := gin.Default()

	InitializeRouters(router)

	if err := router.Run(); err != nil {
		log.Fatal("Ocorreu um erro ao iniciar o servidor.")
	}

}
