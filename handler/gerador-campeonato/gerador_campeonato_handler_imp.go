package gerador_campeonato

import (
	gerador_campeonato "github.com/Micheli97/crud-campeonato-golang/service/gerador-campeonato"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (handler *geradorCampeonatoHandler) GerarCampeonatoHandler(context *gin.Context) {
	context.JSON(http.StatusOK, "Campeonato gerado com sucesso!")

}

func (handler *geradorCampeonatoHandler) ListarTimesCampeonatoHandler(context *gin.Context) {
	context.JSON(http.StatusOK, "Time A x Time B")
}

func NewGeradorHandler(service gerador_campeonato.GeradorCampeonatoServiceInterface) GeradorCampeonatoInterface {
	return &geradorCampeonatoHandler{service: service}
}
