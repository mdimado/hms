package user

import "hospital-management/internal/models"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(user *models.User) (*models.User, error) {
	if err := s.repo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) GetByID(id uint) (*models.User, error) {
	return s.repo.GetByID(id)
}

func (s *Service) GetByUsername(username string) (*models.User, error) {
	return s.repo.GetByUsername(username)
}

func (s *Service) Update(user *models.User) error {
	return s.repo.Update(user)
}
