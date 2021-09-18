package routes

import (
	"agenda/modules/controllers"

	"github.com/gin-gonic/gin"
)

func CarregaRotas(router *gin.Engine) {
	router.POST("/contatos", controllers.SalvarContato)
	router.GET("/contatos", controllers.ListarContatos)
	router.DELETE("/contatos/:id", controllers.DeletarContato)
	router.GET("/contatos/:id", controllers.BuscarContato)
	router.POST("/contatos/:id", controllers.EditarContato)
}
