package http

type CreateItemDTO struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float32 `json:"price" binding:"required"`
	Stock       int     `json:"stock" binding:"required"`
	Status      string  `json:"status" binding:"required"`
}

type UpdateItemDTO struct{}
