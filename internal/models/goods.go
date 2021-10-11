package models

import "gorm.io/gorm"

type Goods struct {
	gorm.Model
	Title       string
	Description string
}
