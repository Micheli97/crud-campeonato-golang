package gerador_campeonato

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GeradorCampeonatoRouter(context *gin.RouterGroup) {
	gerador := context.Group("/geradorCampeonato")

	gerador.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "Jogos do campeonato")

	})

}
