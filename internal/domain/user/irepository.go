package user

import (
	"context"
)

type IRepository interface {
	FindByID(ctx context.Context, id string) (*User, error)
	//FindByEmail(ctx context.Context, email string) (*User, error)
	//Save(ctx context.Context, user *User) error
	//Delete(ctx context.Context, id string) error
}
