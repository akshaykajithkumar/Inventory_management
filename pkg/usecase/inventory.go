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

// func (i *inventoryUseCase) ShowIndividualProducts(id string) (models.InventoryDetails, error) {

// 	product, err := i.repository.ShowIndividualProducts(id)
// 	if err != nil {
// 		return models.InventoryDetails{}, err
// 	}
// 	productID, err := strconv.Atoi(id)
// 	if err != nil {
// 		return models.InventoryDetails{}, err
// 	}
// 	var AdditionalImages []models.ImageInfo
// 	AdditionalImages, err = i.repository.GetImagesFromInventoryID(productID)
// 	if err != nil {
// 		return models.InventoryDetails{}, err
// 	}
// 	InvDetails := models.InventoryDetails{Inventory: product, AdditionalImages: AdditionalImages}

// 	return InvDetails, nil

// }

func (i *inventoryUseCase) ListProducts(page int, limit int) ([]models.InventoryList, error) {

	productDetails, err := i.repository.ListProducts(page, limit)
	if err != nil {
		return []models.InventoryList{}, err
	}
	return productDetails, nil

}

func (i *inventoryUseCase) SearchProducts(key string, page, limit int, sortBY string) ([]models.InventoryList, error) {

	productDetails, err := i.repository.SearchProducts(key, page, limit, sortBY)
	if err != nil {
		return []models.InventoryList{}, err
	}

	return productDetails, nil

}

func (i *orderUseCase) PlaceOrder(userID, productID, quantity int) error {
	// Check if the product exists
	product, err := i.orderRepository.GetProductByID(productID)
	if err != nil {
		return err
	}
	if product == nil {
		return errors.New("product not found")
	}

	// Check if the quantity is available
	if quantity > product.Quantity {
		return errors.New("insufficient quantity available")
	}

	// Create the order
	order := domain.Order{
		UserID:    userID,
		ProductID: productID,
		Quantity:  quantity,
		Status:    "PENDING", // Assuming the initial status is pending
	}

	// Call the repository to save the order
	err = i.orderRepository.CreateOrder(&order)
	if err != nil {
		return err
	}

	// Update the product quantity
	product.Quantity -= quantity
	err = i.orderRepository.UpdateProduct(product)
	if err != nil {

		return err
	}

	return nil
}
