package gerador_campeonato

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (handler *geradorCampeonatoHandler) GerarCampeonatoHandler(context *gin.Context) {
	context.JSON(http.StatusOK, "Campeonato gerado com sucesso!")

}

func (handler *geradorCampeonatoHandler) ListarTimesCampeonatoHandler(context *gin.Context) {
	context.JSON(http.StatusOK, "Time A x Time B")
}

func NewGeradorHandler() GeradorCampeonatoInterface {
	return &geradorCampeonatoHandler{}
}
