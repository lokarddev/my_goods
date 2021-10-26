package entity

import "gorm.io/gorm"

type Dish struct {
	gorm.Model
	Goods       *[]Goods `gorm:"many2many:dish_goods;"`
	Title       string
	Description string
}

type Goods struct {
	gorm.Model
	Title       string
	Description string
}

type List struct {
	gorm.Model
	Dishes      *[]Dish `gorm:"many2many:list_dishes;"`
	Title       string
	Description string
}
