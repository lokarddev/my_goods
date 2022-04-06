package entity

import (
	"github.com/jackc/pgtype"
	"time"
)

type PgxUser struct {
	Id        pgtype.Int4        `json:"id" db:"id"`
	CreatedAt pgtype.Timestamptz `json:"created_at" db:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at" db:"updated_at"`
	Name      pgtype.Varchar     `json:"name" db:"name"`
	Password  pgtype.Varchar     `json:"password" db:"password"`
}

func (m *PgxUser) ToClean() *User {
	return &User{
		Id:        m.Id.Int,
		CreatedAt: m.CreatedAt.Time,
		UpdatedAt: m.UpdatedAt.Time,
		Name:      m.Name.String,
		Password:  m.Password.String,
	}
}

type User struct {
	Id        int32     `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Name      string    `json:"name" db:"name"`
	Password  string    `json:"password" db:"password"`
}

func (m *User) ToPgx() (*PgxUser, error) {
	user := PgxUser{}
	err := user.Id.Set(m.Id)
	err = user.CreatedAt.Set(m.CreatedAt)
	err = user.UpdatedAt.Set(m.UpdatedAt)
	err = user.Name.Set(m.Name)
	err = user.Password.Set(m.Password)
	return &user, err
}
