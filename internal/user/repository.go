package user

import (
	"context"
	"errors"
	"sync"
	"time"
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
	ErrUserNotFound   = errors.New("user not found")
)

type InMermoryUserRepository struct {
	mu     sync.Mutex
	ListDB []User
	counter int64
}

func NewInMermoryUserRepository() *InMermoryUserRepository {
	return &InMermoryUserRepository{
		ListDB: make([]User, 0),
		counter: 0,
	}
}

func (ur *InMermoryUserRepository) Add(ctx context.Context, user User) (ID, error) {
	ur.mu.Lock()
	defer ur.mu.Unlock()
	
	ur.counter++
	id := ID(string(ur.counter))
	
	now := time.Now()
	user.Id = id
	user.CreatedAt = CreatedAt(now)
	user.UpdatedAt = UpdatedAt(now)
	
	ur.ListDB = append(ur.ListDB, user)
	
	return id, nil
}

func (ur *InMermoryUserRepository) GetById(ctx context.Context, id ID) (User, error) {
	ur.mu.Lock()
	defer ur.mu.Unlock()
	
	for _, u := range ur.ListDB {
		if u.Id == id {
			return u, nil
		}
	}
	
	return User{}, ErrUserNotFound
}

func (ur *InMermoryUserRepository) UpdateByID(ctx context.Context, id ID, user User) error {
	ur.mu.Lock()
	defer ur.mu.Unlock()
	
	for i, u := range ur.ListDB {
		if u.Id == id {
			user.Id = id
			user.UpdatedAt = UpdatedAt(time.Now())
			ur.ListDB[i] = user
			return nil
		}
	}
	
	return ErrUserNotFound
}

func (ur *InMermoryUserRepository) DeleteByID(ctx context.Context, id ID) error {
	ur.mu.Lock()
	defer ur.mu.Unlock()
	
	for i, u := range ur.ListDB {
		if u.Id == id {
			ur.ListDB = append(ur.ListDB[:i], ur.ListDB[i+1:]...)
			return nil
		}
	}
	
	return ErrUserNotFound
}

func (ur *InMermoryUserRepository) GetAll(ctx context.Context) ([]User, error) {
	ur.mu.Lock()
	defer ur.mu.Unlock()
	
	users := make([]User, len(ur.ListDB))
	copy(users, ur.ListDB)
	
	return users, nil
}
