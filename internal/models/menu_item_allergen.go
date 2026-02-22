package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/GoAdminGroup/go-admin/modules/db"
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
	q := `INSERT INTO menu_item_allergens (menu_item_id, allergen_id)
		VALUES ($1, $2)
		ON CONFLICT (menu_item_id, allergen_id) 
		DO NOTHING`
	_, err := m.Conn.QueryWithTx(tx, q, menuItemID, allergenID)
	if err != nil {
		fmt.Println("error inserting composite table: ", err)
	}
	return err
}

// RemoveMenuItemAllergens remove old menu item and allergen composite key records
func (m *MenuItemAllergen) RemoveMenuItemAllergens(
	menuItemID string,
	allergenIDs []interface{},
	tx *sql.Tx,
) error {
	t := m.Table(m.TableName)
	if tx != nil {
		t = t.WithTx(tx)
	}
	q := t.Where("menu_item_id", "=", menuItemID)
	if len(allergenIDs) > 0 {

		fmt.Println("tx deleting menu item allergens: menu item id: ", menuItemID)
		return q.WhereNotIn("allergen_id", allergenIDs).Delete()
	}

	fmt.Println("deleting menu item allergens: menu item id: ", menuItemID)
	return q.Delete()

}
