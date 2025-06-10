package auth

import (
	"errors"
	"hospital-management/internal/models"
	"hospital-management/internal/user"
	"hospital-management/pkg/utils"
)

type Service struct {
	userService *user.Service
	jwtSecret   string
}

func NewService(userService *user.Service, jwtSecret string) *Service {
	return &Service{
		userService: userService,
		jwtSecret:   jwtSecret,
	}
}

func (s *Service) Login(req models.LoginRequest) (string, *models.User, error) {
	user, err := s.userService.GetByUsername(req.Username)
	if err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return "", nil, errors.New("invalid credentials")
	}

	if !user.IsActive {
		return "", nil, errors.New("account is deactivated")
	}

	token, err := utils.GenerateToken(user.ID, user.Role, s.jwtSecret)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}

func (s *Service) Register(req models.RegisterRequest) (*models.User, error) {
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  hashedPassword,
		Role:      req.Role,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Phone:     req.Phone,
		IsActive:  true,
	}

	return s.userService.Create(user)
}
