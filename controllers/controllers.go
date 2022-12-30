package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jonasfsilva/go-markets/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "I'm Working")
}

func ListMarkets(w http.ResponseWriter, r *http.Request) {
	models.Markets = []models.Market{
		{Id: 1, Name: "Nome 1", Description: "Market 1"},
		{Id: 2, Name: "Nome 2", Description: "Market 2"},
	}
	json.NewEncoder(w).Encode(models.Markets)
}
