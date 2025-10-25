package main

import (
	"github.com/anvarmx/tranzify-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":3030")
}
