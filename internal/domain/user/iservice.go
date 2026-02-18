package user

import "context"

type Service interface {
	GetByID(ctx context.Context, id string) (*User, error)
	//Create(ctx context.Context, name, email string) (*User, error)
	//Delete(ctx context.Context, id string) error
}
