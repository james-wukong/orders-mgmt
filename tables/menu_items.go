package tables

import (
	"fmt"

	"github.com/GoAdminGroup/go-admin/context"    // 需导入此包
	"github.com/GoAdminGroup/go-admin/modules/db" // 需导入此包

	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	selection "github.com/GoAdminGroup/go-admin/template/types/form/select"
)

func GetMenuitemsTable(dbConn db.Connection) table.Generator {
	return func(ctx *context.Context) table.Table {

		menuItems := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

		info := menuItems.GetInfo().HideFilterArea()
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
		info.AddField("Restaurant Name", "name", db.Varchar).
			// SetTable("restaurants").
			FieldJoin(types.Join{
				Table:     "restaurants",   // The table to join with
				Field:     "restaurant_id", // The foreign key in current table
				JoinField: "id",            // The primary key in joined table
			}).
			// Tell GoAdmin which column specifically to display from the joined table
			FieldDisplay(func(value types.FieldModel) interface{} {
				// 1. Check if the joined value is nil
				if value.Row["restaurants_goadmin_join_name"] == nil {
					return "No Restaurant" // Return a string or empty template.HTML
				}
				// 2. Safely return the value
				return value.Row["restaurants_goadmin_join_name"]
			}).
			FieldSortable()
		info.AddField("Category Name", "name", db.Varchar).
			// SetTable("restaurants").
			FieldJoin(types.Join{
				Table:     "categories",  // The table to join with
				Field:     "category_id", // The foreign key in current table
				JoinField: "id",          // The primary key in joined table
			}).
			// Tell GoAdmin which column specifically to display from the joined table
			FieldDisplay(func(value types.FieldModel) interface{} {
				// 1. Check if the joined value is nil
				if value.Row["categories_goadmin_join_name"] == nil {
					return "No Restaurant" // Return a string or empty template.HTML
				}
				// 2. Safely return the value
				return value.Row["categories_goadmin_join_name"]
			}).
			FieldSortable()
		info.AddField("Category Name", "name", db.Varchar).FieldSortable()
		info.AddField("Category Slug", "slug", db.Varchar).FieldSortable()
		info.AddField("Price", "price", db.Float).FieldSortable()
		info.AddField("Discount Price", "discount_price", db.Float).FieldSortable()
		info.AddField("Is Vegetarian", "is_vegetarian", db.Bool).
			FieldDisplay(func(value types.FieldModel) interface{} {
				if value.Row["is_vegetarian"] == true {
					return "Yes" // Return a string or empty template.HTML
				}
				return "No"
			})
		info.AddField("Is Vegan", "is_vegan", db.Bool).
			FieldDisplay(func(value types.FieldModel) interface{} {
				if value.Row["is_vegan"] == true {
					return "Yes" // Return a string or empty template.HTML
				}
				return "No"
			})
		info.AddField("Is Glueten Free", "is_glueten_free", db.Bool).
			FieldDisplay(func(value types.FieldModel) interface{} {
				if value.Row["is_glueten_free"] == true {
					return "Yes" // Return a string or empty template.HTML
				}
				return "No"
			})
		info.AddField("Is Spicy", "is_spicy", db.Bool).
			FieldDisplay(func(value types.FieldModel) interface{} {
				if value.Row["is_spicy"] == true {
					return "Yes" // Return a string or empty template.HTML
				}
				return "No"
			})
		info.AddField("Spice Level", "spice_level", db.Int).FieldSortable()
		info.AddField("Calories", "calories", db.Int).FieldSortable()
		info.AddField("Preparation Time", "preparation_time", db.Int).FieldSortable()
		info.AddField("Is Featured", "is_featured", db.Bool).
			FieldDisplay(func(value types.FieldModel) interface{} {
				if value.Row["is_featured"] == true {
					return "Yes" // Return a string or empty template.HTML
				}
				return "No"
			})
		info.AddField("Is Available", "is_available", db.Bool).
			FieldDisplay(func(value types.FieldModel) interface{} {
				if value.Row["is_available"] == true {
					return "Yes" // Return a string or empty template.HTML
				}
				return "No"
			})
		info.AddField("Stock Quantity", "stock_quantity", db.Int).FieldSortable()
		info.AddField("Low Stock Threshold", "low_stock_threshold", db.Int).FieldSortable()
		info.AddField("Tags", "tags", db.Text)
		info.AddField("Alergens", "alergens", db.Text)
		info.AddField("Image", "image_url", db.Varchar).FieldImage("50", "50",
			"http://localhost:8081/uploads/")
		info.SetSortField("created_at").SetSortDesc()
		info.SetTable("menu_items").SetTitle("Menuitems").SetDescription("Menuitems")

		formList := menuItems.GetForm()

		formList.AddField("Restaurant Name", "restaurant_id", db.Varchar, form.SelectSingle).
			FieldOptionsFromTable("restaurants", "name", "id").
			FieldOnChooseAjax("category_id", "/admin/api/categories",
				func(ctx *context.Context) (bool, string, interface{}) {
					var opts selection.Options
					restaurantID := ctx.FormValue("value")

					if restaurantID == "" {
						return false, restaurantID, nil
					}

					res, err := db.WithDriver(dbConn).
						Table("categories").
						Where("restaurant_id", "=", restaurantID).
						Select("id", "name").
						All()
					if err != nil {
						fmt.Println("error message from db", err)
						return false, restaurantID, nil
					}
					for _, v := range res {
						opts = append(opts, selection.Option{
							ID:   fmt.Sprint(v["id"]),
							Text: fmt.Sprint(v["name"]),
						})
					}

					return true, "ok", opts
				}).
			FieldMust()
		formList.AddField("Category Name", "category_id", db.Varchar, form.SelectSingle).
			FieldMust()

		formList.AddField("Name", "name", db.Varchar, form.Text)
		formList.AddField("SLug", "slug", db.Varchar, form.Text)
		formList.AddField("Image URL", "image_url", db.Varchar, form.File)
		formList.AddField("Description", "description", db.Varchar, form.TextArea)
		formList.AddField("Is Vegetarian", "is_vegetarian", db.Bool, form.Radio).
			FieldOptions(types.FieldOptions{
				{Text: "Active", Value: "true"},
				{Text: "InActive", Value: "false"},
			}).
			FieldDefault("false")
		formList.AddField("Is Vegan", "is_vegan", db.Bool, form.Radio).
			FieldOptions(types.FieldOptions{
				{Text: "Active", Value: "true"},
				{Text: "InActive", Value: "false"},
			}).
			FieldDefault("false")
		formList.AddField("Is Glueten Free", "is_glueten_free", db.Bool, form.Radio).
			FieldOptions(types.FieldOptions{
				{Text: "Active", Value: "true"},
				{Text: "InActive", Value: "false"},
			}).
			FieldDefault("false")
		formList.AddField("Is Spicy", "is_spicy", db.Bool, form.Radio).
			FieldOptions(types.FieldOptions{
				{Text: "Active", Value: "true"},
				{Text: "InActive", Value: "false"},
			}).
			FieldDefault("false")
		formList.AddField("Spice Level", "spice_level", db.Int, form.Number).
			FieldDefault("10")
		formList.AddField("Calories", "calories", db.Int, form.Number).
			FieldDefault("10")
		formList.AddField("Preparation Time", "preparation_time", db.Int, form.Number).
			FieldDefault("10")
		formList.AddField("Is Featured", "is_featured", db.Bool, form.Radio).
			FieldOptions(types.FieldOptions{
				{Text: "Active", Value: "true"},
				{Text: "InActive", Value: "false"},
			}).
			FieldDefault("false")
		formList.AddField("Is Available", "is_available", db.Bool, form.Radio).
			FieldOptions(types.FieldOptions{
				{Text: "Active", Value: "true"},
				{Text: "InActive", Value: "false"},
			}).
			FieldDefault("false")
		formList.AddField("Stock Quantity", "stock_quantity", db.Int, form.Number).
			FieldDefault("10")
		formList.AddField("Preparation Time", "preparation_time", db.Int, form.Number).
			FieldDefault("10")
		formList.AddField("Low Stock Threshold", "low_stock_threshold", db.Int, form.Number).
			FieldDefault("10")
		formList.AddField("Tags", "tags", db.Varchar, form.Text)
		formList.AddField("Alergens", "alergens", db.Varchar, form.Text)
		formList.AddField("Display Order", "display_order", db.Int, form.Number).
			FieldDefault("10")

		formList.SetTable("menu_items").SetTitle("Menuitems").SetDescription("Menuitems")

		return menuItems
	}
}
