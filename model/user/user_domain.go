package user

import (
	"crypto/md5"
	"encoding/hex"
	rest_err "github.com/Micheli97/crud-campeonato-golang/config/error"
)

// UserDomain não pode ser exportato, podendo ser também escrito de forma privada
// Não pode conter informações da tag (json) porque o Domain não pode ser exportavel
// Objetos(models response, request) junto ao controller/handler são responsáveis por comunicar com os endpoints
type UserDomain struct {
	Email    string
	Password string
	Name     string
}

// NewUserDomain construtor do objeto
func NewUserDomain(email, password, name string) *UserDomain {
	return &UserDomain{email, password, name}
}

// EncryptPassword criptografa a senha do usuário
func (user *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(user.Password))
	user.Password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr
	GetUser(string) (*UserDomain, *rest_err.RestErr)
	UpdateUser(string) *rest_err.RestErr
	DeleteUser(string) *rest_err.RestErr
}
