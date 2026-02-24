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

func GetDeliveryaddressesTable(ctx *context.Context) table.Table {

	deliveryAddresses := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

	info := deliveryAddresses.GetInfo()

	info.AddField("Id", "id", db.UUID)
	info.AddField("User_id", "user_id", db.UUID)
	info.AddField("Is_default", "is_default", db.Bool).FieldBool("true", "false")
	info.AddField("Postal_code", "postal_code", db.Varchar)
	info.AddField("Country", "country", db.Varchar)
	info.AddField("Phone", "phone", db.Varchar)
	info.AddField("Delivery_instructions", "delivery_instructions", db.Text)
	info.AddField("Label", "label", db.Varchar)
	info.AddField("Street_address", "street_address", db.Text)
	info.AddField("Apartment", "apartment", db.Varchar)
	info.AddField("City", "city", db.Varchar)
	info.AddField("State", "state", db.Varchar)
	info.AddField("Latitude", "latitude", db.Numeric)
	info.AddField("Longitude", "longitude", db.Numeric)
	info.AddField("Created_at", "created_at", db.Timestamp)
	// info.AddField("Updated_at", "updated_at", db.Timestamp)

	info.SetTable("delivery_addresses").SetTitle("Deliveryaddresses").SetDescription("Deliveryaddresses")

	formList := deliveryAddresses.GetForm()
	info.AddField("Id", "id", db.UUID).
		FieldDisplay(func(value types.FieldModel) interface{} {
			if val, ok := value.Row["id"].(string); ok {
				return val
			}
			if bytes, ok := value.Row["id"].([]byte); ok {
				return string(bytes)
			}
			return fmt.Sprintf("%s", value.Row["id"])
		}).
		FieldHide()
	formList.AddField("User_id", "user_id", db.UUID, form.Text)
	formList.AddField("Is_default", "is_default", db.Bool, form.Switch).
		FieldOptions(types.FieldOptions{
			{Text: "Default", Value: "true"},
			{Text: "Other", Value: "false"},
		}).
		FieldDefault("false")
	formList.AddField("Postal_code", "postal_code", db.Varchar, form.Text)
	formList.AddField("Country", "country", db.Varchar, form.Text)
	formList.AddField("Phone", "phone", db.Varchar, form.Text)
	formList.AddField("Label", "label", db.Varchar, form.Text)
	formList.AddField("Street_address", "street_address", db.Text, form.Text)
	formList.AddField("Apartment", "apartment", db.Varchar, form.Text)
	formList.AddField("City", "city", db.Varchar, form.Text)
	formList.AddField("State", "state", db.Varchar, form.Text)
	formList.AddField("Latitude", "latitude", db.Decimal, form.Number).FieldDefault("0.00")
	formList.AddField("Longitude", "longitude", db.Decimal, form.Number).FieldDefault("0.00")
	formList.AddField("Delivery_instructions", "delivery_instructions", db.Text, form.RichText)
	formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime).
		FieldDefault(time.Now().Format("2006-01-02 15:04:05")). // Set initial value
		FieldHide().FieldNowWhenInsert()
	formList.AddField("Updated_at", "updated_at", db.Timestamp, form.Datetime).
		FieldDefault(time.Now().Format("2006-01-02 15:04:05")). // Set initial value
		FieldHide().FieldNowWhenUpdate()

	formList.SetTable("delivery_addresses").
		SetTitle("Deliveryaddresses").
		SetDescription("Deliveryaddresses")

	return deliveryAddresses
}
