package db

import (
	"gorm.io/gorm"
)

func CreateIndexes(db *gorm.DB) error {
	// Create the inverted index on product_name column
	if err := db.Exec("CREATE INDEX idx_product_name ON inventories USING gin (to_tsvector('english', product_name));").Error; err != nil {
		return err
	}
	if err := db.Exec("CREATE INDEX idx_order_id ON orders (id);").Error; err != nil {
		return err
	}
	return nil
}
