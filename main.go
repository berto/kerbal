package main

import (
	"fmt"
	"net/http"
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
		router.Use(corsMiddleware())
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

func corsMiddleware() gin.HandlerFunc {
	clientURL := os.Getenv("CLIENT_URL")
	if clientURL == "" {
		clientURL = "*"
	}
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", clientURL)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func main() {
	port := getPort()
	createRouter().Run(fmt.Sprintf(":%s", port))
}
