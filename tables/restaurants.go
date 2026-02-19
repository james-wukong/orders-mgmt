package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetRestaurantsTable(ctx *context.Context) table.Table {

	restaurants := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

	info := restaurants.GetInfo()

	info.AddField("Updated_at", "updated_at", db.Timestamp)
	info.AddField("Latitude", "latitude", db.Numeric)
	info.AddField("Longitude", "longitude", db.Numeric)
	info.AddField("Opening_time", "opening_time", db.Time)
	info.AddField("Closing_time", "closing_time", db.Time)
	info.AddField("Is_open", "is_open", db.Bool)
	info.AddField("Delivery_fee", "delivery_fee", db.Numeric)
	info.AddField("Minimum_order", "minimum_order", db.Numeric)
	info.AddField("Estimated_delivery_time", "estimated_delivery_time", db.Interval)
	info.AddField("Rating", "rating", db.Numeric)
	info.AddField("Total_reviews", "total_reviews", db.Int4)
	info.AddField("Created_at", "created_at", db.Timestamp)
	info.AddField("Id", "id", db.UUID)
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

	info.SetTable("restaurants").SetTitle("Restaurants").SetDescription("Restaurants")

	formList := restaurants.GetForm()
	formList.AddField("Updated_at", "updated_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenUpdate()
	formList.AddField("Latitude", "latitude", db.Numeric, form.Number)
	formList.AddField("Longitude", "longitude", db.Numeric, form.Number)
	formList.AddField("Opening_time", "opening_time", db.Time, form.Datetime)
	formList.AddField("Closing_time", "closing_time", db.Time, form.Datetime)
	formList.AddField("Is_open", "is_open", db.Bool, form.Text)
	formList.AddField("Delivery_fee", "delivery_fee", db.Numeric, form.Number)
	formList.AddField("Minimum_order", "minimum_order", db.Numeric, form.Number)
	formList.AddField("Estimated_delivery_time", "estimated_delivery_time", db.Interval, form.Text)
	formList.AddField("Rating", "rating", db.Numeric, form.Number)
	formList.AddField("Total_reviews", "total_reviews", db.Int4, form.Number)
	formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenInsert()
	formList.AddField("Id", "id", db.UUID, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Slug", "slug", db.Varchar, form.Text)
	formList.AddField("Description", "description", db.Text, form.RichText)
	formList.AddField("Phone", "phone", db.Varchar, form.Text)
	formList.AddField("Email", "email", db.Varchar, form.Email)
	formList.AddField("Address", "address", db.Text, form.RichText)
	formList.AddField("City", "city", db.Varchar, form.Text)
	formList.AddField("State", "state", db.Varchar, form.Text)
	formList.AddField("Postal_code", "postal_code", db.Varchar, form.Text)
	formList.AddField("Country", "country", db.Varchar, form.Text)
	formList.AddField("Logo_url", "logo_url", db.Text, form.RichText)
	formList.AddField("Banner_url", "banner_url", db.Text, form.RichText)

	formList.SetTable("restaurants").SetTitle("Restaurants").SetDescription("Restaurants")

	return restaurants
}
