package entity

import "github.com/jackc/pgtype"

type PgxMeasure struct {
	Id        pgtype.Int4        `json:"id" db:"id"`
	CreatedAt pgtype.Timestamptz `json:"created_at" db:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at" db:"updated_at"`
	Code      pgtype.Varchar     `json:"code" db:"code"`
	Value     pgtype.Varchar     `json:"value" db:"value"`
}

type Measure struct {
}
