package structs

import "gorm.io/gorm"

type PowerRating struct {
	gorm.Model
	Code string
	Name string
}
