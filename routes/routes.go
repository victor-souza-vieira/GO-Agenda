package routes

import (
	"agenda/modules/controllers"

	"github.com/gin-gonic/gin"
)

func CarregaRotas(router *gin.Engine) {
	router.POST("/contatos", controllers.SalvarContato)
	router.GET("/contatos", controllers.ListarContatos)
}
