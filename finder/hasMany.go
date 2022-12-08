package finder

import (
	"encoding/json"
	"fmt"

	"test/models"

	"gorm.io/gorm"
)

func FindHasMany(db *gorm.DB, query string) {
	var customer []models.Customer

	db.Model(&models.Customer{}).Where(query).Preload("CreditCard").Find(&customer)

	bin, _ := json.MarshalIndent(customer, "", "  ")

	fmt.Println(string(bin))
}
