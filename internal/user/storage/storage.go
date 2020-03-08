package storage

import (
	"context"
	"genx/internal/user/model"

	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

// NewUserRepository create new repo
func NewUserRepository(postgredb *sqlx.DB) model.UserRepository {
	return &userRepository{
		db: postgredb,
	}
}

func (r *userRepository) Find(ctx context.Context, filter interface{}) (*model.ResponseUsers, error) {
	usr := &model.ResponseUsers{}
	return usr, nil
}

func (r *userRepository) FindByID(ctx context.Context, id int) (*model.BaseUser, error) {
	usr := &model.BaseUser{}
	return usr, nil
}

func (r *userRepository) Create(ctx context.Context, user model.CreateRequest) (*model.BaseUser, error) {
	usr := &model.BaseUser{}
	return usr, nil
}
func (r *userRepository) Update(ctx context.Context, user model.UpdateRequest) error {
	return nil
}
func (r *userRepository) Delete(ctx context.Context, id int) error {
	return nil
}
