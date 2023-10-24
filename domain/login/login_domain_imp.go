package login

func NewLoginDomain(email, name, password string) LoginDomainInterface {
	return &loginDomain{
		Email:    email,
		Name:     name,
		Password: password,
	}

}
