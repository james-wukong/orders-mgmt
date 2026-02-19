package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetCategoriesTable(ctx *context.Context) table.Table {

	categories := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

	info := categories.GetInfo()

	// info.AddField("Id", "id", db.UUID)
	// info.AddField("ID", "id", db.UUID).
	// 	FieldDisplay(func(value types.FieldModel) interface{} {
	// 		if val, ok := value.Row["id"].(string); ok {
	// 			return val
	// 		}
	// 		if bytes, ok := value.Row["id"].([]byte); ok {
	// 			return string(bytes)
	// 		}
	// 		return value.Row["id"]
	// 	})
	// info.AddField("Restaurant_id", "restaurant_id", db.UUID)
	info.AddField("Restaurant Name", "name", db.Varchar).
		// SetTable("restaurants").
		FieldJoin(types.Join{
			Table:     "restaurants",   // The table to join with
			Field:     "restaurant_id", // The foreign key in current table
			JoinField: "id",            // The primary key in joined table
		}).
		// Tell GoAdmin which column specifically to display from the joined table
		FieldDisplay(func(value types.FieldModel) interface{} {
			// for key, val := range value.Row {
			// 	fmt.Printf("Key: %s | Value: %v\n", key, val)
			// }
			// return value.Value
			// 1. Check if the joined value is nil
			if value.Row["restaurants_goadmin_join_name"] == nil {
				return "No Restaurant" // Return a string or empty template.HTML
			}
			// 2. Safely return the value
			return value.Row["restaurants_goadmin_join_name"]
		}).
		FieldSortable()
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Slug", "slug", db.Varchar)
	info.AddField("Description", "description", db.Text)
	info.AddField("Image", "image_url", db.Text).FieldImage("50", "50",
		"http://localhost:8081/uploads/")
	info.AddField("Is_active", "is_active", db.Bool)
	info.AddField("Display_order", "display_order", db.Int4)
	info.AddField("Created At", "created_at", db.Timestamp)

	info.SetTable("categories").SetTitle("Categories").SetDescription("Categories")

	formList := categories.GetForm()
	formList.AddField("Id", "id", db.UUID, form.Default).
		FieldDisableWhenCreate()
	// formList.AddField("Restaurant_id", "restaurant_id", db.UUID, form.Text)
	formList.AddField("Restaurant Name", "restaurant_id", db.Varchar, form.SelectSingle).
		FieldOptionsFromTable("restaurants", "name", "id").
		FieldPlaceholder("Please select a restaurant").
		FieldMust()
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Slug", "slug", db.Varchar, form.Text)
	formList.AddField("Description", "description", db.Text, form.RichText)
	formList.AddField("Image URL", "image_url", db.Varchar, form.File)
	formList.AddField("Display Order", "display_order", db.Int, form.Number).
		FieldDefault("10")
	formList.AddField("Is Active", "is_active", db.Bool, form.Radio).
		FieldOptions(types.FieldOptions{
			{Text: "Active", Value: "true"},
			{Text: "InActive", Value: "false"},
		}).
		FieldDefault("false")
	formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenInsert()
	formList.AddField("Updated_at", "updated_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenUpdate()

	formList.SetTable("categories").SetTitle("Categories").SetDescription("Categories")

	return categories
}
