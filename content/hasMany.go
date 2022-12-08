package content

import (
	"test/models"

	"gorm.io/gorm"
)

var customer = []models.Customer{
	{
		Name:       "Ares",
		CreditCard: credit1,
	},
	{
		Name:       "Athur",
		CreditCard: credit2,
	},
	{
		Name:       "Bryan",
		CreditCard: credit3,
	},
}

var credit1 = []models.CreditCard{
	{
		Number:     0123456,
		CustomerID: 1,
	},
}

var credit2 = []models.CreditCard{
	{
		Number:     0123456,
		CustomerID: 2,
	},
	{
		Number:     9080706,
		CustomerID: 2,
	},
}

var credit3 = []models.CreditCard{
	{
		Number:     0123456,
		CustomerID: 3,
	},
	{
		Number:     9080706,
		CustomerID: 3,
	},
	{
		Number:     7880323,
		CustomerID: 3,
	},
}

func InjectHasMany(db *gorm.DB) {
	db.Create(&customer)
}
