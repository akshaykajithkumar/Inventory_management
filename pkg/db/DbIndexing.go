package db

import (
	"gorm.io/gorm"
)

//	func CreateIndexes(db *gorm.DB) error {
//		// Create the inverted index on product_name column
//		if err := db.Exec("CREATE INDEX idx_product_name ON inventories USING gin (to_tsvector('english', product_name));").Error; err != nil {
//			return err
//		}
//		if err := db.Exec("CREATE INDEX idx_order_id ON orders (id);").Error; err != nil {
//			return err
//		}
//		return nil
//	}
func CreateIndexes(db *gorm.DB) error {
	// Check if the index idx_product_name exists before attempting to create it
	var indexExists bool
	err := db.Raw("SELECT EXISTS (SELECT 1 FROM pg_indexes WHERE indexname = 'idx_product_name')").Row().Scan(&indexExists)
	if err != nil {
		return err
	}

	// If the index doesn't exist, create it
	if !indexExists {
		if err := db.Exec("CREATE INDEX idx_product_name ON inventories USING gin (to_tsvector('english', product_name));").Error; err != nil {
			return err
		}
	}

	// Check if the index idx_order_id exists before attempting to create it
	err = db.Raw("SELECT EXISTS (SELECT 1 FROM pg_indexes WHERE indexname = 'idx_order_id')").Row().Scan(&indexExists)
	if err != nil {
		return err
	}

	// If the index doesn't exist, create it
	if !indexExists {
		if err := db.Exec("CREATE INDEX idx_order_id ON orders (id);").Error; err != nil {
			return err
		}
	}

	return nil
}
