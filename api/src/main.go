package main

import (
	"bootstrapParadox/src/handler"
	"bootstrapParadox/src/modules/inventory"

	"github.com/gin-gonic/gin"
)

func main() {
	inventory.InitialMigration()
	product := inventory.NewInstance()
	router := gin.Default()

	router.GET("/ocr", handler.GetTheText())
	router.GET("/items", handler.GetItems())
	router.GET("/item/:name", handler.GetItem(product))
	router.POST("/item", handler.PostItems(product))
	router.DELETE("/item/:name", handler.DeleteItem(product))

	router.Run() // listen and serve on 0.0.0.0:8080
}
