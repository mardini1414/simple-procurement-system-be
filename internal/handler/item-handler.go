package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mardini1414/simple-procurement-system-be/internal/dto"
	"github.com/mardini1414/simple-procurement-system-be/internal/service"
	"github.com/mardini1414/simple-procurement-system-be/pkg"
)

type ItemHandler struct {
	s *service.ItemService
}

func NewItemHandler(s *service.ItemService) *ItemHandler {
	return &ItemHandler{
		s: s,
	}
}

func (h *ItemHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateOrUpdateItemRequest

	if err := c.BodyParser(&req); err != nil {
		return pkg.NewApiError(fiber.StatusBadRequest, "Invalid request body", nil)
	}

	if errs := pkg.ValidateStruct(req); len(errs) != 0 {
		return pkg.NewApiError(fiber.StatusBadRequest, "Invalid request body", errs)
	}

	if _, err := h.s.Create(req); err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(pkg.NewSuccessResponse("Item berhasil dibuat", nil))
}

func (h *ItemHandler) GetAll(c *fiber.Ctx) error {
	Items, err := h.s.FindAll()

	if err != nil {
		return err
	}

	return c.JSON(pkg.NewSuccessResponse("Berhasil mengambil Item", Items))
}

func (h *ItemHandler) Update(c *fiber.Ctx) error {
	var req dto.CreateOrUpdateItemRequest
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

	return c.JSON(pkg.NewSuccessResponse("Item berhasil diupdate", nil))
}

func (h *ItemHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return pkg.NewApiError(fiber.StatusBadRequest, "id must uuid", nil)
	}

	if _, err = h.s.Delete(id); err != nil {
		return err
	}

	return c.JSON(pkg.NewSuccessResponse("Item berhasil dihapus", nil))
}
