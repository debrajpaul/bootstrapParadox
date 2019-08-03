package handler

import (
	"bootstrapParadox/src/modules/inventory"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetItem(product *inventory.Product) gin.HandlerFunc {

	return func(c *gin.Context) {
		db, err := gorm.Open("sqlite3", "test.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		log.Println("itemName--> ", c.Param("itemName"))
		db.Where("itemName = ?", c.Param("itemName")).First(&product)

		c.JSON(http.StatusOK, product)
	}
}
