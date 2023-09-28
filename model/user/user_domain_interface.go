package user

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string

	SetPassword(password string)

	SetID(id string)
}

// NewUserDomain construtor do objeto
func NewUserDomain(email, password, name string) UserDomainInterface {
	return &userDomain{
		Name:     name,
		Password: password,
		Email:    email,
	}
}
