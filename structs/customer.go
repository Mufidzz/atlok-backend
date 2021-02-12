package structs

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	IDPEL         string
	FullName      string
	Address       string
	PoleNumber    string
	SubstationID  uint
	PowerRatingID uint
	KWHNumber     string
	KWHBrand      string
	KWHYear       string
	Latitude      string
	Longitude     string
	Verified      bool `gorm:"default:false"`
}

func (CustomerWSubstationPowerRating) TableName() string {
	return "customers"
}

type CustomerWSubstationPowerRating struct {
	Customer
	Substation  Substation  `gorm:"foreignkey:SubstationID;references:ID"`
	PowerRating PowerRating `gorm:"foreignkey:PowerRatingID;references:ID"`
}
