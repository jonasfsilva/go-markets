package main

import (
	"fmt"

	"github.com/jonasfsilva/go-markets/routes"
)

func main() {
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleResquest()
}
