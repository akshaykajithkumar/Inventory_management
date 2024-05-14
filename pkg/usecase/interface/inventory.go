package interfaces

import (
	"main/pkg/domain"
	"main/pkg/utils/models"
)

type InventoryUseCase interface {
	AddInventory(inventory models.Inventory) (models.InventoryResponse, error)
	UpdateInventory(invID int, invData models.UpdateInventory) (models.Inventory, error)
	DeleteInventory(id string) error
	GetInventoryByID(inventoryID string) (domain.Inventory, error)
	SearchProducts(key string, page, limit int, sortBY string) ([]domain.Inventory, error)
}
