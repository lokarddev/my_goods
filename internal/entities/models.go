package entities

import "gorm.io/gorm"

type Dish struct {
	gorm.Model
	Goods       []Goods `gorm:"many2many:dish_goods;"`
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
	Dishes      []Dish `gorm:"many2many:list_dishes;"`
	Title       string
	Description string
}

type User struct {
	gorm.Model
	Login string `json:"login" gorm:"unique"`
	Pass  string `json:"pass"`
}

type Token struct {
	gorm.Model
	User   User `json:"user" gorm:"foreignKey:UserID"`
	UserID int  `json:"user_id"`
}
