package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jonasfsilva/go-markets/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/markets", controllers.MarketList)
	// r.GET("/:nome", controllers.MarketList)
	r.POST("/markets", controllers.MarketCreate)
	// r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	r.Run()
}
