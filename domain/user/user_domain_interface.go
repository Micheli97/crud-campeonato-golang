package user

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetID() string

	SetPassword(password string)

	SetID(id string)
}
