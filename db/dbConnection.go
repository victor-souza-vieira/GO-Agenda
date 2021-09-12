package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

/* Funcao responsavel por devolver um ponteiro para conexao com
para o banco de dados
*/
func ConectarBancoDados() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/ContatinhosGO")

	if err != nil {
		panic(err.Error())
	}

	return db
}
