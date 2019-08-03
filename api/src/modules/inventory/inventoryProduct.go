package inventory

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

type Product struct {
	gorm.Model
	Catogery string `json:"category"`
	Owener   string `json:"owener"`
	ItemName string `json:"itemName"`
	Price    int64  `json:"price"`
}

func (product Product) String() string {
	return fmt.Sprintf("Product<%s %s %s %d>", product.Catogery, product.Owener, product.ItemName, product.Price)
}

func NewInstance() *Product {
	return &Product{}
}

func InitialMigration() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Product{})
}
