package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetCartitemsTable(ctx *context.Context) table.Table {

	cartItems := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

	info := cartItems.GetInfo()

	info.AddField("Updated_at", "updated_at", db.Timestamp)
	info.AddField("Cart_id", "cart_id", db.UUID)
	info.AddField("Menu_item_id", "menu_item_id", db.UUID)
	info.AddField("Quantity", "quantity", db.Int4)
	info.AddField("Unit_price", "unit_price", db.Numeric)
	info.AddField("Id", "id", db.UUID)
	info.AddField("Created_at", "created_at", db.Timestamp)
	info.AddField("Special_instructions", "special_instructions", db.Text)

	info.SetTable("cart_items").SetTitle("Cartitems").SetDescription("Cartitems")

	formList := cartItems.GetForm()
	formList.AddField("Updated_at", "updated_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenUpdate()
	formList.AddField("Cart_id", "cart_id", db.UUID, form.Text)
	formList.AddField("Menu_item_id", "menu_item_id", db.UUID, form.Text)
	formList.AddField("Quantity", "quantity", db.Int4, form.Number)
	formList.AddField("Unit_price", "unit_price", db.Numeric, form.Number)
	formList.AddField("Id", "id", db.UUID, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenInsert()
	formList.AddField("Special_instructions", "special_instructions", db.Text, form.RichText)

	formList.SetTable("cart_items").SetTitle("Cartitems").SetDescription("Cartitems")

	return cartItems
}
