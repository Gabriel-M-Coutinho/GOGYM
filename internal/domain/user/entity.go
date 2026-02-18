package user

import "time"

type User struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
}

// Regra de negócio vive aqui
func (u *User) IsValid() bool {
	return u.Name != "" && u.Email != ""
}
