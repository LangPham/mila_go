package repo

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)
var Repo *gorm.DB

func init() {
	dsn := "host=127.0.0.1 user=postgres password=postgres dbname=mila port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("ERROR: fail create postgres session, %s", err.Error())
		os.Exit(1)
	}
	Repo = db
}
