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

func (o *orderRepository) GetOrders(id, page, limit int) ([]domain.Order, error) {

	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	offset := (page - 1) * limit
	var orders []domain.Order

	//		if err := o.DB.Raw("select * from orders where user_id=? limit ? offset ?", id, limit, offset).Scan(&orders).Error; err != nil {
	//			return []domain.Order{}, err
	//		}
	//		return orders, nil
	//	}
	if err := o.DB.Raw(`
    SELECT
        o.id,
        o.user_id,
        o.address_id,
        pm.payment_method AS payment_method,
        o.payment_id,
        o.price,
        o.ordered_at,
        o.order_status,
        o.payment_status
    FROM
        orders o
    JOIN
        payment_methods pm
    ON
        o.payment_method_id = pm.id
    WHERE
        o.user_id = ?
    LIMIT ? OFFSET ?;
`, id, limit, offset).Scan(&orders).Error; err != nil {
		return []domain.Order{}, err
	}
	return orders, nil

}

func (o *orderRepository) GetOrdersInRange(startDate, endDate time.Time) ([]domain.Order, error) {
	var orders []domain.Order

	// Execute the query to get orders within the specified time range
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

// func (o *orderRepository) OrderItems(userid int, order models.Order, total float64) (int, error) {

// 	var id int
// 	query := `
//     INSERT INTO orders (user_id,address_id ,price,payment_method_id,ordered_at)
//     VALUES (?, ?, ?, ?, ?)
//     RETURNING id
//     `
// 	if err := o.DB.Raw(query, userid, order.AddressID, total, order.PaymentID, time.Now()).Scan(&id).Error; err != nil {
// 		return 0, err
// 	}

// 	return id, nil
// }

// func (o *orderRepository) AddOrderProducts(order_id int, cart []models.GetCart) error {

// 	query := `
//     INSERT INTO order_items (order_id,inventory_id,quantity,total_price)
//     VALUES (?, ?, ?, ?)
//     `

// 	for _, v := range cart {
// 		var inv int
// 		if err := o.DB.Raw("select id from inventories where product_name=?", v.ProductName).Scan(&inv).Error; err != nil {
// 			return err
// 		}

// 		if err := o.DB.Exec(query, order_id, inv, v.Quantity, v.Total).Error; err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

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

func (o *orderRepository) AdminOrders(page, limit int, status string) ([]domain.OrderDetails, error) {
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	offset := (page - 1) * limit
	var orders []domain.OrderDetails
	if err := o.DB.Raw("SELECT orders.id AS order_id, users.name AS username, CONCAT(addresses.house_name, ' ', addresses.street, ' ', addresses.city) AS address, payment_methods.payment_method AS payment_method, orders.price As total FROM orders JOIN users ON users.id = orders.user_id JOIN addresses ON orders.address_id = addresses.id JOIN payment_methods ON orders.payment_method_id=payment_methods.id WHERE order_status = ? limit ? offset ?", status, limit, offset).Scan(&orders).Error; err != nil {
		return []domain.OrderDetails{}, err
	}

	return orders, nil
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

func (r *orderRepository) CreateOrder(order *domain.Order) error {
	// Implement code to insert order into the database using raw SQL query
	query := "INSERT INTO orders (user_id, product_id, quantity, status) VALUES (?, ?, ?, ?)"
	result := r.DB.Exec(query, order.UserID, order.ProductID, order.Quantity, order.Status)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("failed to insert order")
	}
	return nil
}

func (r *orderRepository) UpdateProduct(product *domain.Product) error {
	// Implement code to update product quantity in the database using raw SQL query
	query := "UPDATE products SET quantity = ? WHERE id = ?"
	result := r.DB.Exec(query, product.Quantity, product.ID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("failed to update product")
	}
	return nil
}

// // Implement a database trigger for updating inventory when orders are placed
// // This example assumes you're using PostgreSQL
// // Replace `your_product_table` with the actual name of your product table
// // This trigger assumes that the `orders` table has `product_id` and `quantity` columns
// // and the `products` table has `id` and `quantity` columns
// // You need to execute this SQL query in your database
// const updateInventoryTrigger = `
// CREATE OR REPLACE FUNCTION update_inventory()
// RETURNS TRIGGER AS $$
// BEGIN
//     UPDATE your_product_table
//     SET quantity = quantity - NEW.quantity
//     WHERE id = NEW.product_id;
//     RETURN NEW;
// END;
// $$ LANGUAGE plpgsql;

// CREATE TRIGGER order_placed_trigger
// AFTER INSERT ON orders
// FOR EACH ROW
// EXECUTE FUNCTION update_inventory();
// `

// CREATE OR REPLACE FUNCTION adjust_product_pricing()
// RETURNS TRIGGER AS $$
// BEGIN
//     -- Check if the order exceeds 100 in 1 hour
//     IF (SELECT COUNT(*) FROM orders WHERE ordered_at > current_timestamp - interval '1 hour') > 100 THEN
//         UPDATE products
//         SET price = price * 1.2  -- Increase price by 20%
//         WHERE id IN (SELECT product_id FROM orders WHERE ordered_at > current_timestamp - interval '1 hour');
//     END IF;

//     -- Check if stock is less than 100
//     IF (SELECT COUNT(*) FROM products WHERE id = NEW.product_id AND quantity < 100) > 0 THEN
//         UPDATE products
//         SET price = price * 1.2  -- Increase price by 20%
//         WHERE id = NEW.product_id;
//     END IF;

//     RETURN NEW;
// END;
// $$ LANGUAGE plpgsql;

// CREATE TRIGGER adjust_product_pricing_trigger
// AFTER INSERT ON orders
// FOR EACH ROW
// EXECUTE FUNCTION adjust_product_pricing();
