package models

import (
	"database/sql"
	"errors"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/db/dialect"
	"github.com/google/uuid"
)

type MenuItem struct {
	Base

	ID                uuid.UUID  `json:"id"`
	RestaurantID      uuid.UUID  `json:"restaurant_id"`
	CategoryID        *uuid.UUID `json:"category_id"` // Pointer for nullable field
	Name              string     `json:"name"`
	Slug              string     `json:"slug"`
	Description       string     `json:"description"`
	Price             float64    `json:"price"`
	DiscountPrice     *float64   `json:"discount_price"`
	ImageURL          string     `json:"image_url"`
	IsVegetarian      bool       `json:"is_vegetarian"`
	IsVegan           bool       `json:"is_vegan"`
	IsGlutenFree      bool       `json:"is_gluten_free"`
	IsSpicy           bool       `json:"is_spicy"`
	SpiceLevel        int        `json:"spice_level"`
	Calories          int        `json:"calories"`
	PreparationTime   int        `json:"preparation_time"`
	IsAvailable       bool       `json:"is_available"`
	IsFeatured        bool       `json:"is_featured"`
	StockQuantity     int        `json:"stock_quantity"`
	LowStockThreshold int        `json:"low_stock_threshold"`
	Allergens         []string   `json:"allergens"`
	DisplayOrder      int        `json:"display_order"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}

func NewMenuItem(conn db.Connection) *MenuItem {
	return &MenuItem{
		Base: Base{
			TableName: "menu_items",
			Conn:      conn,
		},
	}
}

func (m *MenuItem) UpdateMenuItem(item map[string][]string, tx *sql.Tx) error {
	t := m.Table(m.TableName)
	if tx != nil {
		t = t.WithTx(tx)
	}
	// 4.2 update menu items with new data
	itemData := make(dialect.H)
	for k, v := range item {
		if k != "id" && len(v) > 0 && !strings.Contains(k, "__") {
			// fmt.Println("item data k: ", k, ", v: ", v[0])
			itemData[k] = v[0]
		}
	}
	_, updateErr := t.
		Where("id", "=", item["id"][0]).
		Update(itemData)
	if db.CheckError(updateErr, db.UPDATE) {
		fmt.Println("updating menu item, with err:", updateErr)
		return updateErr
	}
	return nil
}

func (m *MenuItem) CreateMenuItem(item map[string][]string, tx *sql.Tx) (string, error) {
	var (
		columns      []string
		placeholders []string
		values       []interface{}
	)

	i := 1
	excludes := []string{"id", "created_at", "updated_at"}
	for k, v := range item {
		if !slices.Contains(excludes, k) && len(v) > 0 && !strings.Contains(k, "__") {
			fmt.Println("item data k: ", k, ", v: ", v[0])
			// itemData[k] = v[0]
			columns = append(columns, k)
			placeholders = append(placeholders, fmt.Sprintf("$%d", i))
			values = append(values, v[0])
			i++
		}
	}
	q := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s) ON CONFLICT (restaurant_id, slug) DO NOTHING RETURNING id",
		m.TableName,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)

	row, err := m.Conn.QueryWithTx(tx, q, values...)
	if err != nil {
		return "", err
	}

	if len(row) > 0 {
		return row[0]["id"].(string), nil
	}
	return "", errors.New("no id returned")
}
