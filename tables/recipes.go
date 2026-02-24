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

func GetRecipesTable(dbConn db.Connection) table.Generator {
	return func(ctx *context.Context) table.Table {
		recipes := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

		info := recipes.GetInfo().SetPrimaryKey("id", db.UUID)

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
		info.AddField("Name", "name", db.Varchar).FieldSortable()
		info.AddField("Menu Item", "name", db.Varchar).
			// SetTable("restaurants").
			FieldJoin(types.Join{
				Table:      "menu_items", // The table to join with
				TableAlias: "mi",
				Field:      "menu_item_id", // The foreign key in current table
				JoinField:  "id",           // The primary key in joined table
			}).
			// Tell GoAdmin which column specifically to display from the joined table
			FieldDisplay(func(value types.FieldModel) interface{} {
				// 1. Check if the joined value is nil
				if value.Row["mi_goadmin_join_name"] == nil {
					return "No Menu Item" // Return a string or empty template.HTML
				}
				// 2. Safely return the value
				return value.Row["mi_goadmin_join_name"]
			}).
			FieldSortable()
		info.AddField("Preparation_time", "preparation_time", db.Interval)
		info.AddField("Cooking_time", "cooking_time", db.Interval)
		info.AddField("Serving_size", "serving_size", db.Int4)
		info.AddField("Yield_quantity", "yield_quantity", db.Numeric)
		info.AddField("Yield_unit", "yield_unit", db.Varchar)
		info.AddField("Description", "description", db.Text)
		info.AddField("Instructions", "instructions", db.Text)
		info.AddField("Notes", "notes", db.Text)
		info.AddField("Created_at", "created_at", db.Timestamp)
		info.AddField("Updated_at", "updated_at", db.Timestamp)

		info.SetTable("recipes").SetTitle("Recipes").SetDescription("Recipes")

		formList := recipes.GetForm().SetPrimaryKey("id", db.UUID)

		formList.AddField("Id", "id", db.UUID, form.Default).
			FieldDisableWhenCreate()
		formList.AddField("Name", "name", db.Varchar, form.Text).FieldMust()
		// formList.AddField("Menu Item", "menu_item_id", db.UUID, form.Text).FieldHide()
		formList.AddField("Menu Item", "menu_item_id", db.Varchar, form.SelectSingle).
			// 1. Create the Dropdown List
			FieldOptionInitFn(func(model types.FieldModel) types.FieldOptions {
				var c types.FieldOptions
				// Logic for UPDATING
				records, _ := db.WithDriver(dbConn).
					Table("menu_items").
					Select("id", "name").
					All()

				if len(records) == 0 {
					return nil
				}
				for _, v := range records {
					if model.IsUpdate() && v["id"] == model.Value {
						return types.FieldOptions{
							types.FieldOption{
								Text:     fmt.Sprint(v["name"]),
								Value:    fmt.Sprint(v["id"]),
								Selected: true,
							},
						}
					}
					c = append(c, types.FieldOption{
						Text:  fmt.Sprint(v["name"]),
						Value: fmt.Sprint(v["id"]),
					})
				}
				return c // This will display as plain text in the form
			}).
			FieldNotAllowEdit()
		formList.AddField("Preparation_time", "preparation_time", db.Varchar, form.Text).
			FieldHelpMsg("Format: 00:30:00 for 30 minutes").
			FieldDefault("00:30:00").
			FieldMust()
		formList.AddField("Cooking_time", "cooking_time", db.Varchar, form.Text).
			FieldHelpMsg("Format: 00:30:00 for 30 minutes").
			FieldDefault("00:30:00").
			FieldMust()
		formList.AddField("Serving_size", "serving_size", db.Int4, form.Number).
			FieldDefault("10").
			FieldMust()
		formList.AddField("Yield_quantity", "yield_quantity", db.Numeric, form.Number).
			FieldDefault("10").
			FieldMust()
		formList.AddField("Yield_unit", "yield_unit", db.Varchar, form.Text).FieldMust()
		formList.AddField("Description", "description", db.Text, form.RichText)
		formList.AddField("Instructions", "instructions", db.Text, form.RichText)
		formList.AddField("Notes", "notes", db.Text, form.RichText)
		formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime).
			FieldDefault(time.Now().Format("2006-01-02 15:04:05")). // Set initial value
			FieldHide().FieldNowWhenInsert()
		formList.AddField("Updated_at", "updated_at", db.Timestamp, form.Datetime).
			FieldDefault(time.Now().Format("2006-01-02 15:04:05")). // Set initial value
			FieldHide().FieldNowWhenUpdate()

		formList.SetTable("recipes").SetTitle("Recipes").SetDescription("Recipes")

		return recipes
	}
}
