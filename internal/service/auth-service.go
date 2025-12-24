package service

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mardini1414/simple-procurement-system-be/internal/config"
	"github.com/mardini1414/simple-procurement-system-be/internal/dto"
	"github.com/mardini1414/simple-procurement-system-be/internal/model"
	"github.com/mardini1414/simple-procurement-system-be/internal/repository"
	"github.com/mardini1414/simple-procurement-system-be/pkg"
)

type AuthService struct {
	r   *repository.UserRepository
	cfg *config.Config
}

func NewAuthService(r *repository.UserRepository, cfg *config.Config) *AuthService {
	return &AuthService{
		r:   r,
		cfg: cfg,
	}
}

func (s *AuthService) Login(req dto.LoginRequest) (string, *dto.UserResponse, error) {
	user, err := s.r.FindByUsername(req.Username)

	if err != nil {
		return "", nil, pkg.NewApiError(fiber.ErrNotFound.Code, "email or password is wrong", nil)
	}

	if ok := pkg.ComparePassword(user.Password, req.Password); !ok {
		return "", nil, pkg.NewApiError(fiber.StatusNotFound, "email or password is wrong", nil)
	}

	userClaims := &dto.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
	}

	key := s.cfg.JwtSecret
	exp := time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"sub": userClaims.ID, "user": userClaims, "exp": exp})
	jwt, _ := token.SignedString([]byte(key))

	return jwt, userClaims, nil
}

func (s *AuthService) Register(req dto.RegisterRequest) (*model.User, error) {
	req.Password = pkg.HashPassword(req.Password)
	_, err := s.r.FindByUsername(req.Username)

	if err == nil {
		return nil, pkg.NewApiError(fiber.ErrBadRequest.Code, "Username sudah ada", nil)
	}

	user, _ := s.r.Create(req)

	return user, nil
}
