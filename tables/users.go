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
	models2 "github.com/james-wukong/orders-mgmt/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func GetUsersTable(dbConn db.Connection) table.Generator {
	return func(ctx *context.Context) table.Table {

		users := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

		info := users.GetInfo().SetPrimaryKey("id", db.UUID).HideFilterArea()

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
		info.AddField("Email", "email", db.Varchar).FieldFilterable()
		info.AddField("First Name", "first_name", db.Varchar).FieldFilterable()
		info.AddField("Last Name", "last_name", db.Varchar).FieldFilterable()
		info.AddField("Role", "role", db.Varchar).
			FieldFilterable(types.FilterType{FormType: form.Select}).
			FieldFilterOptions(types.FieldOptions{
				{Text: "Customer", Value: "customer", Selected: true},
				{Text: "Admin", Value: "admin"},
				{Text: "Kitehen", Value: "kitchen"},
				{Text: "Delivery", Value: "delivery"},
				{Text: "Inventory Manager", Value: "inventory_manager"},
			})
		info.AddField("Active", "is_active", db.Boolean).FieldBool("true", "false")
		info.AddField("Created_at", "created_at", db.Timestamp)
		// info.AddField("Updated_at", "updated_at", db.Timestamp)

		info.SetTable("users").SetTitle("Users").SetDescription("Users")

		formList := users.GetForm().SetPrimaryKey("id", db.UUID)
		formList.AddField("Id", "id", db.UUID, form.Default).
			FieldDisableWhenCreate()
		formList.AddField("Email", "email", db.Varchar, form.Email).
			FieldMust().
			FieldPlaceholder("user@example.com")
		formList.AddField("Password", "password_hash", db.Varchar, form.Password).
			FieldMust().
			FieldPostFilterFn(func(value types.PostFieldModel) interface{} {
				// Automatically hash the password if it's being changed
				if value.Value.Value() == "" {
					return ""
				}
				hash, _ := bcrypt.GenerateFromPassword([]byte(value.Value.Value()), bcrypt.DefaultCost)
				return string(hash)
			})
		formList.AddField("First Name", "first_name", db.Varchar, form.Text)
		formList.AddField("Last Name", "last_name", db.Varchar, form.Text)
		formList.AddField("Phone", "phone", db.Varchar, form.Text)
		// Handling the Enum
		formList.AddField("Role", "role", db.Varchar, form.SelectSingle).
			FieldOptions(types.FieldOptions{
				{Text: "Customer", Value: "customer"},
				{Text: "Admin", Value: "admin"},
				{Text: "Kitehen", Value: "kitchen"},
				{Text: "Delivery", Value: "delivery"},
				{Text: "Inventory Manager", Value: "inventory_manager"},
			}).FieldDefault("customer")

		formList.AddField("Active", "is_active", db.Boolean, form.Switch).
			FieldOptions(types.FieldOptions{
				{Text: "Active", Value: "true"},
				{Text: "InActive", Value: "false"},
			}).
			FieldDefault("true")

		// System/Internal fields - Hidden from form but handled automatically
		formList.AddField("Email Verified", "email_verified", db.Boolean, form.Switch).
			FieldOptions(types.FieldOptions{
				{Text: "Verified", Value: "true"},
				{Text: "Not Verified", Value: "false"},
			}).
			FieldDefault("false")
		// FieldHide()

		formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime).
			FieldDefault(time.Now().Format("2006-01-02 15:04:05")). // Set initial value
			FieldHide().FieldNowWhenInsert()
		formList.AddField("Updated_at", "updated_at", db.Timestamp, form.Datetime).
			FieldDefault(time.Now().Format("2006-01-02 15:04:05")). // Set initial value
			FieldHide().FieldNowWhenUpdate()

		// TODO: input validation
		formList.SetPostValidator(func(values form2.Values) error {
			for k, v := range values {
				fmt.Println("k is :", k, ", v is: ", v)
			}

			return nil
		})
		formList.SetInsertFn(func(values form2.Values) error {
			// values 为传入的表单参数
			userModel := models2.NewUser(dbConn)
			_, err := userModel.CreateUser(values, nil)
			return err
		})

		formList.SetTable("users").SetTitle("Users").SetDescription("Users")

		return users
	}
}
