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

	info := categories.GetInfo().HideFilterArea()
	// info.SetPrimaryKey("id", db.Varchar)
	info.SetSortField("created_at").SetSortDesc()
	info.AddField("ID", "id", db.UUID).
		FieldDisplay(func(value types.FieldModel) interface{} {
			// value.Value might be 0, but value.Row["id"]
			// contains the raw string from the database.
			if val, ok := value.Row["id"].(string); ok {
				return val
			}
			// If it's stored as a byte slice (common in some drivers),
			// you might need to convert it:
			if bytes, ok := value.Row["id"].([]byte); ok {
				return string(bytes)
			}
			return value.Row["id"]
		})
	info.AddField("Category Name", "name", db.Varchar).FieldSortable()
	info.AddField("Category Slug", "slug", db.Varchar).FieldSortable()
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
	info.AddField("Image", "image_url", db.Varchar).FieldImage("50", "50",
		"http://localhost:8081/uploads/")
	info.AddField("Created At", "created_at", db.Time).FieldSortable()
	info.AddField("Updated At", "updated_at", db.Time)
	info.SetTable("categories").SetTitle("Categories").SetDescription("Categories")

	formList := categories.GetForm()
	formList.AddField("Restaurant Name", "restaurant_id", db.Varchar, form.SelectSingle).
		FieldOptionsFromTable("restaurants", "name", "id").
		FieldPlaceholder("Please select a restaurant").
		FieldMust()
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("SLug", "slug", db.Varchar, form.Text)
	formList.AddField("Image URL", "image_url", db.Varchar, form.File)
	formList.AddField("Description", "description", db.Varchar, form.TextArea)
	formList.AddField("Display Order", "display_order", db.Int, form.Number).
		FieldDefault("10")
	formList.AddField("Is Active", "is_active", db.Bool, form.Radio).
		FieldOptions(types.FieldOptions{
			{Text: "Active", Value: "true"},
			{Text: "InActive", Value: "false"},
		}).
		FieldDefault("false")

	formList.SetTable("categories").SetTitle("Categories").SetDescription("Categories")

	return categories
}
