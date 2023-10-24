package gerador_campeonato

type geradorCampeonatoDomain struct {
	timeA string
	timeB string
}

func (g *geradorCampeonatoDomain) GetTimes() []string {
	listaJogos := []string{g.timeA, g.timeB}

	return listaJogos

}

func (g *geradorCampeonatoDomain) GetTimeA() string {
	return g.timeA

}

func (g *geradorCampeonatoDomain) GetTimeB() string {
	return g.timeB

}
