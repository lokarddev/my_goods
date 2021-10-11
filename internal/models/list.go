package models

import "gorm.io/gorm"

type List struct {
	gorm.Model
	Dishes      *[]Dish `gorm:"many2many:list_dishes;"`
	Title       string
	Description string
}
