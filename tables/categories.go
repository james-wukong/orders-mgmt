package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetCategoriesTable(ctx *context.Context) table.Table {

	categories := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

	info := categories.GetInfo()

	info.AddField("Id", "id", db.UUID)
	info.AddField("Restaurant_id", "restaurant_id", db.UUID)
	info.AddField("Display_order", "display_order", db.Int4)
	info.AddField("Is_active", "is_active", db.Bool)
	info.AddField("Created_at", "created_at", db.Timestamp)
	info.AddField("Updated_at", "updated_at", db.Timestamp)
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Slug", "slug", db.Varchar)
	info.AddField("Description", "description", db.Text)
	info.AddField("Image_url", "image_url", db.Text)

	info.SetTable("categories").SetTitle("Categories").SetDescription("Categories")

	formList := categories.GetForm()
	formList.AddField("Id", "id", db.UUID, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("Restaurant_id", "restaurant_id", db.UUID, form.Text)
	formList.AddField("Display_order", "display_order", db.Int4, form.Number)
	formList.AddField("Is_active", "is_active", db.Bool, form.Text)
	formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenInsert()
	formList.AddField("Updated_at", "updated_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenUpdate()
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Slug", "slug", db.Varchar, form.Text)
	formList.AddField("Description", "description", db.Text, form.RichText)
	formList.AddField("Image_url", "image_url", db.Text, form.RichText)

	formList.SetTable("categories").SetTitle("Categories").SetDescription("Categories")

	return categories
}
