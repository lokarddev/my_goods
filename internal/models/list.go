package models

import "gorm.io/gorm"

type List struct {
	gorm.Model
	Title       string
	Description string
}
