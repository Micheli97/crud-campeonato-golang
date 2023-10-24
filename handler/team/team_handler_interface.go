package team

import "github.com/gin-gonic/gin"

type TeamHandlerInterface interface {
	CreateTeamHandler(context *gin.Context)
	DeleteTeamHandler(context *gin.Context)
	FindTeamByIDHandler(context *gin.Context)
	FindTeams(context *gin.Context)
	FindTotalTeams(context *gin.Context)
	UpdateTeamHandler(context *gin.Context)
}
