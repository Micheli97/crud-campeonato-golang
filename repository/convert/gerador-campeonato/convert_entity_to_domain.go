package gerador_campeonato

import (
	gerador_campeonato2 "github.com/Micheli97/crud-campeonato-golang/domain/gerador-campeonato"
	gerador_campeonato "github.com/Micheli97/crud-campeonato-golang/entity/gerador-campeonato"
)

func ConvertGeradorCampeonatoToDomain(gerador gerador_campeonato.GeradorCampeonatoEntity) gerador_campeonato2.GeradorCampeonatoDomainInterface {
	return gerador_campeonato2.NewGeradorCampeonatoDomain(gerador.TimeA, gerador.TimeB)
}
