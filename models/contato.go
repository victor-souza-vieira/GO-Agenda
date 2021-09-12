package models

import "agenda/modules/db"

type Contato struct {
	Id       int    `json:id`
	Nome     string `json:nome`
	Email    string `json:email`
	Telefone string `json:telefone`
	Celular  string `json:celular`
	Endereco string `json:endereco`
}

func SalvarContato(nome string, email string, telefone string, celular string, endereco string) {
	db := db.ConectarBancoDados()
	defer db.Close()

	salvarNoBanco, err := db.Prepare("INSERT INTO contatos() values(?,?,?,?,?)")

	if err != nil {
		panic(err.Error())
	}

	salvarNoBanco.Exec(nome, email, telefone, celular, endereco)
}
