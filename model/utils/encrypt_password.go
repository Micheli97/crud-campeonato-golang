package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/Micheli97/crud-campeonato-golang/model/user"
)

// EncryptPassword criptografa a senha do usuário
func EncryptPassword(userDomain user.UserDomainInterface) {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(userDomain.GetPassword()))
	userDomain.SetPassword(hex.EncodeToString(hash.Sum(nil)))
}
