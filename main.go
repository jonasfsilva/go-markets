package main

import (
	"fmt"

	"github.com/jonasfsilva/go-markets/database"
	"github.com/jonasfsilva/go-markets/routes"
)

func main() {
	database.ConectDataBase()
	fmt.Println("Iniciando o servidor Rest com Go")
	// database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
