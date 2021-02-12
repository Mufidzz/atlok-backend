package structs

import "gorm.io/gorm"

type Substation struct {
	gorm.Model
	Code      string
	Name      string
	Latitude  string
	Longitude string
}
