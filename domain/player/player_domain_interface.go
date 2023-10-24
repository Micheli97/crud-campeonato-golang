package player

type PlayerDomainInterface interface {
	GetID() string
	GetName() string
	GetPhoto() string
	GetHeight() float32
	GetWeight() float32
	GetAge() int32
	GetPosition() string
	GetNumber() int32
	GetTeamID() string
	SetID(id string)
}
