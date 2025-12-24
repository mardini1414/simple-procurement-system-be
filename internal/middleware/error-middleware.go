package middleware

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/mardini1414/simple-procurement-system-be/pkg"
)

func ErrorMidleware(c *fiber.Ctx, err error) error {

	code := fiber.StatusInternalServerError

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	var e *pkg.ApiError

	if errors.As(err, &e) {
		code = e.Code
		return c.Status(code).JSON(pkg.NewErrorResponse(e.Message, e.Detail))
	}

	var fiberErr *fiber.Error

	if errors.As(err, &fiberErr) {
		code = fiberErr.Code
		return c.Status(code).JSON(pkg.NewErrorResponse(fiberErr.Message, fiberErr.Error()))
	}

	if err != nil {
		return c.Status(code).JSON(pkg.NewErrorResponse("Internal Server Error", err.Error()))
	}

	return nil
}
