package gerador_campeonato

func NewGeradorCampeonatoDomain(timeA string, timeB string) GeradorCampeonatoDomainInterface {
	return &geradorCampeonatoDomain{timeA: timeA, timeB: timeB}
}
