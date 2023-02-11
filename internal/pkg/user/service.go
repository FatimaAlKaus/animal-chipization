package user

import (
	"context"
	"database/sql"
	"errors"
	"main/pkg/hash"
)

type Service interface {
	Register(ctx context.Context, req RequestRegistration) (ResponseRegistration, error)
}

type service struct {
	rep Repository
}

func (s *service) Register(ctx context.Context, req RequestRegistration) (ResponseRegistration, error) {
	err := s.checkEmailExists(ctx, req.Email)
	if err != nil {
		return ResponseRegistration{}, err
	}

	hash := hash.Password(req.Password)

	userID, err := s.rep.Create(ctx, req.FirstName, req.LastName, req.Email, hash)
	if err != nil {
		return ResponseRegistration{}, err
	}

	return ResponseRegistration{
		ID:        userID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	}, nil
}

type serviceError struct {
	message string
}

func (e serviceError) Error() string {
	return e.message
}

func (s *service) checkEmailExists(ctx context.Context, email string) error {
	_, err := s.rep.GetByEmail(ctx, email)
	if err == nil {
		return serviceError{message: "Пользователь с таким email уже существует"}
	}

	if errors.Is(err, sql.ErrNoRows) {
		return nil
	}

	return err
}

func NewService(rep Repository) Service {
	return &service{rep: rep}
}

type (
	ResponseRegistration struct {
		ID        int    `json:"id"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
	}
)
