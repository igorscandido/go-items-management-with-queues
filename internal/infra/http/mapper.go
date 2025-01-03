package http

import "github.com/igorscandido/go-items-management-with-queues/internal/domain"

func MapCreateItemDTOToDomainItem(dto *CreateItemDTO) domain.Item {
	return domain.Item{
		Name:        dto.Name,
		Description: dto.Description,
		Price:       dto.Price,
		Stock:       dto.Stock,
		Status:      dto.Status,
	}
}
