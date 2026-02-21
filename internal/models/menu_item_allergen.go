package models

import (
	"database/sql"
	"time"

	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/db/dialect"
	"github.com/google/uuid"
)

type MenuItemAllergen struct {
	Base

	MenuItemID uuid.UUID `gorm:"type:uuid;not null" json:"menu_item_id"`
	AllergenID uuid.UUID `gorm:"type:uuid;not null" json:"allergen_id"`

	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func NewMenuItemAllergen(conn db.Connection) *MenuItemAllergen {
	return &MenuItemAllergen{
		Base: Base{
			TableName: "menu_item_allergens",
			Conn:      conn,
		},
	}
}

func (m *MenuItemAllergen) InsertMenuItemAllergen(menuItemID, allergenID string, tx *sql.Tx) error {
	t := m.Table(m.TableName)
	if tx != nil {
		t = m.Table(m.TableName).WithTx(tx)
	}
	_, err := t.Insert(dialect.H{
		"menu_item_id": menuItemID,
		"allergen_id":  allergenID,
	})

	return err
}
