package models

import (
	"database/sql"
	"errors"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/google/uuid"
)

type User struct {
	Base

	ID                     uuid.UUID `json:"id"`
	Email                  string    `json:"email,omitempty"`
	PasswordHash           string    `json:"password_hash,omitempty"`
	FirstName              string    `json:"first_name,omitempty"`
	LastName               string    `json:"last_name,omitempty"`
	Phone                  string    `json:"phone"`
	Role                   string    `json:"role,omitempty"`
	IsActive               bool      `json:"is_active"`
	EmailVerified          bool      `json:"email_verified"`
	EmailVerificationToken *string   `json:"email_verification_token"`
	PasswordResetToken     *string   `json:"password_reset_token"`
	PasswordResetExpires   time.Time `json:"password_reset_expires"`
	LastLogin              time.Time `json:"last_login"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
}

func NewUser(conn db.Connection) *User {
	return &User{
		Base: Base{
			TableName: "users",
			Conn:      conn,
		},
	}
}

func (m *User) GetUserByEmail(email string) (string, error) {
	tag, err := m.Table(m.TableName).
		Where("email", "=", email).
		First()
	if len(tag) > 0 {
		return tag["id"].(string), nil
	}

	return "", err
}

func (m *User) CreateUser(item map[string][]string, tx *sql.Tx) (string, error) {
	var (
		columns      []string
		placeholders []string
		values       []interface{}
	)

	i := 1
	excludes := []string{"id"}
	for k, v := range item {
		if !slices.Contains(excludes, k) && len(v) > 0 && !strings.Contains(k, "__") {
			// fmt.Println("item data k: ", k, ", v: ", v[0])
			// itemData[k] = v[0]
			columns = append(columns, k)
			placeholders = append(placeholders, fmt.Sprintf("$%d", i))
			values = append(values, v[0])
			i++
		}
	}
	q := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s) ON CONFLICT (email) DO NOTHING RETURNING id",
		m.TableName,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)
	// fmt.Println(q)
	row, err := m.Conn.Query(q, values...)
	if err != nil {
		return "", err
	}

	if len(row) > 0 {
		return row[0]["id"].(string), nil
	}
	return "", errors.New("duplicated email")
}
