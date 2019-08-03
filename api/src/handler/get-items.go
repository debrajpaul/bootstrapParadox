package handler

import (
	"bootstrapParadox/src/modules/inventory"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetItems() gin.HandlerFunc {

	return func(c *gin.Context) {
		db, err := gorm.Open("sqlite3", "test.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		product := []*inventory.Product{}
		db.Find(&product)

		c.JSON(http.StatusOK, product)
	}
}
