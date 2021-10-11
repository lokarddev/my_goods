package models

import "gorm.io/gorm"

type Dish struct {
	gorm.Model
	Title       string
	Description string
}
