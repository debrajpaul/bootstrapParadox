package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	vision "cloud.google.com/go/vision/apiv1"
	"github.com/gin-gonic/gin"
)

func GetTheText() gin.HandlerFunc {

	return func(c *gin.Context) {

		result := make([]string, 0)

		ctx := context.Background()

		// Creates a client.
		client, err := vision.NewImageAnnotatorClient(ctx)
		if err != nil {
			log.Fatalf("Failed to create client: %v", err)
		}
		defer client.Close()

		// Sets the name of the image file to annotate.
		filename := "/home/suhail/Downloads/bootstrapParadox/api/src/testdata/bill5.jpeg"

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
		

		

		fmt.Println("Labels:",labels);
		for _, label := range labels {
			fmt.Println(label.Description)
			result = append(result, label.Description)
		}

		c.JSON(http.StatusOK, result)
	}
}
