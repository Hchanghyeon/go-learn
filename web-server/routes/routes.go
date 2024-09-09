package routes

import (
	"web-server/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine){
	memberRoutes := r.Group("/members")
	{
		memberRoutes.GET("/", controllers.GetMembers)
		memberRoutes.POST("/",  controllers.CreateMember)
	}
}