// Gerencia a conexão com o banco de dados
package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var dbName string = "database.db"

// Abre uma conexão com o banco de dados.
func OpenDB() (*sql.DB, error) {

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		panic(err)
	}
	fmt.Println("Conected to ", dbName)
	return db, nil
}
