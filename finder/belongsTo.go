package finder

import (
	"encoding/json"
	"fmt"
	"log"

	"test/models"

	"gorm.io/gorm"
)

func FindBelongsTo(db *gorm.DB, query string) {
	var employees []models.Employee

	err := db.Preload("Company").Where(query).Find(&employees, `company_id = 2`).Error
	if err != nil {
		log.Fatal(err)
	}

	bin, _ := json.MarshalIndent(employees, "", "  ")

	fmt.Println(string(bin))
}
