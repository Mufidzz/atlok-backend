package structs

import "gorm.io/gorm"

type CustomerChange struct {
	gorm.Model
	CurrentCustomerID uint
	NewCustomerID     uint
}

func (CustomerChangeWCustomer) TableName() string {
	return "customer_changes"
}

type CustomerChangeWCustomer struct {
	CustomerChange
	CurrentCustomer Customer `gorm:"foreignkey:ID;references:CurrentCustomerID"`
}
