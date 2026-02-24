package models

import (
	"time"

	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/google/uuid"
)

type UserSession struct {
	Base

	ID              uuid.UUID              `json:"id"`
	UserID          uuid.UUID              `json:"user_id"`
	Token           string                 `json:"token"`
	DeviceInfo      map[string]interface{} `json:"device_info"`
	IPAddress       *string                `json:"ip_address"`
	ClientDevice    string                 `json:"client_device"`
	OperatingSystem string                 `json:"operating_system"`
	UserAgent       string                 `json:"user_agent"`
	ExpiresAt       time.Time              `json:"expires_at"`
	CreatedAt       time.Time              `json:"created_at"`
}

func NewUserSession(conn db.Connection) *UserSession {
	return &UserSession{
		Base: Base{
			TableName: "user_sessions",
			Conn:      conn,
		},
	}
}
