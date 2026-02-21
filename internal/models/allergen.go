package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/google/uuid"
)

type Allergen struct {
	Base

	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func NewAllergen(conn db.Connection) *Allergen {
	return &Allergen{
		Base: Base{
			TableName: "allergens",
			Conn:      conn,
		},
	}
}

func (m *Allergen) InsertAllergen(name string, tx *sql.Tx) (string, error) {
	// 1. Try to find the tag by name
	id, err := m.GetAllergenIDByName(name)
	if err == nil {
		return id, nil
	}
	// 2. Insert if not found
	q := "INSERT INTO allergens (name) VALUES ($1) RETURNING id"
	allergenID, err := m.Conn.QueryWithTx(tx, q, name)
	if err != nil {
		fmt.Println("err during getting allergen id", name, err)
		return "", err
	}
	return allergenID[0]["id"].(string), err
}

func (m *Allergen) GetAllergenIDByName(name string) (string, error) {
	allergen, err := m.Table(m.TableName).
		Where("name", "=", name).
		First()
	if err == nil && len(allergen) > 0 {
		return allergen["id"].(string), nil
	}

	return "", err
}
