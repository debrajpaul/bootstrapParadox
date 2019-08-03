package inventory

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

type Product struct {
	gorm.Model
	Name          string `json:"name"`
	Desc          string `json:"desc"`
	Price         int64  `json:"price"`
	NumberOfItems int64  `json:"numberOfItems"`
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

func NewInstance() *Product {
	return &Product{}
}
