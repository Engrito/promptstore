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
	Service UserRepository
}

func NewService(repo UserRepository) *Service {
	return &Service{
		Service: repo,
	}
}

func (s *Service) Add(ctx context.Context, user User) (ID, error) {
	return s.Service.Add(ctx, user)
}

func (s *Service) GetById(ctx context.Context, id ID) (User, error) {
	return s.Service.GetById(ctx, id)
}

func (s *Service) GetAll(ctx context.Context) ([]User, error) {
	return s.Service.GetAll(ctx)
}

func (s *Service) UpdateByID(ctx context.Context, id ID, user User) error {
	return s.Service.UpdateByID(ctx, id, user)
}

func (s *Service) DeleteByID(ctx context.Context, id ID) error {
	return s.Service.DeleteByID(ctx, id)
}
