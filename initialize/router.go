package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/jocker-l/groq_api_pool/middlewares"
	"github.com/jocker-l/groq_api_pool/router"
)

func InitRouter() *gin.Engine {
	Router := gin.Default()

	Router.Use(middlewares.Cors)

	Router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})

	Router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	v1Group := Router.Group("/v1/")
	router.InitChat(v1Group)

	return Router
}
