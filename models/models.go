package models

import "gorm.io/gorm"

// belongs to model
type Employee struct {
	gorm.Model
	Name      string  `json:"name"`
	CompanyID uint    `json:"company_id"`
	Company   Company `json:"company"`
}

type Company struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// has many model
type Customer struct {
	gorm.Model
	Name       string       `json:"name"`
	CreditCard []CreditCard `json:"credit_card"`
}

type CreditCard struct {
	gorm.Model
	Number     uint `json:"number"`
	CustomerID uint `json:"customer_id"`
}
