package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/google/uuid"
)

type Allergen struct {
	Base

	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewAllergen(conn db.Connection) *Allergen {
	return &Allergen{
		Base: Base{
			TableName: "allergens",
			Conn:      conn,
		},
	}
}

func (m *Allergen) UpsertAllergen(name string, tx *sql.Tx) (string, error) {
	// 1. Upsert a tag
	var (
		allergenID []map[string]interface{}
		err        error
		id         string
	)
	q := `INSERT INTO allergens (name) VALUES ($1) 
	ON CONFLICT (name) 
	DO UPDATE SET updated_at = NOW()
	RETURNING id`
	if tx != nil {
		allergenID, err = m.Conn.QueryWithTx(tx, q, name)
	} else {
		allergenID, err = m.Conn.Query(q, name)
	}
	// 2. Handle potential SQL errors immediately
	if err != nil {
		return "", fmt.Errorf("upsert allergen error: %v", err)
	}
	// 3. Safe extraction of the ID
	rawID, ok := allergenID[0]["id"]
	if !ok || rawID == nil {
		return "", errors.New("upsert allergen: id field missing from result")
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

	fmt.Println("inserted allergen id: ", id)
	return id, nil
}

func (m *Allergen) GetAllergenIDByName(name string) (string, error) {
	allergen, _ := m.Table(m.TableName).
		Where("name", "=", name).
		First()
	if len(allergen) > 0 {
		return allergen["id"].(string), nil
	}

	return "", nil
}
