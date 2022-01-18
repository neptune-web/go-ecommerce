// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// OrderColumns holds the columns for the "order" table.
	OrderColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "delete_time", Type: field.TypeTime, Nullable: true},
		{Name: "order_no", Type: field.TypeString, Size: 64},
		{Name: "transaction_id", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeInt64, Nullable: true},
		{Name: "total_price", Type: field.TypeFloat64},
		{Name: "total_count", Type: field.TypeInt},
		{Name: "final_total_price", Type: field.TypeFloat64},
		{Name: "status", Type: field.TypeInt},
	}
	// OrderTable holds the schema information for the "order" table.
	OrderTable = &schema.Table{
		Name:       "order",
		Columns:    OrderColumns,
		PrimaryKey: []*schema.Column{OrderColumns[0]},
	}
	// OrderDetailColumns holds the columns for the "order_detail" table.
	OrderDetailColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "delete_time", Type: field.TypeTime, Nullable: true},
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "pay_way", Type: field.TypeInt, Default: 1},
		{Name: "client_type", Type: field.TypeInt, Default: 1},
		{Name: "ship_no", Type: field.TypeString},
	}
	// OrderDetailTable holds the schema information for the "order_detail" table.
	OrderDetailTable = &schema.Table{
		Name:       "order_detail",
		Columns:    OrderDetailColumns,
		PrimaryKey: []*schema.Column{OrderDetailColumns[0]},
	}
	// OrderSnapColumns holds the columns for the "order_snap" table.
	OrderSnapColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "snap_img", Type: field.TypeString},
		{Name: "snap_title", Type: field.TypeString},
		{Name: "snap_items", Type: field.TypeString},
		{Name: "snap_address", Type: field.TypeString},
		{Name: "order_id", Type: field.TypeInt64, Nullable: true},
	}
	// OrderSnapTable holds the schema information for the "order_snap" table.
	OrderSnapTable = &schema.Table{
		Name:       "order_snap",
		Columns:    OrderSnapColumns,
		PrimaryKey: []*schema.Column{OrderSnapColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "order_snap_order_order_snap",
				Columns:    []*schema.Column{OrderSnapColumns[5]},
				RefColumns: []*schema.Column{OrderColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OrderSubColumns holds the columns for the "order_sub" table.
	OrderSubColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "delete_time", Type: field.TypeTime, Nullable: true},
		{Name: "order_no", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeInt64, Nullable: true},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "count", Type: field.TypeInt},
		{Name: "final_price", Type: field.TypeFloat64},
		{Name: "status", Type: field.TypeInt},
		{Name: "order_id", Type: field.TypeInt64, Nullable: true},
	}
	// OrderSubTable holds the schema information for the "order_sub" table.
	OrderSubTable = &schema.Table{
		Name:       "order_sub",
		Columns:    OrderSubColumns,
		PrimaryKey: []*schema.Column{OrderSubColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "order_sub_order_order_sub",
				Columns:    []*schema.Column{OrderSubColumns[10]},
				RefColumns: []*schema.Column{OrderColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		OrderTable,
		OrderDetailTable,
		OrderSnapTable,
		OrderSubTable,
	}
)

func init() {
	OrderSnapTable.ForeignKeys[0].RefTable = OrderTable
	OrderSubTable.ForeignKeys[0].RefTable = OrderTable
}
