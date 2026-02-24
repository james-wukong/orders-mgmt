package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/google/uuid"
)

var (
	ErrNotFound = errors.New("out of index")
)

type MenuItemTag struct {
	Base

	MenuItemID uuid.UUID `json:"menu_item_id"`
	TagID      uuid.UUID `json:"tag_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewMenuItemTag(conn db.Connection) *MenuItemTag {
	return &MenuItemTag{
		Base: Base{
			TableName: "menu_item_tags",
			Conn:      conn,
		},
	}
}

func (m *MenuItemTag) InsertMenuItemTag(menuItemID, tagID string, tx *sql.Tx) error {
	fmt.Println("start inserting composite table: tag", tagID, ", menu: ", menuItemID)
	q := `INSERT INTO menu_item_tags (menu_item_id, tag_id)
		VALUES ($1, $2)
		ON CONFLICT (menu_item_id, tag_id) 
		DO NOTHING`
	_, err := m.Conn.QueryWithTx(tx, q, menuItemID, tagID)
	if err != nil {
		fmt.Println("error inserting composite table: ", err)
	}
	return err
}

// RemoveMenuItemTags remove old menu item and tag composite key records
func (m *MenuItemTag) RemoveMenuItemTags(menuItemID string, tagIDs []interface{}, tx *sql.Tx) error {
	t := m.Table(m.TableName)
	if tx != nil {
		t = t.WithTx(tx)
	}
	q := t.Where("menu_item_id", "=", menuItemID)
	if len(tagIDs) > 0 {
		fmt.Println("tx deleting menu item tags: menu item id: ", menuItemID)
		return q.WhereNotIn("tag_id", tagIDs).Delete()
	}

	fmt.Println("deleting menu item tags: menu item id: ", menuItemID)
	return q.Delete()

}
