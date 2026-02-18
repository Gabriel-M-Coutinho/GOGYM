package userservice

import (
	"context"
	"gym/internal/domain/user"
)

type service struct {
	repo user.IRepository
}

func NewService(repo user.IRepository) user.Service {
	return &service{repo: repo}
}

func (s *service) GetByID(ctx context.Context, id string) (*user.User, error) {
	return s.repo.FindByID(ctx, id)
}
