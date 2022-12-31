package models

import "gorm.io/gorm"

type Market struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}

var Markets []Market
