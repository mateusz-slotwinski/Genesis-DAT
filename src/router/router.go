package router

import (
	cors "github.com/gin-contrib/cors"
	gin "github.com/gin-gonic/gin"

	Controllers "GenesisDAT/src/controllers"
	// Database "GenesisDAT/src/database"
	Middlewares "GenesisDAT/src/middlewares"
)

func CreateServer() *gin.Engine {
	R := gin.Default()

	// Database.ConnectDB()

	R.Use(cors.Default())
	R.SetTrustedProxies([]string{"192.168.1.2"})

	R.Use(Middlewares.UseAuthorization)

	SciencesRoutes := R.Group("/sciences")
	{
		var Controller Controllers.SciencesController

		R_Groups := SciencesRoutes.Group("/groups")
		{
			R_Groups.GET("/all", Controller.GetGroups)
		}

	}

	return R
}
