package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mardini1414/simple-procurement-system-be/internal/dto"
	"github.com/mardini1414/simple-procurement-system-be/internal/model"
	"github.com/mardini1414/simple-procurement-system-be/internal/repository"
	"github.com/mardini1414/simple-procurement-system-be/pkg"
)

type SupplierService struct {
	r *repository.SupplierRepository
}

func NewSupplierService(r *repository.SupplierRepository) *SupplierService {
	return &SupplierService{
		r: r,
	}
}

func (s *SupplierService) Create(req dto.CreateOrUpdateSupplierRequest) (*model.Supplier, error) {
	if _, err := s.r.FindByEmail(req.Email); err == nil {
		return nil, pkg.NewApiError(fiber.StatusBadRequest, "Email sudah ada", nil)
	}

	return s.r.Create(req)
}

func (s *SupplierService) Update(req dto.CreateOrUpdateSupplierRequest, id uuid.UUID) (*model.Supplier, error) {
	supplier, err := s.r.FindByEmail(req.Email)

	if err == nil && supplier.ID != id {
		return nil, pkg.NewApiError(fiber.StatusBadRequest, "Email sudah ada", nil)
	}

	if _, err := s.r.FindById(id); err != nil {
		return nil, pkg.NewApiError(fiber.StatusNotFound, "Supplier tidak di temukan", nil)
	}

	return s.r.Update(req, id)
}

func (s *SupplierService) Delete(id uuid.UUID) (*model.Supplier, error) {
	if _, err := s.r.FindById(id); err != nil {
		return nil, pkg.NewApiError(fiber.StatusNotFound, "Supplier tidak di temukan", nil)
	}

	return s.r.Delete(id)
}

func (s *SupplierService) FindAll() (*[]model.Supplier, error) {
	return s.r.FindAll()
}
