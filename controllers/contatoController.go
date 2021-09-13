package controllers

import (
	"agenda/modules/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SalvarContato(gc *gin.Context) {
	var novoContato models.Contato

	if gc.BindJSON(&novoContato) != nil {
		panic("Ocorreu um problema com o bind do Objeto Json")
	}

	lastId, err := models.SalvarContato(novoContato.Nome, novoContato.Email, novoContato.Telefone, novoContato.Celular, novoContato.Endereco)

	if err == nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"status": "Houve um erro ao processar sua requisição"})
	} else {
		novoContato.Id = lastId
		gc.IndentedJSON(http.StatusCreated, novoContato)
	}

}
