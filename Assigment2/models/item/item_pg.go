package item

import (
	"database/sql"
	"fmt"
)

const (
	addNewItemQuery = `
		INSERT INTO items (item_code, description, quantity, order_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id, item_code, description, quantity, order_id;
	`
	retrieveAllItemsQuery = `
		SELECT id, item_code, description, quantity, order_id
		FROM items;
	`
)

// Item represents a single item in the database
type Item struct {
	ID          int64  `json:"id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     int64  `json:"order_id"`
}

// Repository defines the interface for interacting with items
type Repository interface {
	Create(item Item) (Item, error)
	GetAllItems() ([]Item, error)
}

type itemPG struct {
	db *sql.DB
}

// NewItemPG creates a new item repository instance
func NewItemPG(db *sql.DB) Repository {
	return &itemPG{
		db: db,
	}
}

// Create inserts a new item into the database
func (m *itemPG) Create(item Item) (Item, error) {
	row := m.db.QueryRow(addNewItemQuery,
		item.ItemCode, item.Description, item.Quantity, item.OrderID)

	var newItem Item
	err := row.Scan(&newItem.ID, &newItem.ItemCode, &newItem.Description, &newItem.Quantity, &newItem.OrderID)

	if err != nil {
		return Item{}, fmt.Errorf("error creating item: %w", err)
	}

	return newItem, nil
}

// GetAllItems retrieves all items from the database
func (m *itemPG) GetAllItems() ([]Item, error) {
	rows, err := m.db.Query(retrieveAllItemsQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying items: %w", err)
	}
	defer rows.Close() // Ensure rows are closed even in case of errors

	var items []Item
	for rows.Next() {
		var item Item
		err := rows.Scan(&item.ID, &item.ItemCode, &item.Description, &item.Quantity, &item.OrderID)
		if err != nil {
			return nil, fmt.Errorf("error scanning items: %w", err)
		}
		items = append(items, item)
	}

	return items, nil
}
