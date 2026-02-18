package userrepo

import (
	"context"
	"database/sql"
	"gym/internal/domain/user"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) user.IRepository {
	return &repository{db: db}
}

func (r *repository) FindByID(ctx context.Context, id string) (*user.User, error) {
	var u user.User
	query := "SELECT id, name, email, created_at FROM users WHERE id = ?"

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil // ou um erro específico de "não encontrado"
	}
	if err != nil {
		return nil, err
	}

	return &u, nil
}
