package models

import (
	"agenda/modules/db"
	"strconv"
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

func ListarContatos() []Contato {
	db := db.ConectarBancoDados()
	defer db.Close()

	selectTodosContatos, err := db.Query("SELECT * FROM contatos ORDER BY nome ASC")

	if err != nil {
		panic(err.Error())
	}

	contato := Contato{}
	contatos := []Contato{}

	for selectTodosContatos.Next() {
		err := selectTodosContatos.Scan(&contato.Id, &contato.Nome, &contato.Email, &contato.Telefone, &contato.Celular, &contato.Endereco)
		if err != nil {
			panic(err.Error())
		}
		contatos = append(contatos, contato)
	}

	return contatos
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
