package usecase

import (
	"errors"
	domain "main/pkg/domain"

	interfaces "main/pkg/repository/interface"
	services "main/pkg/usecase/interface"
)

type orderUseCase struct {
	orderRepository interfaces.OrderRepository
	userUseCase     services.UserUseCase
}

func NewOrderUseCase(repo interfaces.OrderRepository, userUseCase services.UserUseCase) *orderUseCase {
	return &orderUseCase{
		orderRepository: repo,
		userUseCase:     userUseCase,
	}
}
func (i *orderUseCase) PlaceOrder(userID, productID, quantity int) error {
	// Check if the product exists
	_, err := i.orderRepository.GetProductByID(productID)
	if err != nil {
		return err
	}

	available, err := i.orderRepository.CheckProductAvailability(productID, quantity)
	if err != nil {
		return err
	}
	if !available {
		return errors.New("insufficient quantity available")
	}

	// Get the price of the product
	price, err := i.orderRepository.FindAmountFromProductID(productID)
	if err != nil {
		return errors.New("price not found")
	}

	// Call the repository to place the order
	if err := i.orderRepository.PlaceOrder(userID, productID, quantity, price); err != nil {
		return err
	}

	return nil
}

func (i *orderUseCase) GetOrders(id, page, limit int) ([]domain.Order, error) {

	orders, err := i.orderRepository.GetOrders(id, page, limit)
	if err != nil {
		return []domain.Order{}, err
	}

	return orders, nil

}
