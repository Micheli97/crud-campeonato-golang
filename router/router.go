package router

import (
	"github.com/gin-gonic/gin"
	"log"
)

func SetupRouter() {
	router := gin.Default()

	if err := router.Run(); err != nil {
		log.Fatal("Ocorreu um erro ao iniciar o servidor.")
	}

}
