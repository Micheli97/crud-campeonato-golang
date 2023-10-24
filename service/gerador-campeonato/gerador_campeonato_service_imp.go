package gerador_campeonato

import (
	rest_err "github.com/Micheli97/crud-campeonato-golang/config/error"
	gerador_campeonato "github.com/Micheli97/crud-campeonato-golang/domain/gerador-campeonato"
	gerador_campeonato2 "github.com/Micheli97/crud-campeonato-golang/repository/gerador-campeonato"
)

func (service *geradorCampeonatoService) GerarCampeonatoHandler() *rest_err.RestErr {
	return service.GerarCampeonatoHandler()

}

func (service *geradorCampeonatoService) ListarTimesCampeonatoHandler() ([]gerador_campeonato.GeradorCampeonatoDomainInterface, *rest_err.RestErr) {
	return service.ListarTimesCampeonatoHandler()
}

func NewGeradorCampeonatoService(repository gerador_campeonato2.GeradorCampeonatoRepositoryInterface) GeradorCampeonatoServiceInterface {
	return &geradorCampeonatoService{repository: repository}
}
