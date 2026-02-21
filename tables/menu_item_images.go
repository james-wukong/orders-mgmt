package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetMenuitemimagesTable(ctx *context.Context) table.Table {

	menuItemImages := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

	info := menuItemImages.GetInfo()

	// info.HideNewButton()

	// info.HideEditButton()
	// info.HideDeleteButton()

	info.AddField("Id", "id", db.UUID).
		FieldDisplay(func(value types.FieldModel) interface{} {
			if val, ok := value.Row["id"].(string); ok {
				return val
			}
			if bytes, ok := value.Row["id"].([]byte); ok {
				return string(bytes)
			}
			return value.Row["id"]
		}).
		FieldHide()
	info.AddField("Menu_item_id", "menu_item_id", db.UUID)
	info.AddField("Menu Item", "name", db.Varchar).
		// SetTable("restaurants").
		FieldJoin(types.Join{
			Table:     "menu_items",   // The table to join with
			Field:     "menu_item_id", // The foreign key in current table
			JoinField: "id",           // The primary key in joined table
		}).
		// Tell GoAdmin which column specifically to display from the joined table
		FieldDisplay(func(value types.FieldModel) interface{} {
			// 1. Check if the joined value is nil
			if value.Row["menu_items_goadmin_join_name"] == nil {
				return "No Menu Items" // Return a string or empty template.HTML
			}
			// 2. Safely return the value
			return value.Row["menu_items_goadmin_join_name"]
		}).
		FieldSortable()
	info.AddField("Is_primary", "is_primary", db.Bool).FieldBool("true", "false")
	info.AddField("Image_url", "image_url", db.Text).FieldImage("50", "50",
		"http://localhost:8081/uploads/")
	info.AddField("Display_order", "display_order", db.Int4)
	info.AddField("Created_at", "created_at", db.Timestamp)

	info.SetTable("menu_item_images").SetTitle("Menuitemimages").SetDescription("Menuitemimages")

	formList := menuItemImages.GetForm()
	formList.AddField("Id", "id", db.UUID, form.Default).
		FieldDisableWhenCreate()
	// formList.AddField("Menu_item_id", "menu_item_id", db.UUID, form.Text)
	formList.AddField("Menu Item", "menu_item_id", db.Varchar, form.SelectSingle).
		FieldOptionsFromTable("menu_items", "name", "id").
		FieldPlaceholder("Please select a menu item").
		FieldMust()
	formList.AddField("Is_primary", "is_primary", db.Bool, form.Radio).
		FieldOptions(types.FieldOptions{
			{Text: "Primary", Value: "true"},
			{Text: "Other", Value: "false"},
		}).
		FieldDefault("false")
	formList.AddField("Image_url", "image_url", db.Text, form.File)
	formList.AddField("Display_order", "display_order", db.Int4, form.Number)
	formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenInsert()

	formList.SetTable("menu_item_images").SetTitle("Menuitemimages").SetDescription("Menuitemimages")

	return menuItemImages
}
