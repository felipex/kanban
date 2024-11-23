// Gerencia a conexão com o banco de dados
package database

import (
	"database/sql"
	"fmt"
	"github/felipex/kanban-server/configs"

	_ "github.com/mattn/go-sqlite3"
)

var dbName string = "database.db"

// Abre uma conexão com o banco de dados.
func OpenDB() (*sql.DB, error) {

	db, err := sql.Open("sqlite3", configs.DATABASE_NAME)
	if err != nil {
		panic(err)
	}
	fmt.Println("Conected to ", dbName)
	return db, nil
}
