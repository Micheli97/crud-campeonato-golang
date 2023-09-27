package view

import (
	"github.com/Micheli97/crud-campeonato-golang/handler/model/response"
	"github.com/Micheli97/crud-campeonato-golang/model/user"
)

// ConverteToDomainResponse covnerte o objeto domain para response
func ConverteToDomainResponse(userDomain user.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
	}
}
