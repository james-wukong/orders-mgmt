package tables

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/GoAdminGroup/go-admin/context"    // 需导入此包
	"github.com/GoAdminGroup/go-admin/modules/db" // 需导入此包

	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	selection "github.com/GoAdminGroup/go-admin/template/types/form/select"
	models2 "github.com/james-wukong/orders-mgmt/internal/models"
)

func GetMenuitemsTable(dbConn db.Connection) table.Generator {
	return func(ctx *context.Context) table.Table {

		menuItems := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

		info := menuItems.GetInfo().HideFilterArea().SetPrimaryKey("id", db.UUID)
		info.AddField("ID", "id", db.UUID).
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
		info.AddField("Is Vegetarian", "is_vegetarian", db.Bool).FieldBool("true", "false")
		info.AddField("Is Vegan", "is_vegan", db.Bool).FieldBool("true", "false")
		info.AddField("Is Gluten Free", "is_gluten_free", db.Bool).FieldBool("true", "false")
		info.AddField("Is Spicy", "is_spicy", db.Bool).FieldBool("true", "false")
		info.AddField("Spice Level", "spice_level", db.Int).FieldSortable()
		info.AddField("Calories", "calories", db.Int).FieldSortable()
		info.AddField("Preparation Time", "preparation_time", db.Interval)
		info.AddField("Is Featured", "is_featured", db.Bool).FieldBool("true", "false")
		info.AddField("Is Available", "is_available", db.Bool).FieldBool("true", "false")
		info.AddField("Stock Quantity", "stock_quantity", db.Int).FieldSortable()
		info.AddField("Low Stock Threshold", "low_stock_threshold", db.Int).FieldSortable()
		info.AddField("Tags", "tags", db.Text)
		info.AddField("Alergens", "alergens", db.Text)
		info.AddField("Image", "image_url", db.Varchar).FieldImage("50", "50",
			"http://localhost:8081/uploads/")
		info.SetSortField("created_at").SetSortDesc()
		info.SetTable("menu_items").SetTitle("Menu Items").SetDescription("Menu Items")

		formList := menuItems.GetForm().SetPrimaryKey("id", db.UUID)
		formList.AddField("Id", "id", db.UUID, form.Default).
			FieldDisableWhenCreate()
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
			FieldOptionInitFn(func(val types.FieldModel) types.FieldOptions {
				var c types.FieldOptions
				categories, err := db.WithDriver(dbConn).Table("categories").
					Where("restaurant_id", "=", val.Row["restaurant_id"]).
					Select("id", "name").
					All()
				if err != nil || len(categories) == 0 {
					return nil
				}
				for _, v := range categories {
					opt := types.FieldOption{
						Text:  v["name"].(string),
						Value: v["id"].(string)}
					if v["id"].(string) == val.Row["category_id"].(string) {
						opt.Selected = true
					}
					c = append(c, opt)
				}

				return c
			}).
			FieldMust()

		formList.AddField("Name", "name", db.Varchar, form.Text)
		formList.AddField("SLug", "slug", db.Varchar, form.Text)
		formList.AddField("Price", "price", db.Decimal, form.Currency).
			FieldDefault("10")
		formList.AddField("Discount Price", "discount_price", db.Decimal, form.Currency).
			FieldDefault("8")
		formList.AddField("Image URL", "image_url", db.Varchar, form.File).
			FieldPostFilterFn(func(value types.PostFieldModel) interface{} {
				// value.Value contains the path GoAdmin generated
				fmt.Println("File saved at:", value.Value)
				return value
			})
		formList.AddField("Description", "description", db.Varchar, form.RichText)
		formList.AddField("Is Vegetarian", "is_vegetarian", db.Bool, form.Switch).
			FieldOptions(types.FieldOptions{
				{Text: "Yes", Value: "true"},
				{Text: "No", Value: "false"},
			}).
			FieldDefault("false")
		formList.AddField("Is Vegan", "is_vegan", db.Bool, form.Switch).
			FieldOptions(types.FieldOptions{
				{Text: "Yes", Value: "true"},
				{Text: "No", Value: "false"},
			}).
			FieldDefault("false")
		formList.AddField("Is Gluten Free", "is_gluten_free", db.Bool, form.Switch).
			FieldOptions(types.FieldOptions{
				{Text: "Yes", Value: "true"},
				{Text: "No", Value: "false"},
			}).
			FieldDefault("false")
		formList.AddField("Is Spicy", "is_spicy", db.Bool, form.Switch).
			FieldOptions(types.FieldOptions{
				{Text: "Yes", Value: "true"},
				{Text: "No", Value: "false"},
			}).
			FieldDefault("false")
		formList.AddField("Spice Level", "spice_level", db.Int, form.Number).
			// FieldOptionExt(map[string]interface{}{
			// 	"min":  1,
			// 	"max":  5,
			// 	"step": 1,
			// }).
			FieldMust().
			FieldDefault("1").
			FieldHelpMsg("Between 1 - 5 (Include 1 and 5)")
		formList.AddField("Calories", "calories", db.Int, form.Number).
			FieldDefault("10")
		formList.AddField("Preparation Time", "preparation_time", db.Varchar, form.Text).
			FieldHelpMsg("Format: 00:10:00 for 10 minutes")
		formList.AddField("Is Featured", "is_featured", db.Bool, form.Switch).
			FieldOptions(types.FieldOptions{
				{Text: "Yes", Value: "true"},
				{Text: "No", Value: "false"},
			}).
			FieldDefault("false")
		formList.AddField("Is Available", "is_available", db.Bool, form.Switch).
			FieldOptions(types.FieldOptions{
				{Text: "Yes", Value: "true"},
				{Text: "No", Value: "false"},
			}).
			FieldDefault("false")
		formList.AddField("Stock Quantity", "stock_quantity", db.Int, form.Number).
			FieldDefault("10")
		formList.AddField("Low Stock Threshold", "low_stock_threshold", db.Int, form.Number).
			FieldDefault("10")
		// TODO: multi-select for tags and alergens
		formList.AddField("Tags", "tags", db.Varchar, form.Select).
			// 1. Enable "Tags" mode so Enter/Comma creates a new bubble
			FieldOptionExt(map[string]interface{}{
				"tags":            true,
				"tokenSeparators": []string{",", " "},
			}).
			// 2. Clean the data when loading from DB (stripping PostgreSQL {} braces)
			FieldOptionInitFn(func(val types.FieldModel) types.FieldOptions {
				// 编辑时的显示，根据行数据 val 返回options
				var c types.FieldOptions
				tags, _ := db.WithDriver(dbConn).Table("menu_item_tags").
					LeftJoin("tags", "tags.id", "=", "menu_item_tags.tag_id").
					Where("menu_item_tags.menu_item_id", "=", val.Row["id"]).
					Select("id", "name").
					All()
				if len(tags) == 0 {
					return nil
				}
				for _, v := range tags {
					opt := types.FieldOption{
						Text:     v["name"].(string),
						Value:    v["name"].(string),
						Selected: true,
					}
					c = append(c, opt)
				}

				return c
			}).
			FieldPlaceholder("Type and press Enter")

		formList.AddField("Allergens", "allergens", db.Varchar, form.Select).
			FieldOptionExt(map[string]interface{}{
				"tags":            true,
				"tokenSeparators": []string{",", " "},
			}).
			FieldOptionInitFn(func(val types.FieldModel) types.FieldOptions {
				// 编辑时的显示，根据行数据 val 返回options
				var c types.FieldOptions
				allergens, _ := db.WithDriver(dbConn).Table("menu_item_allergens").
					LeftJoin("allergens", "allergens.id", "=", "menu_item_allergens.allergen_id").
					Where("menu_item_allergens.menu_item_id", "=", val.Row["id"]).
					Select("id", "name").
					All()
				if len(allergens) == 0 {
					return nil
				}
				for _, v := range allergens {
					opt := types.FieldOption{
						Text:     v["name"].(string),
						Value:    v["name"].(string),
						Selected: true,
					}
					c = append(c, opt)
				}

				return c
			}).
			FieldPlaceholder("Type and press Enter")
		formList.AddField("Display Order", "display_order", db.Int, form.Number).
			FieldDefault("10")
		formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime).
			FieldDefault(time.Now().Format("2006-01-02 15:04:05")). // Set initial value
			FieldHide().FieldNowWhenInsert()
		formList.AddField("Updated_at", "updated_at", db.Timestamp, form.Datetime).
			FieldDefault(time.Now().Format("2006-01-02 15:04:05")). // Set initial value
			FieldHide().FieldNowWhenUpdate()

		// rewrite update function
		formList.SetUpdateFn(func(values form2.Values) error {
			// 1. validate input
			if values.IsEmpty("name", "slug") {
				return errors.New("name and slug can not be empty")
			}
			values.RemoveSysRemark()
			tags := values["tags[]"]
			values.Delete("tags[]")
			allergens := values["allergens[]"]
			values.Delete("allergens[]")

			// 2. start transaction
			_, txErr := db.WithDriver(dbConn).
				WithTransaction(func(tx *sql.Tx) (error, map[string]interface{}) {
					tagModel := models2.NewTag(dbConn)
					allergenModel := models2.NewAllergen(dbConn)
					menuItemModel := models2.NewMenuItem(dbConn)
					menuItemTagModel := models2.NewMenuItemTag(dbConn)
					menuItemAllergenModel := models2.NewMenuItemAllergen(dbConn)
					// 2.1 update menu items with new data
					_ = menuItemModel.UpdateMenuItem(values, tx)

					// 2.2 Insert tags and menu item tags to database
					if len(tags) > 0 {
						tagIDs := []interface{}{}
						for _, v := range tags {
							tagID, err := tagModel.UpsertTag(v, tx)
							if err == nil {
								_ = menuItemTagModel.InsertMenuItemTag(values.Get("id"), tagID, tx)
							}
							tagIDs = append(tagIDs, tagID)

						}
						// 2.4 Remove old records in composite table
						menuItemTagModel.RemoveMenuItemTags(values.Get("id"), tagIDs, tx)
					}
					// 2.3 Insert allergens and menu item tags to database
					if len(allergens) > 0 {
						allergenIDs := []interface{}{}
						for _, v := range allergens {
							allergenID, err := allergenModel.UpsertAllergen(v, tx)
							if err == nil {
								_ = menuItemAllergenModel.InsertMenuItemAllergen(
									values.Get("id"),
									allergenID,
									tx,
								)
							}
							allergenIDs = append(allergenIDs, allergenID)
						}
						// 2.4 Remove old records in composite table
						menuItemAllergenModel.RemoveMenuItemAllergens(values.Get("id"), allergenIDs, tx)
					}
					return nil, nil
				})
			fmt.Println("itxErr : ", txErr)
			return txErr
		})

		// TODO: update insert function
		formList.SetInsertFn(func(values form2.Values) error {
			// 1. validate input
			if values.IsEmpty("name", "slug") {
				return errors.New("name and slug can not be empty")
			}
			values.RemoveSysRemark()
			tags := values["tags[]"]
			values.Delete("tags[]")
			allergens := values["allergens[]"]
			values.Delete("allergens[]")

			// 2. start transaction
			_, txErr := db.WithDriver(dbConn).
				WithTransaction(func(tx *sql.Tx) (error, map[string]interface{}) {
					tagModel := models2.NewTag(dbConn)
					allergenModel := models2.NewAllergen(dbConn)
					menuItemModel := models2.NewMenuItem(dbConn)
					menuItemTagModel := models2.NewMenuItemTag(dbConn)
					menuItemAllergenModel := models2.NewMenuItemAllergen(dbConn)
					// 2.1 update menu items with new data
					menuItemID, _ := menuItemModel.CreateMenuItem(values, tx)

					// 2.2 Insert tags and menu item tags to database
					if len(tags) > 0 {
						for _, v := range tags {
							tagID, err := tagModel.UpsertTag(v, tx)
							if err == nil {
								_ = menuItemTagModel.InsertMenuItemTag(menuItemID, tagID, tx)
							}

						}
					}
					// 2.3 Insert allergens and menu item tags to database
					if len(allergens) > 0 {
						for _, v := range allergens {
							allergenID, err := allergenModel.UpsertAllergen(v, tx)
							if err == nil {
								_ = menuItemAllergenModel.InsertMenuItemAllergen(
									menuItemID,
									allergenID,
									tx,
								)
							}
						}
					}
					return nil, nil
				})
			fmt.Println("itxErr : ", txErr)
			return txErr
		})

		formList.SetTable("menu_items").SetTitle("Menu Items").SetDescription("Menu Items")

		return menuItems
	}
}
