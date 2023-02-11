package user

import (
	"context"
	"database/sql"
)

type postgres struct {
	db *sql.DB
}

func (db *postgres) GetByEmail(ctx context.Context, email string) (Model, error) {
	row := db.db.QueryRowContext(ctx, `
SELECT id,
       first_name,
       last_name,
       email
FROM   account
WHERE  email = $1`, email)
	if row.Err() != nil {
		return Model{}, row.Err()
	}

	var user Model
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)

	return user, err
}

func (db *postgres) Create(ctx context.Context, firstName, lastName, email, pwdHash string) (int, error) {
	var userID int
	err := db.db.QueryRowContext(ctx, `
INSERT INTO account
            (first_name,
             last_name,
             email,
             pwd_hash)
VALUES     ($1,
            $2,
            $3,
            $4)
RETURNING id`,
		firstName,
		lastName,
		email,
		pwdHash).Scan(&userID)

	return userID, err
}

func (db *postgres) GetByID(ctx context.Context, userID int) (Model, error) {
	row := db.db.QueryRowContext(ctx, `
SELECT id,
       first_name,
       last_name,
       email
FROM   account
WHERE  id = $1`, userID)
	if row.Err() != nil {
		return Model{}, row.Err()
	}

	var user Model
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)

	return user, err
}

func (db *postgres) Search(ctx context.Context, firstName, lastName, email string, limit, offset int) ([]Model, error) {
	panic("unimplemented")
}

func NewPostgresRepo(db *sql.DB) Repository {
	return &postgres{db: db}
}
