package routes

import (
	"github.com/anvarmx/tranzify-api/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")

	{
		api.POST("/morseToText", handlers.MorseToTextHandler)
		api.POST("/textToMorse", handlers.TextToMorseHandler)
	}

	r.GET("/", handlers.RootHandler)

	r.NoRoute(handlers.NoRouteHandler)
}
