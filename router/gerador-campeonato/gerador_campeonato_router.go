package gerador_campeonato

import (
	gerador_campeonato "github.com/Micheli97/crud-campeonato-golang/handler/gerador-campeonato"
	"github.com/gin-gonic/gin"
)

func GeradorCampeonatoRouter(context *gin.RouterGroup, handler gerador_campeonato.GeradorCampeonatoInterface) {
	gerador := context.Group("/geradorCampeonato")

	gerador.POST("", handler.GerarCampeonatoHandler)
	gerador.GET("", handler.ListarTimesCampeonatoHandler)

}
