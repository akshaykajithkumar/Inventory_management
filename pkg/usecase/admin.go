package usecase

import (
	"errors"
	"main/pkg/domain"
	"main/pkg/helper"
	interfaces "main/pkg/repository/interface"
	services "main/pkg/usecase/interface"
	"main/pkg/utils/models"

	"golang.org/x/crypto/bcrypt"
)

type adminUseCase struct {
	adminRepository interfaces.AdminRepository
}

func NewAdminUseCase(repo interfaces.AdminRepository) services.AdminUseCase {
	return &adminUseCase{
		adminRepository: repo,
	}
}

func (ad *adminUseCase) LoginHandler(adminDetails models.AdminLogin) (models.TokenAdmin, error) {

	// getting details of the admin based on the email provided
	adminCompareDetails, err := ad.adminRepository.LoginHandler(adminDetails)
	if err != nil {
		return models.TokenAdmin{}, err
	}

	// compare password from database and that provided from admins
	err = bcrypt.CompareHashAndPassword([]byte(adminCompareDetails.Password), []byte(adminDetails.Password))
	if err != nil {
		return models.TokenAdmin{}, err
	}

	accessTokenString, refreshTokenString, err := helper.GenerateTokensAdmin(adminCompareDetails)

	if err != nil {
		return models.TokenAdmin{}, err
	}

	return models.TokenAdmin{
		Username:     adminCompareDetails.Username,
		RefreshToken: refreshTokenString,
		AccessToken:  accessTokenString,
	}, nil

}

func (ad *adminUseCase) GetUsers(page int, limit int) ([]models.UserDetailsAtAdmin, error) {

	userDetails, err := ad.adminRepository.GetUsers(page, limit)
	if err != nil {
		return []models.UserDetailsAtAdmin{}, err
	}

	return userDetails, nil

}

// stats
func (a *adminUseCase) UserStats() (models.UserStats, error) {
	//finding total users
	totalUsers, err := a.adminRepository.CountUsers()
	if err != nil {
		return models.UserStats{}, errors.New("could not find users data")
	}
	//finding total revenue
	totalRevenue, err := a.adminRepository.GetTotalRevenue()
	if err != nil {
		return models.UserStats{}, errors.New("could not find revenue data")
	}
	//finding total orders
	totalOrders, err := a.adminRepository.CountOrders()
	if err != nil {
		return models.UserStats{}, errors.New("could not find orders data")
	}
	return models.UserStats{
		TotalUsers:            totalUsers,
		AverageRevenuePerUser: totalRevenue / float32(totalUsers),
		AverageOrderPerUSer:   float32(totalOrders) / float32(totalUsers),
	}, nil
}

func (a *adminUseCase) OrderStats() (models.OrderStats, error) {
	//finding total orders
	totalOrders, err := a.adminRepository.CountOrders()
	if err != nil {
		return models.OrderStats{}, errors.New("could not find orders data")
	}
	//finding total revenue
	totalRevenue, err := a.adminRepository.GetTotalRevenue()
	if err != nil {
		return models.OrderStats{}, errors.New("could not find revenue data")
	}
	//finding todays revenue
	totalRevenueToday, err := a.adminRepository.GetTotalRevenueToday()
	if err != nil {
		return models.OrderStats{}, errors.New("could not find revenue data")
	}
	return models.OrderStats{
		TotalOrders:       totalOrders,
		TotalRevenue:      totalRevenue,
		TodaysRevenue:     totalRevenueToday,
		AverageOrderValue: totalRevenue / float32(totalOrders),
	}, nil
}

func (a *adminUseCase) InventoryStats() (models.InventoryStats, error) {
	//finding most sold product of all time
	mostSoldProduct, err := a.adminRepository.GetMostSoldProduct()
	if err != nil {
		return models.InventoryStats{}, errors.New("could not find product data")
	}
	//finding todays most sold product
	trendingProduct, err := a.adminRepository.GetTrendingProduct()
	if err != nil {
		return models.InventoryStats{}, errors.New("could not find product data")
	}
	//finding total number of products
	totalProducts, err := a.adminRepository.CountProducts()
	if err != nil {
		return models.InventoryStats{}, errors.New("could not find product data")
	}
	//finding total stock
	totalStock, err := a.adminRepository.CountStock()
	if err != nil {
		return models.InventoryStats{}, errors.New("could not find product data")
	}
	totalPrice, err := a.adminRepository.GetTotalPrice()
	if err != nil {
		return models.InventoryStats{}, errors.New("could not find product data")
	}

	return models.InventoryStats{
		MostSoldProduct: mostSoldProduct,
		TrendingProduct: trendingProduct,
		TotalProducts:   totalProducts,
		TotalStock:      totalStock,
		AveragePrice:    totalPrice / float32(totalProducts),
	}, nil
}
func (i *adminUseCase) AdminOrders(page, limit int, status string) ([]domain.Order, error) {

	if status != "PENDING" && status != "SHIPPED" && status != "CANCELED" && status != "RETURNED" && status != "DELIVERED" {
		return []domain.Order{}, errors.New("invalid status type")

	}
	orders, err := i.adminRepository.AdminOrders(page, limit, status)
	if err != nil {
		return []domain.Order{}, err
	}

	return orders, nil

}

func (i *adminUseCase) GetOrder(id int) (domain.Order, error) {

	orders, err := i.adminRepository.AdminGetOrder(id)
	if err != nil {
		return domain.Order{}, err
	}

	return orders, nil

}
func (i *adminUseCase) ChangeOrderStatus(orderID, status string) error {
	// Check if the status is valid
	if status != "PENDING" && status != "SHIPPED" && status != "CANCELED" && status != "RETURNED" && status != "DELIVERED" {
		return errors.New("invalid status type")
	}

	// Call repository method to change the order status
	err := i.adminRepository.ChangeOrderStatus(orderID, status)
	if err != nil {
		return err
	}

	return nil
}
