package tables

import (
	"time"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetSuppliersTable(ctx *context.Context) table.Table {

	suppliers := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

	info := suppliers.GetInfo().SetPrimaryKey("id", db.UUID)

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
	info.AddField("Contact_person", "contact_person", db.Varchar)
	info.AddField("Email", "email", db.Varchar)
	info.AddField("Phone", "phone", db.Varchar)
	info.AddField("Address", "address", db.Text)
	info.AddField("Minimum_order_amount", "minimum_order_amount", db.Numeric)
	info.AddField("Is_active", "is_active", db.Bool).FieldBool("true", "false")
	info.AddField("Rating", "rating", db.Numeric)
	info.AddField("City", "city", db.Varchar)
	info.AddField("State", "state", db.Varchar)
	info.AddField("Postal_code", "postal_code", db.Varchar)
	info.AddField("Country", "country", db.Varchar)
	info.AddField("Payment_terms", "payment_terms", db.Varchar)
	info.AddField("Delivery_days", "delivery_days", db.Varchar)

	info.AddField("Notes", "notes", db.Text)
	info.AddField("Created_at", "created_at", db.Timestamp)
	info.AddField("Updated_at", "updated_at", db.Timestamp)

	info.SetTable("suppliers").SetTitle("Suppliers").SetDescription("Suppliers")

	formList := suppliers.GetForm().SetPrimaryKey("id", db.UUID)
	formList.AddField("Id", "id", db.UUID, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("Name", "name", db.Varchar, form.Text).FieldMust()
	formList.AddField("Contact_person", "contact_person", db.Varchar, form.Text).
		FieldMust()
	formList.AddField("Email", "email", db.Varchar, form.Email)
	formList.AddField("Phone", "phone", db.Varchar, form.Text)
	formList.AddField("Address", "address", db.Text, form.Text)
	formList.AddField("Minimum_order_amount", "minimum_order_amount", db.Decimal, form.Currency).
		FieldDefault("10.00").
		FieldMust()
	formList.AddField("Is_active", "is_active", db.Bool, form.Switch).
		FieldOptions(types.FieldOptions{
			{Text: "Active", Value: "true"},
			{Text: "Inactive", Value: "false"},
		}).
		FieldDefault("false")
	formList.AddField("Rating", "rating", db.Decimal, form.Number).
		FieldDefault("2").
		FieldMust()
	formList.AddField("City", "city", db.Varchar, form.Text)
	formList.AddField("State", "state", db.Varchar, form.Text)
	formList.AddField("Postal_code", "postal_code", db.Varchar, form.Text)
	formList.AddField("Country", "country", db.Varchar, form.Text)
	formList.AddField("Payment_terms", "payment_terms", db.Varchar, form.Text)
	formList.AddField("Delivery_days", "delivery_days", db.Varchar, form.Text)

	formList.AddField("Notes", "notes", db.Text, form.RichText)
	formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime).
		FieldDefault(time.Now().Format("2006-01-02 15:04:05")). // Set initial value
		FieldHide().FieldNowWhenInsert()
	formList.AddField("Updated_at", "updated_at", db.Timestamp, form.Datetime).
		FieldDefault(time.Now().Format("2006-01-02 15:04:05")). // Set initial value
		FieldHide().FieldNowWhenUpdate()

	formList.SetTable("suppliers").SetTitle("Suppliers").SetDescription("Suppliers")

	return suppliers
}
