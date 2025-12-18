package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/romitdubey1/user-api/db/sqlc"
)

type UserRepository struct {
	q *sqlc.Queries
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		q: sqlc.New(db),
	}
}

func (r *UserRepository) Create(
	ctx context.Context,
	name string,
	dob time.Time,
) (sqlc.User, error) {

	return r.q.CreateUser(ctx, sqlc.CreateUserParams{
		Name: name,
		Dob:  dob,
	})
}

func (r *UserRepository) GetByID(
	ctx context.Context,
	id int32,
) (sqlc.User, error) {

	return r.q.GetUserByID(ctx, id)
}

func (r *UserRepository) List(ctx context.Context) ([]sqlc.User, error) {
	return r.q.ListUsers(ctx)
}

func (r *UserRepository) Update(
	ctx context.Context,
	id int32,
	name string,
	dob time.Time,
) (sqlc.User, error) {

	return r.q.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:   id,
		Name: name,
		Dob:  dob,
	})
}

func (r *UserRepository) Delete(
	ctx context.Context,
	id int32,
) error {

	return r.q.DeleteUser(ctx, id)
}
