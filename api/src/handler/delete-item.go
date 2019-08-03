package handler

import (
	"bootstrapParadox/src/modules/inventory"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func DeleteItem(product *inventory.Product) gin.HandlerFunc {

	return func(c *gin.Context) {

		db, err := gorm.Open("sqlite3", "test.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		log.Println("name--> ", c.Param("name"))
		db.Where("Name = ?", c.Param("name")).Find(&product)
		db.Delete(&product)

		c.JSON(http.StatusOK, map[string]string{
			"message": "successfullly deleted",
		})
	}
}
