package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mardini1414/simple-procurement-system-be/internal/dto"
	"github.com/mardini1414/simple-procurement-system-be/internal/service"
	"github.com/mardini1414/simple-procurement-system-be/pkg"
)

type AuthHandler struct {
	s *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{
		s: s,
	}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req dto.LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return pkg.NewApiError(fiber.StatusBadRequest, "Login gagal", nil)
	}

	if errs := pkg.ValidateStruct(req); len(errs) != 0 {
		return pkg.NewApiError(fiber.StatusBadRequest, "Login gagal", errs)
	}

	token, user, err := h.s.Login(req)

	if err != nil {
		return err
	}

	return c.JSON(pkg.NewSuccessResponse("Login success", fiber.Map{
		"user":         user,
		"access_token": token,
	}))
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req dto.RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return pkg.NewApiError(fiber.StatusBadRequest, "Register gagal", nil)
	}

	if errs := pkg.ValidateStruct(req); len(errs) != 0 {
		return pkg.NewApiError(fiber.StatusBadRequest, "Register gagal", errs)
	}

	user, err := h.s.Register(req)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(pkg.NewSuccessResponse("Register success", fiber.Map{
		"username": user.Username,
	}))
}
