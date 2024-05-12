package interfaces

import (
	"main/pkg/domain"
	"main/pkg/utils/models"
)

type InventoryRepository interface {
	AddInventory(inventory models.Inventory) (models.InventoryResponse, error)

	CheckInventory(pid int) (bool, error)
	UpdateInventory(pid int, invData models.UpdateInventory) (models.Inventory, error)
	DeleteInventory(id string) error
	CheckStock(inventory_id int) (int, error)
	CheckPrice(inventory_id int) (float64, error)
	GetInventoryByID(inventoryID string) (domain.Inventory, error)
	SearchProducts(key string, page, limit int, sortBY string) ([]domain.Inventory, error)
}
