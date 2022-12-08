package content

import (
	"log"

	"test/models"

	"gorm.io/gorm"
)

var company = []models.Company{
	{
		ID:   1,
		Name: "Shelby Limited",
	},
	{
		ID:   2,
		Name: "Sabini's Corp",
	},
	{
		ID:   3,
		Name: "John Inc",
	},
}

var employee = []models.Employee{
	{
		Name:      "Thomas",
		CompanyID: 1,
	},
	{
		Name:      "Arthur",
		CompanyID: 1,
	},
	{
		Name:      "John's",
		CompanyID: 1,
	},
	{
		Name:      "Arman",
		CompanyID: 2,
	},
	{
		Name:      "Lance",
		CompanyID: 3,
	},
}

func InjectBelongsTo(db *gorm.DB) {
	if err := db.Create(&company).Error; err != nil {
		log.Fatal(err)
	}
	if err := db.Create(&employee).Error; err != nil {
		log.Fatal(err)
	}
}
