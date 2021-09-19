package controllers

import (
	"agenda/modules/domain/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SalvarContato(gc *gin.Context) {
	var novoContato models.Contato

	if gc.BindJSON(&novoContato) != nil {
		panic("Ocorreu um problema com o bind do Objeto Json")
	}

	lastId, err := models.SalvarContato(novoContato.Nome, novoContato.Email, novoContato.Telefone, novoContato.Celular, novoContato.Endereco)

	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"status": "Houve um erro ao processar sua requisição"})
	} else {
		novoContato.Id = lastId
		gc.IndentedJSON(http.StatusCreated, novoContato)
	}
}

func ListarContatos(gc *gin.Context) {

	contatos := models.ListarContatos()

	if len(contatos) > 0 {
		gc.IndentedJSON(http.StatusOK, contatos)
	} else {
		gc.AbortWithStatus(http.StatusNoContent)
	}
}

func DeletarContato(gc *gin.Context) {
	models.DeletarContato(gc.Param("id"))
	gc.IndentedJSON(http.StatusNoContent, nil)
}

func BuscarContato(gc *gin.Context) {
	contato := models.BuscarContato(gc.Param("id"))

	if contato.Id > 0 {
		gc.IndentedJSON(http.StatusOK, contato)
	} else {
		gc.AbortWithStatus(http.StatusNotFound)
	}

}

func EditarContato(gc *gin.Context) {
	var contatoParaEdicao models.Contato

	if gc.Bind(&contatoParaEdicao) != nil {
		panic("Ocorreu um problema com o bind do objeto JSON para edição.")
	}

	contatoEditado := models.EditarContato(contatoParaEdicao, gc.Param("id"))

	if contatoEditado.Id > 0 {
		gc.IndentedJSON(http.StatusOK, contatoEditado)
	} else {
		gc.AbortWithStatus(http.StatusBadRequest)
	}

}

func EditarParcialContato(gc *gin.Context) {
	var contatoParaEditar models.Contato
	if gc.BindQuery(&contatoParaEditar) != nil {
		panic("Erro ao realizar o bind das variaveis")
	}

	if contatoParaEditar.Id == 0 {
		gc.AbortWithStatus(http.StatusBadRequest)
		panic("Id de contato não informado na requisição")
	}

	fmt.Println(contatoParaEditar)
	contatoEditado := models.EditarParcialContato(contatoParaEditar)

	if contatoEditado.Id > 0 {
		gc.IndentedJSON(http.StatusOK, contatoEditado)
	} else {
		gc.AbortWithStatus(http.StatusBadRequest)
	}
}
