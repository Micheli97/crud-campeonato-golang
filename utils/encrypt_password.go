package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/Micheli97/crud-campeonato-golang/domain/login"
	"github.com/Micheli97/crud-campeonato-golang/domain/user"
)

// EncryptPassword criptografa a senha do usuário
func EncryptPassword(userDomain user.UserDomainInterface) {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(userDomain.GetPassword()))
	userDomain.SetPassword(hex.EncodeToString(hash.Sum(nil)))
}

// EncryptPassword criptografa a senha do usuário
func EncryptPasswordLogin(loginDomain login.LoginDomainInterface) {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(loginDomain.GetPassword()))
	loginDomain.SetPassword(hex.EncodeToString(hash.Sum(nil)))
}
