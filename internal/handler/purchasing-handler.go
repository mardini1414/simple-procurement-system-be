package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mardini1414/simple-procurement-system-be/internal/dto"
	"github.com/mardini1414/simple-procurement-system-be/internal/service"
	"github.com/mardini1414/simple-procurement-system-be/pkg"
)

type PurchasingHandler struct {
	s *service.PurchasingService
}

func NewPurchasingHandler(s *service.PurchasingService) *PurchasingHandler {
	return &PurchasingHandler{
		s: s,
	}
}

func (h *PurchasingHandler) Create(c *fiber.Ctx) error {
	var req dto.CreatePurchasingRequest

	if err := c.BodyParser(&req); err != nil {
		return pkg.NewApiError(fiber.StatusBadRequest, "Invalid request body", nil)
	}

	if errs := pkg.ValidateStruct(req); len(errs) != 0 {
		return pkg.NewApiError(fiber.StatusBadRequest, "Invalid request body", errs)
	}

	err := h.s.Create(req, c)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(pkg.NewSuccessResponse("Success membuat purchase", nil))
}

func (h *PurchasingHandler) FindAll(c *fiber.Ctx) error {
	purchasings, err := h.s.FindAll()

	if err != nil {
		return err
	}

	return c.JSON(pkg.NewSuccessResponse("Success mengambil purchasings", purchasings))
}

func (h *PurchasingHandler) FindById(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return pkg.NewApiError(fiber.StatusBadRequest, "id must uuid", nil)
	}

	purchasing, err := h.s.FindById(id)

	if err != nil {
		return err
	}

	return c.JSON(pkg.NewSuccessResponse("Success mengambil detail purchasing", purchasing))
}
