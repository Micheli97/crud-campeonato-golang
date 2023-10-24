package gerador_campeonato

import (
	gerador_campeonato "github.com/Micheli97/crud-campeonato-golang/service/gerador-campeonato"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (handler *geradorCampeonatoHandler) GerarCampeonatoHandler(context *gin.Context) {

	if err := handler.service.GerarCampeonatoHandler(); err != nil {
		context.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	context.JSON(http.StatusOK, "Campeonato gerado com sucesso!")

}

func (handler *geradorCampeonatoHandler) ListarTimesCampeonatoHandler(context *gin.Context) {

	times, err := handler.service.ListarTimesCampeonatoHandler()

	if err != nil {
		context.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	context.JSON(http.StatusOK, times)
}

func NewGeradorHandler(service gerador_campeonato.GeradorCampeonatoServiceInterface) GeradorCampeonatoInterface {
	return &geradorCampeonatoHandler{service: service}
}
