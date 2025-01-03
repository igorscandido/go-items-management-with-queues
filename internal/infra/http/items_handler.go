package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/igorscandido/go-items-management-with-queues/internal/ports"
)

type ItemsHandler struct {
	createItemUsecase ports.CreateItemUsecase
}

func NewItemsHandler(createItemUsecase ports.CreateItemUsecase) *ItemsHandler {
	return &ItemsHandler{
		createItemUsecase,
	}
}

func (h *ItemsHandler) CreateItem(c *gin.Context) {
	var dto CreateItemDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item := MapCreateItemDTOToDomainItem(&dto)
	if err := h.createItemUsecase.PublishItemToQueue(&item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
