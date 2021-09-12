package main

import (
	"agenda/modules/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.CarregaRotas(router)

	router.Run("localhost:8080")

}
