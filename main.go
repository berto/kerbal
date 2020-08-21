package main

import (
	"fmt"
	"os"

	"github.com/berto/kerbal/views"
	"github.com/gin-gonic/gin"
)

const portENV = "PORT"
const defaultPort = "3000"

func createRouter() *gin.Engine {
	router := gin.New()
	{
		router.Use(gin.Recovery())
		router.Use(gin.Logger())
		router.LoadHTMLGlob("./client/*.html")
		router.Static("/js", "./client/js")
		router.Static("/css", "./client/css")
	}
	{
		router.GET("/", views.Home)
		router.GET("/download", views.Download)
		views.KerbalRoutes(router.Group("/kerbal"))
		views.APIRoutes(router.Group("/api"))
	}
	return router
}

func getPort() string {
	port := os.Getenv(portENV)
	if port == "" {
		return defaultPort
	}
	return port
}

func main() {
	port := getPort()
	createRouter().Run(fmt.Sprintf(":%s", port))
}
