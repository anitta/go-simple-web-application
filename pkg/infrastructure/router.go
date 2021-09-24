package infrastructure

import (
	"github.com/anitta/go-simple-web-application/pkg/interfaces/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router(simpleController controllers.SimpleController, allowOrigins []string) *gin.Engine {
	router := gin.New()

	config := cors.DefaultConfig()
	config.AllowOrigins = allowOrigins
	config.AllowMethods = []string{"GET"}
	router.Use(cors.New(config))

	router.GET("/", func(c *gin.Context) { simpleController.Index(c) })

	return router
}
