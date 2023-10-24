package login

type LoginDomainInterface interface {
	GetID() string
	GetName() string
	GetEmail() string
	SetPassword(password string)
	GetPassword() string
	SetID(id string)
}
