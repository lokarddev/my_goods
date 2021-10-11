package models

import "gorm.io/gorm"

type Dish struct {
	gorm.Model
	Goods       *[]Goods `gorm:"many2many:dish_goods;"`
	Title       string
	Description string
}
