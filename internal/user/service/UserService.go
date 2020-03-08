package service

import (
	"context"
	"genx/internal/user/model"
)

// UserService application layer of usera
type UserService interface {
	Find(ctc context.Context, filter interface{}) (*model.ResponseUsers, error)
	FindByID(ctx context.Context, id int) (*model.BaseUser, error)
	Create(ctx context.Context, user model.CreateRequest) (*model.BaseUser, error)
	Update(ctx context.Context, user model.UpdateRequest) error
	Delete(ctx context.Context, id int) error
}

type userService struct {
	model.UserRepository
}

// NewUserService create new service
func NewUserService(repo model.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) Find(ctx context.Context, filter interface{}) (*model.ResponseUsers, error) {
	return s.UserRepository.Find(ctx, filter)
}

func (s *userService) FindByID(ctx context.Context, id int) (*model.BaseUser, error) {
	usr, err := s.UserRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return usr.Publish(), nil
}

func (s *userService) Create(ctx context.Context, req model.CreateRequest) (*model.BaseUser, error) {
	// Check permission
	// validate request
	// Hash password
	usr, err := s.UserRepository.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return usr.Publish(), nil
}

func (s *userService) Update(ctx context.Context, user model.UpdateRequest) error {
	return s.UserRepository.Update(ctx, user)
}

func (s *userService) Delete(ctx context.Context, id int) error {
	return s.UserRepository.Delete(ctx, id)
}
