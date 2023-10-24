package gerador_campeonato

type geradorCampeonatoDomain struct {
	timeA string
	timeB string
}

func (g *geradorCampeonatoDomain) GetTime() []string {
	listaJogos := []string{g.timeA, g.timeB}

	return listaJogos

}
