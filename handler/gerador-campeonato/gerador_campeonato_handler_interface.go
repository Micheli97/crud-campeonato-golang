package gerador_campeonato

import "github.com/gin-gonic/gin"

type GeradorCampeonatoInterface interface {
	GerarCampeonatoHandler(context *gin.Context)
	ListarTimesCampeonatoHandler(context *gin.Context)
}
