package repository

import (
	"fmt"
	"main/pkg/domain"
	interfaces "main/pkg/repository/interface"
	"main/pkg/utils/models"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type adminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRepository {
	return &adminRepository{
		DB: DB,
	}
}

func (ad *adminRepository) LoginHandler(adminDetails models.AdminLogin) (domain.Admin, error) {

	var adminCompareDetails domain.Admin
	if err := ad.DB.Raw("select * from admins where email = ? ", adminDetails.Email).Scan(&adminCompareDetails).Error; err != nil {
		return domain.Admin{}, err
	}
	return adminCompareDetails, nil
}

func (ad *adminRepository) GetUserByID(id string) (domain.User, error) {

	user_id, err := strconv.Atoi(id)
	if err != nil {
		return domain.User{}, err
	}

	query := fmt.Sprintf("select * from users where id = '%d'", user_id)
	var userDetails domain.User

	if err := ad.DB.Raw(query).Scan(&userDetails).Error; err != nil {
		return domain.User{}, err
	}

	return userDetails, nil
}

func (ad *adminRepository) GetUsers(page int, limit int) ([]models.UserDetailsAtAdmin, error) {
	// pagination purpose -
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	offset := (page - 1) * limit
	var userDetails []models.UserDetailsAtAdmin

	if err := ad.DB.Raw("select id,name,email,phone,permission from users limit ? offset ?", limit, offset).Scan(&userDetails).Error; err != nil {
		return []models.UserDetailsAtAdmin{}, err
	}

	return userDetails, nil

}

//stats

func (a *adminRepository) CountUsers() (int, error) {
	var count int64
	query := "SELECT COUNT(*) FROM users"
	if err := a.DB.Raw(query).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func (a *adminRepository) GetTotalRevenue() (float32, error) {
	var totalRevenue float32
	if err := a.DB.Model(&domain.Order{}).Select("SUM(quantity * price)").Scan(&totalRevenue).Error; err != nil {
		return 0, err
	}
	return totalRevenue, nil
}

func (a *adminRepository) GetTotalRevenueToday() (float32, error) {
	var todaysRevenue float32
	if err := a.DB.Model(&domain.Order{}).Select("SUM(quantity * price)").Where("DATE(ordered_at) = CURRENT_DATE").Scan(&todaysRevenue).Error; err != nil {
		return 0, err
	}
	return todaysRevenue, nil
}

func (a *adminRepository) CountOrders() (int, error) {
	var count int64
	if err := a.DB.Model(&domain.Order{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func (a *adminRepository) CountProducts() (int, error) {
	var count int64
	query := "SELECT COUNT(*) FROM inventories"
	if err := a.DB.Raw(query).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func (a *adminRepository) CountStock() (int, error) {
	var totalStock int
	query := "SELECT SUM(stock) FROM inventories"
	if err := a.DB.Raw(query).Scan(&totalStock).Error; err != nil {
		return 0, err
	}
	return totalStock, nil
}

func (a *adminRepository) GetTotalPrice() (float32, error) {
	var totalPrice float32
	query := "SELECT SUM(price) FROM inventories"
	if err := a.DB.Raw(query).Scan(&totalPrice).Error; err != nil {
		return 0, err
	}
	return totalPrice, nil
}

func (a *adminRepository) GetMostSoldProduct() (string, error) {
	var mostSoldProductID uint
	var mostSoldProductName string

	query := `
        SELECT inventory_id
        FROM orders
        GROUP BY inventory_id
        ORDER BY SUM(quantity) DESC
        LIMIT 1
    `
	if err := a.DB.Raw(query).Scan(&mostSoldProductID).Error; err != nil {
		return "", err
	}

	// Assuming you have a Product table associated with the Inventory table.
	if err := a.DB.Model(&domain.Inventory{}).Where("id = ?", mostSoldProductID).Select("product_name").Scan(&mostSoldProductName).Error; err != nil {
		return "", err
	}

	return mostSoldProductName, nil
}

func (a *adminRepository) GetTrendingProduct() (string, error) {
	var trendingProductID uint
	var trendingProductName string

	// Get the start and end of today.
	todayStart := time.Now().Truncate(24 * time.Hour)
	todayEnd := todayStart.Add(24 * time.Hour)

	query := `
        SELECT inventory_id
        FROM orders
        WHERE ordered_at >= ? AND ordered_at < ?
        GROUP BY inventory_id
        ORDER BY SUM(quantity) DESC
        LIMIT 1
    `
	if err := a.DB.Raw(query, todayStart, todayEnd).Scan(&trendingProductID).Error; err != nil {
		return "", err
	}

	// Assuming  Product table associated with the Inventory table.
	if err := a.DB.Model(&domain.Inventory{}).Where("id = ?", trendingProductID).Select("product_name").Scan(&trendingProductName).Error; err != nil {
		return "", err
	}

	return trendingProductName, nil
}
func (o *adminRepository) AdminGetOrder(id int) (domain.Order, error) {
	var order domain.Order

	// Execute raw SQL query to retrieve the order by ID
	if err := o.DB.Raw(`
        SELECT
            id,
            user_id,
            payment_id,
            price,
            ordered_at,
            order_status
        FROM
            orders
        WHERE
            id = ?
    `, id).Scan(&order).Error; err != nil {
		return domain.Order{}, err
	}

	return order, nil
}

// ChangeOrderStatus changes the status of a specific order
func (a *adminRepository) ChangeOrderStatus(orderID, status string) error {
	// Prepare the SQL query
	query := "UPDATE orders SET order_status = ? WHERE id = ?"

	// Execute the SQL query
	result := a.DB.Exec(query, status, orderID)
	if result.Error != nil {
		return result.Error // Return error if there's any issue with the query execution
	}

	return nil
}

func (i *adminRepository) AdminOrders(page, limit int, status string) ([]domain.Order, error) {
	// Execute raw SQL query to retrieve orders
	var orders []domain.Order
	query := "SELECT * FROM orders WHERE order_status = ? LIMIT ? OFFSET ?"
	if err := i.DB.Raw(query, status, limit, (page-1)*limit).Scan(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}
