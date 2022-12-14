// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CouponsColumns holds the columns for the "coupons" table.
	CouponsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "delete_time", Type: field.TypeTime, Nullable: true},
		{Name: "title", Type: field.TypeString},
		{Name: "start_time", Type: field.TypeTime},
		{Name: "end_time", Type: field.TypeTime},
		{Name: "description", Type: field.TypeString},
		{Name: "full_money", Type: field.TypeFloat64},
		{Name: "minus", Type: field.TypeFloat64},
		{Name: "rate", Type: field.TypeFloat64},
		{Name: "type", Type: field.TypeInt},
		{Name: "valitiy", Type: field.TypeInt},
		{Name: "activity_id", Type: field.TypeInt64, Nullable: true},
		{Name: "remark", Type: field.TypeString},
		{Name: "whole_store", Type: field.TypeInt},
	}
	// CouponsTable holds the schema information for the "coupons" table.
	CouponsTable = &schema.Table{
		Name:       "coupons",
		Columns:    CouponsColumns,
		PrimaryKey: []*schema.Column{CouponsColumns[0]},
	}
	// CouponTemplatesColumns holds the columns for the "coupon_templates" table.
	CouponTemplatesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "delete_time", Type: field.TypeTime, Nullable: true},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "full_money", Type: field.TypeFloat64},
		{Name: "minus", Type: field.TypeFloat64},
		{Name: "discount", Type: field.TypeFloat64},
		{Name: "type", Type: field.TypeInt},
	}
	// CouponTemplatesTable holds the schema information for the "coupon_templates" table.
	CouponTemplatesTable = &schema.Table{
		Name:       "coupon_templates",
		Columns:    CouponTemplatesColumns,
		PrimaryKey: []*schema.Column{CouponTemplatesColumns[0]},
	}
	// CouponTypesColumns holds the columns for the "coupon_types" table.
	CouponTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "delete_time", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "code", Type: field.TypeInt},
		{Name: "description", Type: field.TypeString},
	}
	// CouponTypesTable holds the schema information for the "coupon_types" table.
	CouponTypesTable = &schema.Table{
		Name:       "coupon_types",
		Columns:    CouponTypesColumns,
		PrimaryKey: []*schema.Column{CouponTypesColumns[0]},
	}
	// UserCouponColumns holds the columns for the "user_coupon" table.
	UserCouponColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "delete_time", Type: field.TypeTime, Nullable: true},
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "coupon_id", Type: field.TypeInt64, Nullable: true},
		{Name: "status", Type: field.TypeInt, Default: 1},
		{Name: "order_id", Type: field.TypeInt},
	}
	// UserCouponTable holds the schema information for the "user_coupon" table.
	UserCouponTable = &schema.Table{
		Name:       "user_coupon",
		Columns:    UserCouponColumns,
		PrimaryKey: []*schema.Column{UserCouponColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CouponsTable,
		CouponTemplatesTable,
		CouponTypesTable,
		UserCouponTable,
	}
)

func init() {
	UserCouponTable.Annotation = &entsql.Annotation{
		Table: "user_coupon",
	}
}
