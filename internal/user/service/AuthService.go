package service

import (
	"context"
	"genx/internal/user/model"
)

// AuthService authentication service
type AuthService interface {
	Login(ctx context.Context, user model.LoginRequest) (*model.Token, error)
}

type authService struct {
	model.UserRepository
}

// NewAuthService create new service
func NewAuthService(repo model.UserRepository) AuthService {
	return &authService{repo}
}

func (a *authService) Login(ctx context.Context, req model.LoginRequest) (*model.Token, error) {
	// Validate request
	// Find user by email
	usr, err := a.UserRepository.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	// Check password
	claims := model.Claims{
		UserID: usr.ID,
	}
	accessToken, err := a.generateToken(claims)
	if err != nil {
		return nil, err
	}
	token := &model.Token{
		AccessToken: accessToken,
	}
	return token, nil
}

func (a *authService) generateToken(claims model.Claims) (string, error) {
	return "token", nil
}
