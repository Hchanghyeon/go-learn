package main

import (
	"web-server/routes"

	"github.com/gin-gonic/gin"
)

func main(){
	var engine *gin.Engine = gin.Default()
	routes.SetupRoutes(engine)

	engine.Run()
}