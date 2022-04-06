package entity

import (
	"github.com/jackc/pgtype"
	"time"
)

type PgxSession struct {
	Id           pgtype.Int4        `json:"id" db:"id"`
	UserId       pgtype.Int4        `json:"user_id" db:"user_id"`
	CreatedAt    pgtype.Timestamptz `json:"created_at" db:"created_at"`
	RefreshToken pgtype.UUID        `json:"refresh_token" db:"refresh_token"`
	UserAgent    pgtype.Varchar     `json:"user_agent" db:"user_agent"`
	Fingerprint  pgtype.Varchar     `json:"fingerprint" db:"fingerprint"`
	Ip           pgtype.Varchar     `json:"ip" db:"ip"`
	ExpiresIn    pgtype.Int8        `json:"expires_in" db:"expires_in"`
}

func (m *PgxSession) ToClean() *Session {
	var refresh string
	_ = m.RefreshToken.AssignTo(&refresh)

	return &Session{
		Id:           m.Id.Int,
		UserId:       m.UserId.Int,
		CreatedAt:    m.CreatedAt.Time,
		RefreshToken: refresh,
		UserAgent:    m.UserAgent.String,
		Fingerprint:  m.Fingerprint.String,
		Ip:           m.Ip.String,
		ExpiresIn:    m.ExpiresIn.Int,
	}
}

type Session struct {
	Id           int32     `json:"id" db:"id"`
	UserId       int32     `json:"user_id" db:"user_id"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	RefreshToken string    `json:"refresh_token" db:"refresh_token"`
	UserAgent    string    `json:"user_agent" db:"user_agent"`
	Fingerprint  string    `json:"fingerprint" db:"fingerprint"`
	Ip           string    `json:"ip" db:"ip"`
	ExpiresIn    int64     `json:"expires_in" db:"expires_in"`
}

func (m *Session) ToPgx() (*PgxSession, error) {
	session := PgxSession{}
	err := session.Id.Set(m.Id)
	err = session.UserId.Set(m.UserId)
	err = session.CreatedAt.Set(m.CreatedAt)
	err = session.RefreshToken.Set(m.RefreshToken)
	err = session.UserAgent.Set(m.UserAgent)
	err = session.Fingerprint.Set(m.Fingerprint)
	err = session.Ip.Set(m.Ip)
	err = session.ExpiresIn.Set(m.ExpiresIn)
	return &session, err
}
