package main

import (
	"github.com/jonasfsilva/go-markets/database"
	"github.com/jonasfsilva/go-markets/routes"
)

func main() {
	database.ConectDataBase()
	routes.HandleRequests()
}
