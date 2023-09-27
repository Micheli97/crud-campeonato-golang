package user

import (
	"crypto/md5"
	"encoding/hex"
)

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string

	EncryptPassword()

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

// EncryptPassword criptografa a senha do usuário
func (user *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(user.Password))
	user.Password = hex.EncodeToString(hash.Sum(nil))
}
