package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type D struct {
	DB *gorm.DB
}

func New() *D {
	dialector := postgres.New(postgres.Config{
		DSN:                  "host=localhost port=5432 user=aliffatulmf password=aliffatulmf dbname=modeling sslmode=disable TimeZone=Asia/Jakarta",
		PreferSimpleProtocol: true,
	})

	startTime := time.Now()
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Time: %s\n", time.Since(startTime))
	return &D{db}
}

func (d *D) Close() {
	db, _ := d.DB.DB()
	db.Close()
}
