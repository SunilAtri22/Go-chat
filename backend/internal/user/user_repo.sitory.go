package user

import (
	"context"

	"github.com/gocql/gocql"
)

type DBTX interface {
	Query(query string, values ...interface{}) *gocql.Query
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	query := "INSERT INTO users (id, username, password, email) VALUES (?, ?, ?, ?)"
	err := r.db.Query(query, gocql.TimeUUID(), user.Username, user.Password, user.Email).Exec()
	if err != nil {
		return &User{}, err
	}

	return user, nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	u := User{}
	query := "SELECT id, email, username, password FROM users WHERE email = ? LIMIT 1"
	iter := r.db.Query(query, email).Iter()

	if iter.Scan(&u.ID, &u.Email, &u.Username, &u.Password) {
		return &u, nil
	}

	return &User{}, gocql.ErrNotFound
}
