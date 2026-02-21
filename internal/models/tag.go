package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/google/uuid"
)

type Tag struct {
	Base

	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func NewTag(conn db.Connection) *Tag {
	return &Tag{
		Base: Base{
			TableName: "tags",
			Conn:      conn,
		},
	}
}

func (m *Tag) InsertTag(name string, tx *sql.Tx) (string, error) {
	// 1. Try to find the tag by name
	id, err := m.GetTagIDByName(name)
	if err == nil {
		return id, nil
	}
	// 2. Insert if not found
	q := "INSERT INTO tags (name) VALUES ($1) RETURNING id"
	tagID, err := m.Conn.QueryWithTx(tx, q, name)
	if err != nil {
		fmt.Println("err during getting tag id", name, err)
		return "", err
	}
	return tagID[0]["id"].(string), err
}

func (m *Tag) GetTagIDByName(name string) (string, error) {
	tag, err := m.Table(m.TableName).
		Where("name", "=", name).
		First()
	if err == nil && len(tag) > 0 {
		return tag["id"].(string), nil
	}

	return "", err
}
