// Pacote peincipal da aplicacao
package main

import (
	_ "github/felipex/kanban-server/docs"
	"github/felipex/kanban-server/internal/server"
)

func main() {
	server.InitServer()
}
