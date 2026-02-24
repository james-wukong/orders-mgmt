package tables

import (
	"fmt"
	"time"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetRestaurantsTable(ctx *context.Context) table.Table {

	restaurants := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

	info := restaurants.GetInfo().SetPrimaryKey("id", db.UUID)

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
	info.AddField("Opening_time", "opening_time", db.Time).
		FieldDisplay(func(value types.FieldModel) interface{} {
			t, err := time.Parse(time.RFC3339, value.Value)
			if err != nil {
				return err
			}
			return t.Format("15:04:05")
		})
	info.AddField("Closing_time", "closing_time", db.Time).
		FieldDisplay(func(value types.FieldModel) interface{} {
			t, err := time.Parse(time.RFC3339, value.Value)
			if err != nil {
				return err
			}
			return t.Format("15:04:05")
		})
	info.AddField("Is_open", "is_open", db.Bool).FieldBool("true", "false")
	info.AddField("Delivery_fee", "delivery_fee", db.Numeric)
	info.AddField("Minimum_order", "minimum_order", db.Int4)
	info.AddField("Estimated_delivery_time", "estimated_delivery_time", db.Interval)
	info.AddField("Rating", "rating", db.Numeric)
	info.AddField("Total_reviews", "total_reviews", db.Int4)
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Slug", "slug", db.Varchar)
	info.AddField("Description", "description", db.Text)
	info.AddField("Phone", "phone", db.Varchar)
	info.AddField("Email", "email", db.Varchar)
	info.AddField("Address", "address", db.Text)
	info.AddField("City", "city", db.Varchar)
	info.AddField("State", "state", db.Varchar)
	info.AddField("Postal_code", "postal_code", db.Varchar)
	info.AddField("Country", "country", db.Varchar)
	info.AddField("Logo_url", "logo_url", db.Text)
	info.AddField("Banner_url", "banner_url", db.Text)
	info.AddField("Latitude", "latitude", db.Numeric)
	info.AddField("Longitude", "longitude", db.Numeric)
	info.AddField("Created_at", "created_at", db.Timestamp)
	info.AddField("Updated_at", "updated_at", db.Timestamp)

	info.SetTable("restaurants").SetTitle("Restaurants").SetDescription("Restaurants")

	formList := restaurants.GetForm().SetPrimaryKey("id", db.UUID)
	formList.AddField("Id", "id", db.UUID, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("Name", "name", db.Varchar, form.Text).FieldMust()
	formList.AddField("Slug", "slug", db.Varchar, form.Text).FieldMust()
	formList.AddField("Opening_time", "opening_time", db.Time, form.Text).
		FieldPlaceholder("Select Time").
		FieldHelpMsg("Format: HH:MM:SS").
		FieldMust()
	formList.AddField("Closing_time", "closing_time", db.Time, form.Text).
		FieldPlaceholder("Select Time").
		FieldHelpMsg("Format: HH:MM:SS").
		FieldMust()
	formList.AddField("Is_open", "is_open", db.Bool, form.Switch).
		FieldOptions(types.FieldOptions{
			{Text: "Open", Value: "true"},
			{Text: "Closed", Value: "false"},
		}).
		FieldDefault("false")
	formList.AddField("Delivery_fee", "delivery_fee", db.Decimal, form.Currency).
		FieldDefault("10.00").
		FieldMust()
	formList.AddField("Minimum_order", "minimum_order", db.Decimal, form.Currency).
		FieldDefault("1.00").
		FieldMust()
	formList.AddField("Estimated_delivery_time", "estimated_delivery_time", db.Varchar, form.Text).
		FieldHelpMsg("Format: 00:30:00 for 30 minutes").
		FieldMust()
	formList.AddField("Rating", "rating", db.Numeric, form.Number).
		FieldDefault("2").
		FieldMust()
	formList.AddField("Total_reviews", "total_reviews", db.Int4, form.Number).
		FieldDefault("10").
		FieldMust()
	formList.AddField("Description", "description", db.Text, form.RichText)
	formList.AddField("Phone", "phone", db.Varchar, form.Text)
	formList.AddField("Email", "email", db.Varchar, form.Email)
	formList.AddField("Address", "address", db.Text, form.Text)
	formList.AddField("City", "city", db.Varchar, form.Text)
	formList.AddField("State", "state", db.Varchar, form.Text)
	formList.AddField("Postal_code", "postal_code", db.Varchar, form.Text)
	formList.AddField("Country", "country", db.Varchar, form.Text)
	formList.AddField("Logo_url", "logo_url", db.Text, form.File)
	formList.AddField("Banner_url", "banner_url", db.Text, form.File)
	formList.AddField("Latitude", "latitude", db.Numeric, form.Number).
		FieldDefault("0").
		FieldMust()
	formList.AddField("Longitude", "longitude", db.Numeric, form.Number).
		FieldDefault("0").
		FieldMust()
	formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime).
		FieldDefault(time.Now().Format("2006-01-02 15:04:05")). // Set initial value
		FieldHide().FieldNowWhenInsert()
	formList.AddField("Updated_at", "updated_at", db.Timestamp, form.Datetime).
		FieldDefault(time.Now().Format("2006-01-02 15:04:05")). // Set initial value
		FieldHide().FieldNowWhenUpdate()

	formList.SetPostValidator(func(values form2.Values) error {
		for k, v := range values {
			fmt.Println("values k: ", k, ", v: ", v)
		}
		return nil
	})
	formList.SetTable("restaurants").SetTitle("Restaurants").SetDescription("Restaurants")

	return restaurants
}
