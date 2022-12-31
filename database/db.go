package database

import (
	"log"

	"github.com/jonasfsilva/go-markets/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectDataBase() {
	stringDeConexao := "host=192.168.80.2 user=markets password=markets@123 dbname=markets port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
	DB.AutoMigrate(&models.Market{})
}
