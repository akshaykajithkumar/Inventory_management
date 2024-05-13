// package db

// import (
// 	"errors"

// 	"gorm.io/gorm"
// )

// // SetUpDBTriggers sets up database triggers
// func SetUpDBTriggers(db *gorm.DB) error {
// 	// Create trigger to update inventory stock after placing an order
// 	if err := db.Exec(placeOrderTriggerSQL).Error; err != nil {
// 		return errors.New("failed to create place_order trigger")
// 	}

// 	// Create trigger to adjust product price based on demand and availability
// 	if err := db.Exec(adjustPriceTriggerSQL).Error; err != nil {
// 		return errors.New("failed to create adjust_price trigger")
// 	}

// 	return nil
// }

// var (
// 	// SQL trigger to update inventory stock after placing an order
// 	placeOrderTriggerSQL = `CREATE TRIGGER update_inventory_after_order
//   AFTER INSERT ON orders
//   FOR EACH ROW
//   BEGIN
//     UPDATE inventories
//     SET stock = stock - NEW.quantity
//     WHERE id = NEW.inventory_id;
//   END;`

// 	// SQL trigger to adjust product price based on demand and availability
// 	adjustPriceTriggerSQL = `CREATE OR REPLACE FUNCTION adjust_product_price()
// RETURNS TRIGGER AS $$
// DECLARE
//     demand_count INT;
//     availability_count INT;
// BEGIN
//     -- Calculate demand within the last 2 hours
//     SELECT COUNT(*)
//     INTO demand_count
//     FROM orders
//     WHERE ordered_at >= NOW() - INTERVAL '2 hours' AND price >= 100;

//     -- Calculate availability (stock below 100)
//     SELECT COUNT(*)
//     INTO availability_count
//     FROM inventories
//     WHERE stock < 100;

//     -- Update product price if conditions are met
//     IF demand_count > 0 OR availability_count > 0 THEN
//         UPDATE inventories
//         SET price = price * 1.20; -- Increase price by 20%
//     END IF;

//     RETURN NEW;
// END;
// $$ LANGUAGE plpgsql;

// CREATE TRIGGER adjust_price_trigger
// AFTER INSERT ON orders
// FOR EACH ROW
// EXECUTE FUNCTION adjust_product_price();`
// )
package db

import (
	"errors"

	"gorm.io/gorm"
)

// SetUpDBTriggers sets up database triggers
func SetUpDBTriggers(db *gorm.DB) error {
	// Create trigger to update inventory stock after placing an order
	if err := db.Exec(placeOrderTriggerSQL).Error; err != nil {
		return errors.New("failed to create place_order trigger")
	}
	// Create trigger to adjust product price based on demand and availability
	if err := db.Exec(adjustPriceTriggerSQL).Error; err != nil {
		return errors.New("failed to create adjust_price trigger")
	}

	return nil
}

var (
	// SQL trigger to update inventory stock after placing an order
	placeOrderTriggerSQL = `
CREATE OR REPLACE FUNCTION update_inventory_stock()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE inventories
    SET stock = stock - NEW.quantity
    WHERE id = NEW.inventory_id;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_trigger WHERE tgname = 'update_inventory_after_order'
    ) THEN
        CREATE TRIGGER update_inventory_after_order
        AFTER INSERT ON orders
        FOR EACH ROW
        EXECUTE FUNCTION update_inventory_stock();
    END IF;
END $$;
`

	// SQL trigger to adjust product price based on demand and availability
	adjustPriceTriggerSQL = `
CREATE OR REPLACE FUNCTION adjust_product_price()
RETURNS TRIGGER AS $$
DECLARE
    demand_count INT;
    availability_count INT;
BEGIN
    -- Calculate demand within the last 2 hours
    SELECT COUNT(*)
    INTO demand_count
    FROM orders
    WHERE ordered_at >= NOW() - INTERVAL '2 hours';

    -- Calculate availability (stock below 100)
    SELECT COUNT(*)
    INTO availability_count
    FROM inventories
    WHERE stock < 100;

    -- Update product price if conditions are met
    IF demand_count > 0 OR availability_count > 0 THEN
        UPDATE inventories
        SET price = price * 1.20; -- Increase price by 20%
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_trigger WHERE tgname = 'adjust_price_trigger'
    ) THEN
        CREATE TRIGGER adjust_price_trigger
        AFTER INSERT ON orders
        FOR EACH ROW
        EXECUTE FUNCTION adjust_product_price();
    END IF;
END $$;
`
)
