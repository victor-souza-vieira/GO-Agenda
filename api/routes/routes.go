package routes

import (
	"agenda/modules/api/controllers"

	"github.com/gin-gonic/gin"
)

func CarregaRotas(router *gin.Engine) {
	router.POST("/contatos", controllers.SalvarContato)
	router.PUT("/contatos/:id", controllers.EditarContato)
	router.PATCH("/contatos", controllers.EditarParcialContato)
	router.DELETE("/contatos/:id", controllers.DeletarContato)
	router.GET("/contatos", controllers.ListarContatos)
	router.GET("/contatos/:id", controllers.BuscarContato)
}
