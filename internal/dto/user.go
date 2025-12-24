package dto

import "github.com/google/uuid"

type (
	UserResponse struct {
		ID       uuid.UUID `json:"id"`
		Username string    `json:"username"`
		Role     string    `json:"role"`
	}
)
