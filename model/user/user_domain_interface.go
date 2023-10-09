package user

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetID() string

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

// NewUpdateUserDomain construtor do objeto
func NewUpdateUserDomain(email, name string) UserDomainInterface {
	return &userDomain{
		Name:  name,
		Email: email,
	}
}
