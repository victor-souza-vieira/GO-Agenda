package models

import (
	"agenda/modules/db"
	exceptions "agenda/modules/exception"
	"strconv"
)

type Contato struct {
	Id       int64  `json:"id" form:"id"`
	Nome     string `json:"nome" form:"nome"`
	Email    string `json:"email" form:"email"`
	Telefone string `json:"telefone" form:"telefone"`
	Celular  string `json:"celular" form:"celular"`
	Endereco string `json:"endereco" form:"endereco"`
}

func SalvarContato(nome string, email string, telefone string, celular string, endereco string) (int64, error) {
	db := db.ConectarBancoDados()
	defer db.Close()

	salvarNoBanco, err := db.Prepare("INSERT INTO contatos() values(null, ?,?,?,?,?)")

	if err != nil {
		return 0, exceptions.CustomError{
			Erro:       err.Error(),
			Data:       "Não foi possível salvar o contato",
			StatusCode: 500}
	}

	sqlResult, err := salvarNoBanco.Exec(nome, email, telefone, celular, endereco)

	if err != nil {
		return 0, exceptions.CustomError{
			Erro:       err.Error(),
			Data:       "Não foi possível salvar o contato",
			StatusCode: 500}
	}

	return sqlResult.LastInsertId()
}

// TODO Construir funcao para retornar o erro baseado apenas na mensagem
/*
return nil, exceptions.CustomError{
			Erro:       err.Error(),
			Data:       "Não foi possível realizar sua consulta.",
			StatusCode: 500}
*/

func ListarContatos() ([]Contato, error) {
	db := db.ConectarBancoDados()
	defer db.Close()

	selectTodosContatos, err := db.Query("SELECT * FROM contatos ORDER BY nome ASC")

	if err != nil {
		return nil, exceptions.CustomError{
			Erro:       err.Error(),
			Data:       "Não foi possível realizar sua consulta.",
			StatusCode: 500}
	}

	contato := Contato{}
	contatos := []Contato{}

	for selectTodosContatos.Next() {
		err := selectTodosContatos.Scan(&contato.Id, &contato.Nome, &contato.Email, &contato.Telefone, &contato.Celular, &contato.Endereco)
		if err != nil {
			return nil, exceptions.CustomError{
				Erro:       err.Error(),
				Data:       "Houve um erro no mapeamento da consulta.",
				StatusCode: 500}
		}
		contatos = append(contatos, contato)
	}

	return contatos, nil
}

func DeletarContato(idContato string) {

	db := db.ConectarBancoDados()
	defer db.Close()

	contatoDeletado, err := db.Prepare("DELETE FROM contatos WHERE id = ?")

	if err != nil {
		panic(err.Error())
	}

	contatoDeletado.Exec(idContato)
}

func BuscarContato(idContato string) Contato {
	db := db.ConectarBancoDados()
	defer db.Close()

	contatoBD, err := db.Query("SELECT * FROM contatos WHERE id = ?", idContato)

	if err != nil {
		panic(err.Error())
	}

	contatoSaida := Contato{}
	if contatoBD.Next() {
		err := contatoBD.Scan(&contatoSaida.Id, &contatoSaida.Nome, &contatoSaida.Email, &contatoSaida.Telefone, &contatoSaida.Celular, &contatoSaida.Endereco)

		if err != nil {
			panic(err.Error())
		}
	}

	return contatoSaida
}

func EditarContato(contato Contato, idContato string) Contato {
	db := db.ConectarBancoDados()
	defer db.Close()

	stmtContato, err := db.Prepare("UPDATE contatos SET nome = ?, email = ?, telefone = ?, celular = ?, endereco = ? WHERE id = ?")

	if err != nil {
		panic(err.Error())
	}

	sqlResult, err := stmtContato.Exec(contato.Nome, contato.Email, contato.Telefone, contato.Celular, contato.Endereco, idContato)

	if err != nil {
		panic(err.Error())
	}

	linhasAfetadas, err := sqlResult.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	if linhasAfetadas <= 0 {
		return Contato{}
	}

	contato.Id, _ = strconv.ParseInt(idContato, 10, 64)

	return contato
}

func EditarParcialContato(contato Contato) Contato {
	db := db.ConectarBancoDados()
	defer db.Close()

	result, err := db.Query("SELECT * FROM contatos WHERE id = ?", contato.Id)

	if err != nil {
		panic(err.Error())
	}

	contatoBD := Contato{}
	if result.Next() {
		err := result.Scan(&contatoBD.Id, &contatoBD.Nome, &contatoBD.Email, &contatoBD.Telefone, &contatoBD.Celular, &contatoBD.Endereco)
		if err != nil {
			panic(err.Error())
		}
	}

	if contato.Nome == "" {
		contato.Nome = contatoBD.Nome
	}
	if contato.Email == "" {
		contato.Email = contatoBD.Email
	}
	if contato.Telefone == "" {
		contato.Telefone = contatoBD.Telefone
	}
	if contato.Celular == "" {
		contato.Celular = contatoBD.Celular
	}
	if contato.Endereco == "" {
		contato.Endereco = contatoBD.Endereco
	}

	return EditarContato(contato, strconv.Itoa(int(contato.Id)))
}
