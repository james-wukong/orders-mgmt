package models

import (
	"database/sql"
	"time"

	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/db/dialect"
	"github.com/google/uuid"
)

type MenuItemTag struct {
	Base

	MenuItemID uuid.UUID `gorm:"type:uuid;not null" json:"menu_item_id"`
	TagID      uuid.UUID `gorm:"type:uuid;not null" json:"tag_id"`

	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
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
	t := m.Table(m.TableName)
	if tx != nil {
		t = m.Table(m.TableName).WithTx(tx)
	}
	_, err := t.Insert(dialect.H{
		"menu_item_id": menuItemID,
		"tag_id":       tagID,
	})

	return err
}
