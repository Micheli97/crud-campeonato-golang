package user

// Não pode conter informações da tag (json) porque o Domain não pode ser exportavel
// Objetos(models response, request) junto ao controller/handler são responsáveis por comunicar com os endpoints
type userDomain struct {
	ID       string
	Email    string
	Password string
	Name     string
}

func (user *userDomain) SetID(id string) {
	user.ID = id
}

func (user *userDomain) GetEmail() string {
	return user.Email
}

func (user *userDomain) GetPassword() string {
	return user.Password
}

func (user *userDomain) GetName() string {
	return user.Name
}

func (user *userDomain) GetID() string {
	return user.ID
}

func (user *userDomain) SetPassword(password string) {
	user.Password = password
}
