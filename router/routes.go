package router

import (
	gerador_campeonato2 "github.com/Micheli97/crud-campeonato-golang/handler/gerador-campeonato"
	login2 "github.com/Micheli97/crud-campeonato-golang/handler/login"
	"github.com/Micheli97/crud-campeonato-golang/handler/team"
	user2 "github.com/Micheli97/crud-campeonato-golang/handler/user"
	gerador_campeonato "github.com/Micheli97/crud-campeonato-golang/router/gerador-campeonato"
	"github.com/Micheli97/crud-campeonato-golang/router/login"
	team2 "github.com/Micheli97/crud-campeonato-golang/router/team"
	"github.com/Micheli97/crud-campeonato-golang/router/user"
	"github.com/gin-gonic/gin"
)

// InitializeRouters inicializa as rotas
func InitializeRouters(router *gin.Engine, userHandler user2.UserHandlerInterface, loginHandler login2.LoginHandlerInterface, teamHandler team.TeamHandlerInterface, geradorHandler gerador_campeonato2.GeradorCampeonatoInterface) {

	v1 := router.Group("api/v1")

	user.UserRouters(v1, userHandler)
	login.LoginRouter(v1, loginHandler)
	team2.TeamRouter(v1, teamHandler)
	gerador_campeonato.GeradorCampeonatoRouter(v1, geradorHandler)

}
