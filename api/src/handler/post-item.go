package handler

import (
	"bootstrapParadox/src/modules/inventory"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type postItemsPostRequest struct {
	Name          string `json:"name"`
	Desc          string `json:"desc"`
	Price         int64  `json:"price"`
	NumberOfItems int64  `json:"numberOfItems"`
}

func PostItems(product *inventory.Product) gin.HandlerFunc {

	return func(c *gin.Context) {

		reqBody := postItemsPostRequest{}
		err := c.Bind(&reqBody)
		if err != nil {
			log.Fatal(err)
		}

		// DB connection
		db, err := gorm.Open("sqlite3", "test.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		product.Name = reqBody.Name
		product.Desc = reqBody.Desc
		product.Price = reqBody.Price
		product.NumberOfItems = reqBody.NumberOfItems

		db.Create(&product)

		c.JSON(http.StatusOK, map[string]string{
			"message": "successfullly created",
		})
	}
}
