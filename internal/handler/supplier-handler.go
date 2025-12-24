package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mardini1414/simple-procurement-system-be/internal/dto"
	"github.com/mardini1414/simple-procurement-system-be/internal/service"
	"github.com/mardini1414/simple-procurement-system-be/pkg"
)

type SupplierHandler struct {
	s *service.SupplierService
}

func NewSupplierHandler(s *service.SupplierService) *SupplierHandler {
	return &SupplierHandler{
		s: s,
	}
}

func (h *SupplierHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateOrUpdateSupplierRequest

	if err := c.BodyParser(&req); err != nil {
		return pkg.NewApiError(fiber.StatusBadRequest, "Invalid request body", nil)
	}

	if errs := pkg.ValidateStruct(req); len(errs) != 0 {
		return pkg.NewApiError(fiber.StatusBadRequest, "Invalid request body", errs)
	}

	if _, err := h.s.Create(req); err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(pkg.NewSuccessResponse("Supplier berhasil dibuat", nil))
}

func (h *SupplierHandler) GetAll(c *fiber.Ctx) error {
	suppliers, err := h.s.FindAll()

	if err != nil {
		return err
	}

	return c.JSON(pkg.NewSuccessResponse("Berhasil mengambil supplier", suppliers))
}

func (h *SupplierHandler) Update(c *fiber.Ctx) error {
	var req dto.CreateOrUpdateSupplierRequest
	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return pkg.NewApiError(fiber.StatusBadRequest, "id must uuid", nil)
	}

	if err := c.BodyParser(&req); err != nil {
		return pkg.NewApiError(fiber.StatusBadRequest, "Invalid request body", nil)
	}

	if errs := pkg.ValidateStruct(req); len(errs) != 0 {
		return pkg.NewApiError(fiber.StatusBadRequest, "Invalid request body", errs)
	}

	if _, err := h.s.Update(req, id); err != nil {
		return err
	}

	return c.JSON(pkg.NewSuccessResponse("Supplier berhasil diupdate", nil))
}

func (h *SupplierHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return pkg.NewApiError(fiber.StatusBadRequest, "id must uuid", nil)
	}

	if _, err = h.s.Delete(id); err != nil {
		return err
	}

	return c.JSON(pkg.NewSuccessResponse("Supplier berhasil dihapus", nil))
}
