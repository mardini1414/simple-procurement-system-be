package dto

type (
	CreateOrUpdateItemRequest struct {
		Name  string `json:"name" validate:"required,min=3,max=50"`
		Stock int    `json:"stock" validate:"required,gte=0"`
		Price int64  `json:"price" validate:"required,gte=0"`
	}
)
