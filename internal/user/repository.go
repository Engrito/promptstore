package user

import (
	"context"
	"errors"
	"sync"
)

type UserRepository interface {
	Add(ctx context.Context, user User) (ID, error)
	GetById(ctx context.Context, id ID) (User, error)
	UpdateByID(ctx context.Context, id ID, user User) error
	DeleteByID(ctx context.Context, id ID) error
	GetAll(ctx context.Context) ([]User, error)
}

var (
	ErrNotImplemented = errors.New("not implemented")
)

type InMermoryUserRepository struct {
	mu     sync.Mutex
	ListDB []User
}

func NewInMermoryUserRepository() *InMermoryUserRepository {
	return &InMermoryUserRepository{
		ListDB: make([]User, 0),
	}
}

func (ur *InMermoryUserRepository) Add(ctx context.Context, user User) (ID, error) {
	return "", ErrNotImplemented
}

func (ur *InMermoryUserRepository) GetById(ctx context.Context, id ID) (User, error) {
	return User{}, ErrNotImplemented
}

func (ur *InMermoryUserRepository) UpdateByID(ctx context.Context, id ID, user User) error {
	return ErrNotImplemented
}

func (ur *InMermoryUserRepository) DeleteByID(ctx context.Context, id ID) error {
	return ErrNotImplemented
}

func (ur *InMermoryUserRepository) GetAll(ctx context.Context) ([]User, error) {
	return []User{}, ErrNotImplemented
}
