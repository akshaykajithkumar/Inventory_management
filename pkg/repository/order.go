package repository

import (
	"errors"
	"main/pkg/domain"
	interfaces "main/pkg/repository/interface"
	"time"

	"gorm.io/gorm"
)

type orderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) interfaces.OrderRepository {
	return &orderRepository{
		DB: db,
	}
}

func (a *orderRepository) GetOrders(userID, page, limit int) ([]domain.Order, error) {
	var orders []domain.Order
	offset := (page - 1) * limit

	query := `
        SELECT * FROM orders
        WHERE user_id = ? 
        ORDER BY ordered_at DESC
        LIMIT ? OFFSET ?
    `

	rows, err := a.DB.Raw(query, userID, limit, offset).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order domain.Order
		err := rows.Scan(&order.ID, &order.UserID, &order.InventoryID, &order.Quantity, &order.Price, &order.OrderedAt, &order.OrderStatus)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (o *orderRepository) GetOrdersInRange(startDate, endDate time.Time) ([]domain.Order, error) {
	var orders []domain.Order

	if err := o.DB.Raw("SELECT * FROM orders WHERE  ordered_at BETWEEN ? AND ?", startDate, endDate).Scan(&orders).Error; err != nil {
		return []domain.Order{}, err
	}
	return orders, nil
}

func (o *orderRepository) GetProductsQuantity() ([]domain.ProductReport, error) {

	var products []domain.ProductReport

	if err := o.DB.Raw("select inventory_id,quantity from order_items").Scan(&products).Error; err != nil {
		return []domain.ProductReport{}, err
	}
	return products, nil
}

func (o *orderRepository) GetProductNameFromID(id int) (string, error) {
	var product string

	if err := o.DB.Raw("SELECT product_name FROM inventories WHERE id=?", id).Scan(&product).Error; err != nil {
		return "", err
	}
	return product, nil
}

func (o *orderRepository) CancelOrder(orderid int) error {

	if err := o.DB.Exec("update orders set order_status='CANCELED' where id=?", orderid).Error; err != nil {
		return err
	}

	return nil
}

func (o *orderRepository) EditOrderStatus(status string, id int) error {

	if err := o.DB.Exec("update orders set order_status=? where id=?", status, id).Error; err != nil {
		return err
	}

	return nil
}

func (o *orderRepository) MarkAsPaid(orderID int) error {

	if err := o.DB.Exec("update orders set payment_status='PAID' where id=?", orderID).Error; err != nil {
		return err
	}

	return nil
}

func (o *orderRepository) CheckOrder(orderID string, userID int) error {

	var count int
	err := o.DB.Raw("select count(*) from orders where order_id = ?", orderID).Scan(&count).Error
	if err != nil {
		return err
	}
	if count < 0 {
		return errors.New("no such order exist")
	}
	var checkUser int
	err = o.DB.Raw("select user_id from orders where order_id = ?", orderID).Scan(&checkUser).Error
	if err != nil {
		return err
	}

	if userID != checkUser {
		return errors.New("the order is not did by this user")
	}

	return nil
}

func (o *orderRepository) FindAmountFromOrderID(id int) (float64, error) {

	var amount float64
	err := o.DB.Raw("select price from orders where id = ?", id).Scan(&amount).Error
	if err != nil {
		return 0, err
	}

	return amount, nil
}

func (o *orderRepository) FindUserIdFromOrderID(id int) (int, error) {

	var user_id int
	err := o.DB.Raw("select user_id from orders where id = ?", id).Scan(&user_id).Error
	if err != nil {
		return 0, err
	}

	return user_id, nil
}

func (o *orderRepository) GetProductByID(productID int) (bool, error) {
	var count int64
	err := o.DB.Raw("SELECT COUNT(*) FROM inventories WHERE id = ?", productID).Scan(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
func (o *orderRepository) CheckProductAvailability(productID int, quantity int) (bool, error) {
	var availableStock int
	err := o.DB.Raw("SELECT stock FROM inventories WHERE id = ?", productID).Scan(&availableStock).Error
	if err != nil {
		return false, err
	}

	return availableStock >= quantity, nil
}
func (o *orderRepository) PlaceOrder(userID, productID, quantity int, price float64) error {

	totalPrice := float64(quantity) * price

	err := o.DB.Exec("INSERT INTO orders (user_id, inventory_id, quantity, price, ordered_at, order_status) VALUES (?, ?, ?, ?, ?, ?)",
		userID, productID, quantity, totalPrice, time.Now(), "PENDING").Error
	if err != nil {
		return err
	}
	return nil
}
func (o *orderRepository) FindAmountFromProductID(productID int) (float64, error) {
	var price float64
	err := o.DB.Raw("SELECT price FROM inventories WHERE id = ?", productID).Scan(&price).Error
	if err != nil {
		return 0, err
	}

	return price, nil
}
