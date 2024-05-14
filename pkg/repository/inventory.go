package repository

import (
	"errors"
	"main/pkg/domain"
	interfaces "main/pkg/repository/interface"
	"main/pkg/utils/models"
	"strconv"

	"gorm.io/gorm"
)

type inventoryRepository struct {
	DB *gorm.DB
}

func NewInventoryRepository(DB *gorm.DB) interfaces.InventoryRepository {
	return &inventoryRepository{
		DB: DB,
	}
}

func (i *inventoryRepository) AddInventory(inventory models.Inventory) (models.InventoryResponse, error) {

	query := `
		INSERT INTO inventories (product_name, description, stock, price)
		VALUES ( ?, ?, ?, ?);
		`
	i.DB.Exec(query, inventory.ProductName, inventory.Description, inventory.Stock, inventory.Price)

	var inventoryResponse models.InventoryResponse

	return inventoryResponse, nil

}

func (i *inventoryRepository) CheckInventory(pid int) (bool, error) {
	var k int
	err := i.DB.Raw("SELECT COUNT(*) FROM inventories WHERE id=?", pid).Scan(&k).Error
	if err != nil {
		return false, err
	}

	if k == 0 {
		return false, err
	}

	return true, err
}

func (i *inventoryRepository) UpdateInventory(pid int, invData models.UpdateInventory) (models.Inventory, error) {

	// Check the database connection
	if i.DB == nil {
		return models.Inventory{}, errors.New("database connection is nil")
	}

	if invData.ProductName != "" && invData.ProductName != "string" {
		if err := i.DB.Exec("UPDATE inventories SET product_name = ? WHERE id= ?", invData.ProductName, pid).Error; err != nil {
			return models.Inventory{}, err
		}
	}
	if invData.Description != "" && invData.Description != "string" {
		if err := i.DB.Exec("UPDATE inventories SET description = ? WHERE id= ?", invData.Description, pid).Error; err != nil {
			return models.Inventory{}, err
		}
	}
	if invData.Stock != 0 {
		if err := i.DB.Exec("UPDATE inventories SET stock =  ? WHERE id= ?", invData.Stock, pid).Error; err != nil {
			return models.Inventory{}, err
		}
	}

	if invData.Price != 0 {
		if err := i.DB.Exec("UPDATE inventories SET price =  ? WHERE id= ?", invData.Price, pid).Error; err != nil {
			return models.Inventory{}, err
		}
	}
	// Retrieve the update
	var inventory models.Inventory
	if err := i.DB.Raw("SELECT * FROM inventories WHERE id=?", pid).Scan(&inventory).Error; err != nil {
		return models.Inventory{}, err
	}

	return inventory, nil
}

func (i *inventoryRepository) DeleteInventory(inventoryID string) error {
	id, err := strconv.Atoi(inventoryID)
	if err != nil {
		return errors.New("converting into integer not happened")
	}

	result := i.DB.Exec("DELETE FROM inventories WHERE id = ?", id)

	if result.RowsAffected < 1 {
		return errors.New("no records with that ID exist")
	}

	return nil
}

func (i *inventoryRepository) CheckStock(pid int) (int, error) {
	var k int
	if err := i.DB.Raw("SELECT stock FROM inventories WHERE id=?", pid).Scan(&k).Error; err != nil {
		return 0, err
	}
	return k, nil
}

func (i *inventoryRepository) CheckPrice(pid int) (float64, error) {
	var k float64
	err := i.DB.Raw("SELECT price FROM inventories WHERE id=?", pid).Scan(&k).Error
	if err != nil {
		return 0, err
	}

	return k, nil
}

// func (a *inventoryRepository) SearchProducts(key string, page, limit int, sortBy string) ([]domain.Inventory, error) {
// 	var inventories []domain.Inventory
// 	offset := (page - 1) * limit

// 	// Prepare the SQL query
// 	var query string
// 	if sortBy == "asc" {
// 		query = `
// 			 SELECT id, product_name, description, stock, price
// 			 FROM inventories
// 			 WHERE to_tsvector('english', product_name) @@ to_tsquery('english', ?)
// 			 ORDER BY price ASC
// 			 LIMIT ? OFFSET ?
// 		 `
// 	} else if sortBy == "desc" {
// 		query = `
// 			 SELECT id, product_name, description, stock, price
// 			 FROM inventories
// 			 WHERE to_tsvector('english', product_name) @@ to_tsquery('english', ?)
// 			 ORDER BY price DESC
// 			 LIMIT ? OFFSET ?
// 		 `
// 	}

// 	// Execute the query
// 	rows, err := a.DB.Raw(query, key, limit, offset).Rows()
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	// Iterate through the rows and scan each into an Inventory struct
// 	for rows.Next() {
// 		var inventory domain.Inventory
// 		if err := rows.Scan(&inventory.ID, &inventory.ProductName, &inventory.Description, &inventory.Stock, &inventory.Price); err != nil {
// 			return nil, err
// 		}
// 		inventories = append(inventories, inventory)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}

//		return inventories, nil
//	}
func (a *inventoryRepository) SearchProducts(key string, page, limit int, sortBy string) ([]domain.Inventory, error) {
	var inventories []domain.Inventory
	offset := (page - 1) * limit

	// Prepare the search key for prefix matching
	searchKey := key + ":*"

	// Prepare the SQL query
	var query string
	if sortBy == "asc" {
		query = `
			 SELECT id, product_name, description, stock, price
			 FROM inventories
			 WHERE to_tsvector('english', product_name) @@ to_tsquery('english', ?)
			 ORDER BY price ASC
			 LIMIT ? OFFSET ?
		 `
	} else if sortBy == "desc" {
		query = `
			 SELECT id, product_name, description, stock, price
			 FROM inventories
			 WHERE to_tsvector('english', product_name) @@ to_tsquery('english', ?)
			 ORDER BY price DESC
			 LIMIT ? OFFSET ?
		 `
	}

	// Execute the query
	rows, err := a.DB.Raw(query, searchKey, limit, offset).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate through the rows and scan each into an Inventory struct
	for rows.Next() {
		var inventory domain.Inventory
		if err := rows.Scan(&inventory.ID, &inventory.ProductName, &inventory.Description, &inventory.Stock, &inventory.Price); err != nil {
			return nil, err
		}
		inventories = append(inventories, inventory)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return inventories, nil
}

func (i *inventoryRepository) GetInventoryByID(inventoryID string) (domain.Inventory, error) {

	var inventory domain.Inventory

	err := i.DB.Raw("SELECT * FROM inventories WHERE id = ?", inventoryID).Scan(&inventory).Error
	if err != nil {
		// If there's an error or the inventory doesn't exist, return an empty inventory object and the error
		return domain.Inventory{}, err
	}

	return inventory, nil
}
