package model

import (
	"context"
	"genx/internal/db"
	"time"
)

// UserRepository interface
type UserRepository interface {
	Find(ctx context.Context, filter interface{}) (*ResponseUsers, error)
	FindByID(ctx context.Context, id int) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, user CreateRequest) (*User, error)
	Update(ctx context.Context, user UpdateRequest) error
	Delete(ctx context.Context, id int) error
}

// BaseUser model
type BaseUser struct {
	ID         int       `json:"id" validate:"required,uuid"`
	FirstName  string    `json:"first_name" validate:"required"`
	LastName   string    `json:"last_name" validate:"required"`
	Email      string    `json:"email" validate:"required,email"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	ArchivedAt time.Time `json:"archived_at"`
}

// User model
type User struct {
	BaseUser
	PasswordHash string `json:"password_hash" validate:"required"`
	PasswordSalt string `json:"password_salt" validate:"required"`
}

// CreateRequest model
type CreateRequest struct {
	FirstName       string `json:"first_name" validate:"required"`
	LastName        string `json:"last_name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required"`
	PasswordConfirm string `json:"password_confirm" validate:"required,eqfield=Password"`
}

// UpdateRequest update information
type UpdateRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
}

// ListOfUser list of users
type ListOfUser []BaseUser

// ResponseUsers list of users
type ResponseUsers struct {
	Meta db.Pagination `json:"meta_data"`
	Data ListOfUser    `json:"data"`
}

// LoginRequest login information
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// Publish convert to Base model
func (u *User) Publish() *BaseUser {
	usr := &BaseUser{
		u.ID,
		u.FirstName,
		u.LastName,
		u.Email,
		u.CreatedAt,
		u.UpdatedAt,
		u.ArchivedAt,
	}
	return usr
}

// Hash convert from plain text to hash passowrd
func (u *User) Hash(pwd string) error {
	return nil
}

// Verify validate password
func (u *User) Verify(pwd string) error {
	return nil
}
