package main

import (
	"github.com/gin-gonic/gin"
	"stockPredict/api"

	"stockPredict/pkg/orm"
)

func main() {
	/* gin */
	router := gin.Default()
	v1 := router.Group("/v1")
	api.StockRouter(v1)
	
	/* orm */
	go func() {
		orm.Connect()
	}()
	router.Run(":3000")
}
