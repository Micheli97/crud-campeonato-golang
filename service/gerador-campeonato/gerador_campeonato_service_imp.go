package gerador_campeonato

import (
	rest_err "github.com/Micheli97/crud-campeonato-golang/config/error"
	gerador_campeonato "github.com/Micheli97/crud-campeonato-golang/domain/gerador-campeonato"
)

func (service *geradorCampeonatoService) GerarCampeonatoHandler() *rest_err.RestErr {
	return nil

}

func (service *geradorCampeonatoService) ListarTimesCampeonatoHandler() ([]gerador_campeonato.GeradorCampeonatoDomainInterface, *rest_err.RestErr) {
	return nil, nil
}

func NewGeradorCampeonatoService() GeradorCampeonatoServiceInterface {
	return &geradorCampeonatoService{}
}
