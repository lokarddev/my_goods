package entities

import "github.com/jackc/pgtype"

type Dish struct {
	ID          pgtype.Int4    `json:"id" db:"id"`
	Title       pgtype.Varchar `json:"title" db:"title"`
	Description pgtype.Varchar `json:"description" db:"description"`
}

type Goods struct {
	Title       string
	Description string
}

type List struct {
	Dishes      []Dish
	Title       string
	Description string
}

type User struct {
	Login string `json:"login"`
	Pass  string `json:"pass"`
}

type Token struct {
	User   User `json:"user"`
	UserID int  `json:"user_id"`
}
