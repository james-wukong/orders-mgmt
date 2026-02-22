package models

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/db/dialect"
	"github.com/google/uuid"
)

type MenuItem struct {
	Base

	ID                uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	RestaurantID      uuid.UUID  `gorm:"type:uuid;not null" json:"restaurant_id"`
	CategoryID        *uuid.UUID `gorm:"type:uuid" json:"category_id"` // Pointer for nullable field
	Name              string     `gorm:"type:varchar(255);not null" json:"name"`
	Slug              string     `gorm:"type:varchar(255);not null" json:"slug"`
	Description       string     `gorm:"type:text" json:"description"`
	Price             float64    `gorm:"type:decimal(10,2);not null" json:"price"`
	DiscountPrice     *float64   `gorm:"type:decimal(10,2)" json:"discount_price"`
	ImageURL          string     `gorm:"type:text" json:"image_url"`
	IsVegetarian      bool       `gorm:"default:false" json:"is_vegetarian"`
	IsVegan           bool       `gorm:"default:false" json:"is_vegan"`
	IsGlutenFree      bool       `gorm:"default:false" json:"is_gluten_free"`
	IsSpicy           bool       `gorm:"default:false" json:"is_spicy"`
	SpiceLevel        int        `gorm:"type:integer" json:"spice_level"`
	Calories          int        `gorm:"type:integer" json:"calories"`
	PreparationTime   int        `gorm:"type:integer" json:"preparation_time"`
	IsAvailable       bool       `gorm:"default:true" json:"is_available"`
	IsFeatured        bool       `gorm:"default:false" json:"is_featured"`
	StockQuantity     int        `gorm:"type:integer" json:"stock_quantity"`
	LowStockThreshold int        `gorm:"type:integer;default:10" json:"low_stock_threshold"`
	Allergens         []string   `gorm:"type:text[]" json:"allergens"`
	DisplayOrder      int        `gorm:"type:integer;default:0" json:"display_order"`
	CreatedAt         time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt         time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
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
			fmt.Println("item data k: ", k, ", v: ", v[0])
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
