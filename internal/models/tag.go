package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/google/uuid"
)

type Tag struct {
	Base

	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewTag(conn db.Connection) *Tag {
	return &Tag{
		Base: Base{
			TableName: "tags",
			Conn:      conn,
		},
	}
}

func (m *Tag) UpsertTag(name string, tx *sql.Tx) (string, error) {
	// 1. Upsert a tag
	var (
		tagID []map[string]interface{}
		err   error
		id    string
	)
	q := `INSERT INTO tags (name) VALUES ($1) 
	ON CONFLICT (name) 
	DO UPDATE SET updated_at = NOW()
	RETURNING id`
	if tx != nil {
		tagID, err = m.Conn.QueryWithTx(tx, q, name)
	} else {
		tagID, err = m.Conn.Query(q, name)
	}
	// 2. Handle potential SQL errors immediately
	if err != nil {
		fmt.Println("error inserting tag: ", err)
		return "", fmt.Errorf("upsert tag error: %v", err)
	}
	// 3. Safe extraction of the ID
	rawID, ok := tagID[0]["id"]
	if !ok || rawID == nil {
		fmt.Println("error extraction of the ID")
		return "", errors.New("upsert tag: id field missing from result")
	}

	// 4. Flexible type conversion (handles string or []byte)
	switch v := rawID.(type) {
	case string:
		id = v
	case []byte:
		id = string(v)
	default:
		id = fmt.Sprintf("%v", v)
	}

	fmt.Println("inserted tag id: ", id)
	return id, nil
}

func (m *Tag) GetTagIDByName(name string) (string, error) {
	tag, _ := m.Table(m.TableName).
		Where("name", "=", name).
		First()
	if len(tag) > 0 {
		return tag["id"].(string), nil
	}

	return "", nil
}
