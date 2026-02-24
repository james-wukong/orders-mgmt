package tables

import (
	"time"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetIngredientcategoriesTable(ctx *context.Context) table.Table {

	ingredientCategories := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

	info := ingredientCategories.GetInfo().SetPrimaryKey("id", db.UUID)

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
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Display_order", "display_order", db.Int4)
	info.AddField("Description", "description", db.Text)
	info.AddField("Created_at", "created_at", db.Timestamp)

	info.SetTable("ingredient_categories").SetTitle("Ingredientcategories").SetDescription("Ingredientcategories")

	formList := ingredientCategories.GetForm().SetPrimaryKey("id", db.UUID)
	formList.AddField("Id", "id", db.UUID, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("Name", "name", db.Varchar, form.Text).FieldMust()
	formList.AddField("Description", "description", db.Text, form.RichText)
	formList.AddField("Display_order", "display_order", db.Int4, form.Number).
		FieldDefault("10").
		FieldMust()
	formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime).
		FieldDefault(time.Now().Format("2006-01-02 15:04:05")). // Set initial value
		FieldHide().FieldNowWhenInsert()

	formList.SetTable("ingredient_categories").SetTitle("Ingredientcategories").SetDescription("Ingredientcategories")

	return ingredientCategories
}
