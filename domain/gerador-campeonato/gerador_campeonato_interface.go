package gerador_campeonato

type GeradorCampeonatoDomainInterface interface {
	GetTimes() []string
	GetTimeA() string
	GetTimeB() string
}
