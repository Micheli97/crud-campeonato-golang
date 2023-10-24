package router

import (
	gerador_campeonato2 "github.com/Micheli97/crud-campeonato-golang/handler/gerador-campeonato"
	login2 "github.com/Micheli97/crud-campeonato-golang/handler/login"
	"github.com/Micheli97/crud-campeonato-golang/handler/team"
	"github.com/Micheli97/crud-campeonato-golang/handler/user"
	"github.com/gin-gonic/gin"
	"log"
)

// SetupRouter configura o servidor
func SetupRouter(userHandler user.UserHandlerInterface, loginHandler login2.LoginHandlerInterface, teamHandler team.TeamHandlerInterface, geradorHandler gerador_campeonato2.GeradorCampeonatoInterface) {
	router := gin.Default()

	InitializeRouters(router, userHandler, loginHandler, teamHandler, geradorHandler)

	if err := router.Run(); err != nil {
		log.Fatal("Ocorreu um erro ao iniciar o servidor.")
	}

}
