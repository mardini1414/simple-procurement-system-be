package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mardini1414/simple-procurement-system-be/internal/dto"
	"github.com/mardini1414/simple-procurement-system-be/internal/model"
	"github.com/mardini1414/simple-procurement-system-be/internal/repository"
	"github.com/mardini1414/simple-procurement-system-be/pkg"
)

type ItemService struct {
	r *repository.ItemRepository
}

func NewItemService(r *repository.ItemRepository) *ItemService {
	return &ItemService{
		r: r,
	}
}

func (s *ItemService) Create(req dto.CreateOrUpdateItemRequest) (*model.Item, error) {
	return s.r.Create(req)
}

func (s *ItemService) Update(req dto.CreateOrUpdateItemRequest, id uuid.UUID) (*model.Item, error) {
	if _, err := s.r.FindById(id); err != nil {
		return nil, pkg.NewApiError(fiber.StatusNotFound, "Item tidak di temukan", nil)
	}

	return s.r.Update(req, id)
}

func (s *ItemService) Delete(id uuid.UUID) (*model.Item, error) {
	if _, err := s.r.FindById(id); err != nil {
		return nil, pkg.NewApiError(fiber.StatusNotFound, "Item tidak di temukan", nil)
	}

	return s.r.Delete(id)
}

func (s *ItemService) FindAll() (*[]model.Item, error) {
	return s.r.FindAll()
}
