package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/mardini1414/simple-procurement-system-be/internal/config"
	"github.com/mardini1414/simple-procurement-system-be/internal/dto"
	"github.com/mardini1414/simple-procurement-system-be/pkg"
)

type JWTClaims struct {
	Sub  uuid.UUID `json:"sub"`
	User dto.UserResponse
	jwt.RegisteredClaims
}

func AuthMiddleware(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return pkg.NewApiError(fiber.StatusUnauthorized, "missing authorization header", nil)
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return pkg.NewApiError(fiber.StatusUnauthorized, "invalid authorization format", nil)
		}

		tokenStr := parts[1]

		token, err := jwt.ParseWithClaims(
			tokenStr,
			&JWTClaims{},
			func(token *jwt.Token) (any, error) {
				return []byte(cfg.JwtSecret), nil
			},
		)

		if err != nil || !token.Valid {
			return pkg.NewApiError(fiber.StatusUnauthorized, "invalid or expired access token", err.Error())
		}

		claims := token.Claims.(*JWTClaims)

		c.Locals("user_id", claims.Sub)
		c.Locals("role", claims.User.Role)

		return c.Next()
	}
}
