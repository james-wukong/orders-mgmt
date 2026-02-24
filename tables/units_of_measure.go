package tables

import (
	"time"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetUnitsofmeasureTable(ctx *context.Context) table.Table {

	unitsOfMeasure := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

	info := unitsOfMeasure.GetInfo()

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
	info.AddField("Conversion_factor", "conversion_factor", db.Numeric)
	info.AddField("Base_unit", "base_unit", db.Varchar)
	info.AddField("Type", "type", db.Varchar)
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Abbreviation", "abbreviation", db.Varchar)
	info.AddField("Created_at", "created_at", db.Timestamp)

	info.SetTable("units_of_measure").SetTitle("Unitsofmeasure").SetDescription("Unitsofmeasure")

	formList := unitsOfMeasure.GetForm()
	formList.AddField("Id", "id", db.UUID, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("Conversion_factor", "conversion_factor", db.Numeric, form.Number)
	formList.AddField("Base_unit", "base_unit", db.Varchar, form.Text)
	formList.AddField("Type", "type", db.Varchar, form.Text)
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Abbreviation", "abbreviation", db.Varchar, form.Text)
	formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime).
		FieldDefault(time.Now().Format("2006-01-02 15:04:05")). // Set initial value
		FieldHide().FieldNowWhenInsert()

	formList.SetTable("units_of_measure").SetTitle("Unitsofmeasure").SetDescription("Unitsofmeasure")

	return unitsOfMeasure
}
