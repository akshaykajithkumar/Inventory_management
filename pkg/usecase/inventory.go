package usecase

import (
	"errors"
	"main/pkg/domain"
	interfaces "main/pkg/repository/interface"
	services "main/pkg/usecase/interface"
	"main/pkg/utils/models"
)

type inventoryUseCase struct {
	repository interfaces.InventoryRepository
}

func NewInventoryUseCase(repo interfaces.InventoryRepository) services.InventoryUseCase {
	return &inventoryUseCase{
		repository: repo,
	}
}

func (i *inventoryUseCase) AddInventory(inventory models.Inventory) (models.InventoryResponse, error) {

	//send the url and save it in database
	InventoryResponse, err := i.repository.AddInventory(inventory)
	if err != nil {
		return models.InventoryResponse{}, err
	}

	return InventoryResponse, nil

}

func (i *inventoryUseCase) UpdateInventory(invID int, invData models.UpdateInventory) (models.Inventory, error) {

	result, err := i.repository.CheckInventory(invID)
	if err != nil {

		return models.Inventory{}, err
	}

	if !result {
		return models.Inventory{}, errors.New("there is no inventory as you mentioned")
	}

	newinv, err := i.repository.UpdateInventory(invID, invData)
	if err != nil {
		return models.Inventory{}, err
	}

	return newinv, err
}

func (i *inventoryUseCase) DeleteInventory(inventoryID string) error {

	err := i.repository.DeleteInventory(inventoryID)
	if err != nil {
		return err
	}
	return nil

}

func (i *inventoryUseCase) SearchProducts(key string, page, limit int, sortBY string) ([]domain.Inventory, error) {

	productDetails, err := i.repository.SearchProducts(key, page, limit, sortBY)
	if err != nil {
		return []domain.Inventory{}, err
	}

	return productDetails, nil

}

func (i *inventoryUseCase) GetInventoryByID(inventoryID string) (domain.Inventory, error) {
	// Call the repository function to fetch the inventory details by ID
	inventory, err := i.repository.GetInventoryByID(inventoryID)
	if err != nil {
		return domain.Inventory{}, err
	}

	// Return the fetched inventory details
	return inventory, nil
}
