package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jonasfsilva/go-markets/controllers"
)

func HandleResquest() {
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8000"
	}

	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/markets", controllers.ListMarkets)
	log.Fatal(http.ListenAndServe(":8000", r))
}
