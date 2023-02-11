package user

import "context"

type Repository interface {
	Create(ctx context.Context, firstName, lastName, email, pwdHash string) (int, error)
	GetByID(ctx context.Context, userID int) (Model, error)
	GetByEmail(ctx context.Context, email string) (Model, error)
	Search(ctx context.Context, firstName, lastName, email string, limit, offset int) ([]Model, error)
}
