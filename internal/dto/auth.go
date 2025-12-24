package dto

type (
	LoginRequest struct {
		Username string `json:"username" validate:"required,min=3,max=50"`
		Password string `json:"password" validate:"required,min=6,max=225"`
	}

	RegisterRequest struct {
		Username string `json:"username" validate:"required,min=3,max=50"`
		Password string `json:"password" validate:"required,min=6,max=225"`
	}
)
