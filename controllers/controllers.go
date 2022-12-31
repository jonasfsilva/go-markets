package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jonasfsilva/go-markets/database"
	"github.com/jonasfsilva/go-markets/models"
)

func MarketList(c *gin.Context) {
	var markets []models.Market
	database.DB.Find(&markets)
	c.JSON(200, markets)
}

func MarketCreate(c *gin.Context) {
	var market models.Market
	if err := c.ShouldBindJSON(&market); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&market)
	c.JSON(http.StatusOK, market)
}
