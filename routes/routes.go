package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jonasfsilva/go-markets/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/markets", controllers.MarketList)
	r.POST("/markets", controllers.MarketCreate)
	r.GET("/markets/:id", controllers.MarketRetrieve)
	r.DELETE("/markets/:id", controllers.MarketDelete)
	r.PUT("/markets/:id", controllers.MarketUpdate)
	r.GET("/markets/search/:name", controllers.SearchMarketFromName)

	r.Run()
}
