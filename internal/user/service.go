package user

import "context"

type UserService interface {
	Add(ctx context.Context, user User) (ID, error)
	GetById(ctx context.Context, id ID) (User, error)
	GetAll(ctx context.Context) ([]User, error)
	UpdateByID(ctx context.Context, id ID, user User) error
	DeleteByID(ctx context.Context, id ID) error
}

type Service struct {
	Service UserService
}

func (s *Service) Add(ctx context.Context, user User) (ID, error) {
	return ID(""), ErrNotImplemented
}

func (s *Service) GetById(ctx context.Context, id ID) (User, error) {
	return User{}, ErrNotImplemented
}

func (s *Service) GetAll(ctx context.Context) ([]User, error) {
	return []User{}, ErrNotImplemented
}

func (s *Service) UpdateByID(ctx context.Context, id ID, user User) error {
	return ErrNotImplemented
}

func (s *Service) DeleteByID(ctx context.Context, id ID, user User) error {
	return ErrNotImplemented
}
