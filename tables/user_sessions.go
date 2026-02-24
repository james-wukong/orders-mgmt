package tables

import (
	"fmt"
	"time"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetUsersessionsTable(ctx *context.Context) table.Table {

	userSessions := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

	info := userSessions.GetInfo().HideFilterArea()
	info.AddField("Id", "id", db.UUID).
		FieldDisplay(func(value types.FieldModel) interface{} {
			if val, ok := value.Row["id"].(string); ok {
				return val
			}
			if bytes, ok := value.Row["id"].([]byte); ok {
				return string(bytes)
			}
			return fmt.Sprintf("%s", value.Row["id"])
		}).
		FieldHide()
	info.AddField("User ID", "user_id", db.UUID).FieldFilterable()
	info.AddField("Token", "token", db.Varchar).FieldHide()
	info.AddField("IP Address", "ip_address", db.Varchar).FieldFilterable()
	info.AddField("Device", "client_device", db.Varchar).FieldFilterable()
	info.AddField("OS", "operating_system", db.Varchar).FieldFilterable()
	info.AddField("Expires At", "expires_at", db.Timestamp)
	info.AddField("Created_at", "created_at", db.Timestamp)
	// info.AddField("Updated_at", "updated_at", db.Timestamp)
	info.SetTable("user_sessions").SetTitle("Usersessions").SetDescription("Usersessions")

	formList := userSessions.GetForm()
	formList.AddField("Id", "id", db.UUID, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("User ID", "user_id", db.UUID, form.Text).FieldMust()
	formList.AddField("Token", "token", db.Varchar, form.Text).FieldMust()

	// JSONB Field handling
	formList.AddField("Device Info", "device_info", db.JSON, form.Code).
		FieldDefault(`{}`)

	formList.AddField("IP Address", "ip_address", db.Varchar, form.Text)
	formList.AddField("Client Device", "client_device", db.Varchar, form.Text).
		FieldDefault("Unknown")
	formList.AddField("Operating System", "operating_system", db.Varchar, form.Text).
		FieldDefault("Windows 10")
	formList.AddField("User Agent", "user_agent", db.Varchar, form.TextArea)

	formList.AddField("Expires At", "expires_at", db.Timestamp, form.Datetime).FieldMust()
	formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime).
		FieldDefault(time.Now().Format("2006-01-02 15:04:05")). // Set initial value
		FieldHide().FieldNowWhenInsert()
	formList.AddField("Updated_at", "updated_at", db.Timestamp, form.Datetime).
		FieldDefault(time.Now().Format("2006-01-02 15:04:05")). // Set initial value
		FieldHide().FieldNowWhenUpdate()
	formList.SetTable("user_sessions").SetTitle("Usersessions").SetDescription("Usersessions")

	return userSessions
}
