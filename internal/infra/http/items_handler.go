package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/igorscandido/go-items-management-with-queues/internal/ports"
)

type ItemsHandler struct {
	createItemUsecase ports.InsertItemUsecase
}

func NewItemsHandler(insertItemUsecase ports.InsertItemUsecase) *ItemsHandler {
	return &ItemsHandler{
		insertItemUsecase,
	}
}

func (h *ItemsHandler) InsertItem(c *gin.Context) {
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
