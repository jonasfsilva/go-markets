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

func MarketRetrieve(c *gin.Context) {
	var market models.Market
	id := c.Params.ByName("id")
	database.DB.First(&market, id)

	if market.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Mercado não encontrado"})
		return
	}

	c.JSON(http.StatusOK, market)
}

func MarketDelete(c *gin.Context) {
	var market models.Market
	id := c.Params.ByName("id")
	database.DB.Delete(&market, id)

	c.JSON(http.StatusOK, gin.H{"data": "Mercado deletado com sucesso"})
}

func MarketUpdate(c *gin.Context) {
	var market models.Market
	id := c.Params.ByName("id")
	database.DB.First(&market, id)

	if err := c.ShouldBindJSON(&market); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&market).UpdateColumns(market)
	c.JSON(http.StatusOK, market)
}

func SearchMarketFromName(c *gin.Context) {
	var market models.Market
	name := c.Param("name")
	database.DB.Where(&models.Market{Name: name}).First(&market)

	if market.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Mercado não encontrado"})
		return
	}

	c.JSON(http.StatusOK, market)
}
