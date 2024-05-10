package interfaces

import (
	"main/pkg/utils/models"
)

type InventoryUseCase interface {
	AddInventory(inventory models.Inventory) (models.InventoryResponse, error)
	UpdateInventory(invID int, invData models.UpdateInventory) (models.Inventory, error)
	DeleteInventory(id string) error

	ListProducts(page int, limit int) ([]models.InventoryList, error)
	SearchProducts(key string, page, limit int, sortBY string) ([]models.InventoryList, error)
}
