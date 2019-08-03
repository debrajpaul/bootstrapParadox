package handler

import (
	"bootstrapParadox/src/modules/inventory"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	vision "cloud.google.com/go/vision/apiv1"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetTheText(product *inventory.Product) gin.HandlerFunc {

	return func(c *gin.Context) {

		// DB connection
		db, err := gorm.Open("sqlite3", "test.db")
		if err != nil {
			log.Fatal(err)
		}

		result := make([]string, 0)
		index := 0

		ctx := context.Background()

		// Creates a client.
		client, err := vision.NewImageAnnotatorClient(ctx)
		if err != nil {
			log.Fatalf("Failed to create client: %v", err)
		}
		defer client.Close()

		// Sets the name of the image file to annotate.
		filename := "/home/paul/Documents/github/bootstrapParadox/api/src/testdata/bill5.jpeg"

		file, err := os.Open(filename)
		if err != nil {
			log.Fatalf("Failed to read file: %v", err)
		}
		defer file.Close()
		// image
		image, err := vision.NewImageFromReader(file)
		if err != nil {
			log.Fatalf("Failed to create image: %v", err)
		}

		labels, err := client.DetectTexts(ctx, image, nil, 10)
		if err != nil {
			log.Fatalf("Failed to detect labels: %v", err)
		}

		// labels, err := client.DetectDocumentText(ctx, image, nil)
		// if err != nil {
		// 	log.Fatalf("Failed to detect labels: %v", err)
		// }

		fmt.Println("Labels:- ")
		for _, label := range labels {
			fmt.Println(label.Description)
			result = append(result, label.Description)
		}

		result = strings.Split(result[0], "\n")

		product.Catogery = "food"
		product.Owener = result[0]
		for i, data := range result {
			if data == "Qty Price Amount" {
				index = i
				return
			}
		}
		product.ItemName = result[index+1]
		product.Price = 20

		db.Create(&product)

		c.JSON(http.StatusOK, result)
	}
}
