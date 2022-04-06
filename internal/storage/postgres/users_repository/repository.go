package users_repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"log"
	"my_goods/internal/entity"
	"my_goods/internal/entity/dto"
	"my_goods/internal/storage/postgres"
	"my_goods/pkg/env"
	"time"
)

type UsersRepository struct {
	postgres.Repository
}

func NewUsersRepository(db postgres.PgxPoolInterface) *UsersRepository {
	return &UsersRepository{Repository: postgres.Repository{
		DB:  db,
		Ctx: context.Background(),
	}}
}

func (r *UsersRepository) GetUserByName(userName string) (entity.User, bool) {
	user := entity.PgxUser{}
	query := fmt.Sprintf("SELECT id, name, password FROM %s WHERE name=$1", postgres.Users)
	rows, err := r.DB.Query(r.Ctx, query, userName)
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Name, &user.Password); err != nil {
			return *user.ToClean(), false
		}
	}
	switch user.Id.Status {
	case pgtype.Present:
		return *user.ToClean(), true
	default:
		return *user.ToClean(), false
	}
}

func (r *UsersRepository) CreateUser(input dto.LoginRequest) (entity.User, error) {
	user := entity.PgxUser{}
	query := fmt.Sprintf("INSERT INTO %s (name, password) VALUES ($1, $2) RETURNING id, name, password", postgres.Users)
	err := r.DB.BeginTxFunc(r.Ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		rows, err := r.DB.Query(r.Ctx, query, input.Name, input.Password)
		defer rows.Close()
		for rows.Next() {
			if err = rows.Scan(&user.Id, &user.Name, &user.Password); err != nil {
				log.Println(err)
			}
		}
		return err
	})
	return *user.ToClean(), err
}

func (r *UsersRepository) CreateSession(userId int32, refresh string) (entity.Session, error) {
	session := entity.PgxSession{}
	expiresIn := time.Now().Add(time.Duration(env.RefreshTTL) * 24 * time.Hour).Unix()

	if err := r.deleteSession(userId); err != nil {
		return *session.ToClean(), err
	}

	query := fmt.Sprintf("INSERT INTO %s (user_id, refresh_token, expires_in) VALUES ($1, $2, $3) "+
		"RETURNING id, user_id, refresh_token, expires_in", postgres.Session)
	err := r.DB.BeginTxFunc(r.Ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		rows, err := r.DB.Query(r.Ctx, query, userId, refresh, expiresIn)
		defer rows.Close()
		for rows.Next() {
			if err = rows.Scan(&session.Id, &session.UserId, &session.RefreshToken, &session.ExpiresIn); err != nil {
				log.Println(err)
			}
		}
		return err
	})
	return *session.ToClean(), err
}

func (r *UsersRepository) deleteSession(userId int32) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id=$1", postgres.Session)
	err := r.DB.BeginTxFunc(r.Ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		_, err := r.DB.Exec(r.Ctx, query, userId)
		if err != nil {
			log.Println(err)
		}
		return err
	})
	return err
}

func (r *UsersRepository) GetSession(token string) (entity.Session, error) {
	session := entity.PgxSession{}
	query := fmt.Sprintf("SELECT id, user_id, refresh_token, expires_in FROM %s WHERE refresh_token=$1",
		postgres.Session)
	rows, err := r.DB.Query(r.Ctx, query, token)
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&session.Id, &session.UserId, &session.RefreshToken, &session.ExpiresIn); err != nil {
			log.Println(err)
		}
	}
	return *session.ToClean(), err
}
