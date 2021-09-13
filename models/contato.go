package models

import (
	"agenda/modules/db"
)

type Contato struct {
	Id       int64  `json:"id"`
	Nome     string `json:"nome"`
	Email    string `json:"email"`
	Telefone string `json:"telefone"`
	Celular  string `json:"celular"`
	Endereco string `json:"endereco"`
}

/*Funcao responsavel por inserir um contato no banco de dados*/
func SalvarContato(nome string, email string, telefone string, celular string, endereco string) (int64, error) {
	db := db.ConectarBancoDados()
	defer db.Close()

	salvarNoBanco, err := db.Prepare("INSERT INTO contatos() values(null, ?,?,?,?,?)")

	if err != nil {
		panic(err.Error())
	}

	sqlResult, err := salvarNoBanco.Exec(nome, email, telefone, celular, endereco)

	if err != nil {
		panic(err.Error())
	}

	return sqlResult.LastInsertId()
}
