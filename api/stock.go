package api

import (
	"net/http"
	"stockPredict/pkg/orm"
	"stockPredict/pkg/orm/models"

	"github.com/gin-gonic/gin"
)

func StockRouter(r *gin.RouterGroup) {
	/* in memory */
	stock := r.Group("/stock")
	stock.GET("/", getAllStock)
}

func getAllStock(c *gin.Context) {
	var stock []models.Stock
	orm.Database.Find(&stock)
	c.JSON(http.StatusOK, stock)
}
