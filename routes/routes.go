package routes

import (
	"net/http"

	"agenda/modules/controllers"
)

func CarregaRotas() {
	rotasEscrita()
}

/*Tem como objetivo as rotas que fazem escrita em linhas de tabelas do
banco de dados*/
func rotasEscrita() {
	http.HandleFunc("/contatos", controllers.SalvarContato)
}
