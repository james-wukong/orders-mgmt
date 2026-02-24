package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetCartsTable(ctx *context.Context) table.Table {

	carts := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

	info := carts.GetInfo()

	info.AddField("Id", "id", db.UUID)
	info.AddField("User_id", "user_id", db.UUID)
	info.AddField("Restaurant_id", "restaurant_id", db.UUID)
	info.AddField("Created_at", "created_at", db.Timestamp)
	info.AddField("Updated_at", "updated_at", db.Timestamp)
	info.AddField("Session_id", "session_id", db.Varchar)

	info.SetTable("carts").SetTitle("Carts").SetDescription("Carts")

	formList := carts.GetForm()
	formList.AddField("Id", "id", db.UUID, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("User_id", "user_id", db.UUID, form.Text)
	formList.AddField("Restaurant_id", "restaurant_id", db.UUID, form.Text)
	formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenInsert()
	formList.AddField("Updated_at", "updated_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenUpdate()
	formList.AddField("Session_id", "session_id", db.Varchar, form.Text)

	formList.SetTable("carts").SetTitle("Carts").SetDescription("Carts")

	return carts
}
