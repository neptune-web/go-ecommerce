// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ActivitiesColumns holds the columns for the "activities" table.
	ActivitiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "delete_time", Type: field.TypeTime, Nullable: true},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "start_time", Type: field.TypeTime},
		{Name: "end_time", Type: field.TypeTime},
		{Name: "remark", Type: field.TypeString},
		{Name: "online", Type: field.TypeInt},
		{Name: "entrance_img", Type: field.TypeString},
		{Name: "internal_top_img", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
	}
	// ActivitiesTable holds the schema information for the "activities" table.
	ActivitiesTable = &schema.Table{
		Name:       "activities",
		Columns:    ActivitiesColumns,
		PrimaryKey: []*schema.Column{ActivitiesColumns[0]},
	}
	// BannersColumns holds the columns for the "banners" table.
	BannersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "delete_time", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "title", Type: field.TypeString},
		{Name: "img", Type: field.TypeString},
	}
	// BannersTable holds the schema information for the "banners" table.
	BannersTable = &schema.Table{
		Name:       "banners",
		Columns:    BannersColumns,
		PrimaryKey: []*schema.Column{BannersColumns[0]},
	}
	// BannerItemsColumns holds the columns for the "banner_items" table.
	BannerItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "delete_time", Type: field.TypeTime, Nullable: true},
		{Name: "img", Type: field.TypeString},
		{Name: "keyword", Type: field.TypeString},
		{Name: "type", Type: field.TypeInt},
		{Name: "name", Type: field.TypeString},
		{Name: "banner_id", Type: field.TypeInt64, Nullable: true},
	}
	// BannerItemsTable holds the schema information for the "banner_items" table.
	BannerItemsTable = &schema.Table{
		Name:       "banner_items",
		Columns:    BannerItemsColumns,
		PrimaryKey: []*schema.Column{BannerItemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "banner_items_banners_banner_item",
				Columns:    []*schema.Column{BannerItemsColumns[8]},
				RefColumns: []*schema.Column{BannersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CategoriesColumns holds the columns for the "categories" table.
	CategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "delete_time", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "is_root", Type: field.TypeInt},
		{Name: "img", Type: field.TypeString},
		{Name: "index", Type: field.TypeInt},
		{Name: "online", Type: field.TypeInt},
		{Name: "level", Type: field.TypeInt},
		{Name: "parent_id", Type: field.TypeInt64, Nullable: true},
	}
	// CategoriesTable holds the schema information for the "categories" table.
	CategoriesTable = &schema.Table{
		Name:       "categories",
		Columns:    CategoriesColumns,
		PrimaryKey: []*schema.Column{CategoriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "categories_categories_children",
				Columns:    []*schema.Column{CategoriesColumns[11]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ChargesColumns holds the columns for the "charges" table.
	ChargesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "delete_time", Type: field.TypeTime, Nullable: true},
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "amount", Type: field.TypeString},
		{Name: "charge_no", Type: field.TypeString},
		{Name: "transaction_id", Type: field.TypeString},
		{Name: "pay_way", Type: field.TypeInt},
		{Name: "client_type", Type: field.TypeInt},
	}
	// ChargesTable holds the schema information for the "charges" table.
	ChargesTable = &schema.Table{
		Name:       "charges",
		Columns:    ChargesColumns,
		PrimaryKey: []*schema.Column{ChargesColumns[0]},
	}
	// GridCategoryColumns holds the columns for the "grid_category" table.
	GridCategoryColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "delete_time", Type: field.TypeTime, Nullable: true},
		{Name: "title", Type: field.TypeString},
		{Name: "img", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "category_id", Type: field.TypeInt},
		{Name: "root_category_id", Type: field.TypeInt},
	}
	// GridCategoryTable holds the schema information for the "grid_category" table.
	GridCategoryTable = &schema.Table{
		Name:       "grid_category",
		Columns:    GridCategoryColumns,
		PrimaryKey: []*schema.Column{GridCategoryColumns[0]},
	}
	// RefundsColumns holds the columns for the "refunds" table.
	RefundsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "delete_time", Type: field.TypeTime, Nullable: true},
		{Name: "refund_no", Type: field.TypeString},
		{Name: "transaction_id", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeInt64, Nullable: true},
		{Name: "reason", Type: field.TypeString},
		{Name: "order_id", Type: field.TypeInt64, Nullable: true},
		{Name: "order_sub_id", Type: field.TypeInt64, Nullable: true},
		{Name: "status", Type: field.TypeInt},
	}
	// RefundsTable holds the schema information for the "refunds" table.
	RefundsTable = &schema.Table{
		Name:       "refunds",
		Columns:    RefundsColumns,
		PrimaryKey: []*schema.Column{RefundsColumns[0]},
	}
	// ThemeColumns holds the columns for the "theme" table.
	ThemeColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "delete_time", Type: field.TypeTime, Nullable: true},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "tpl_name", Type: field.TypeString},
		{Name: "entrance_img", Type: field.TypeString},
		{Name: "extend", Type: field.TypeString},
		{Name: "internal_top_img", Type: field.TypeString},
		{Name: "title_img", Type: field.TypeString},
		{Name: "online", Type: field.TypeInt},
	}
	// ThemeTable holds the schema information for the "theme" table.
	ThemeTable = &schema.Table{
		Name:       "theme",
		Columns:    ThemeColumns,
		PrimaryKey: []*schema.Column{ThemeColumns[0]},
	}
	// ThemeSpuColumns holds the columns for the "theme_spu" table.
	ThemeSpuColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "spu_id", Type: field.TypeInt64},
		{Name: "theme_id", Type: field.TypeInt64, Nullable: true},
	}
	// ThemeSpuTable holds the schema information for the "theme_spu" table.
	ThemeSpuTable = &schema.Table{
		Name:       "theme_spu",
		Columns:    ThemeSpuColumns,
		PrimaryKey: []*schema.Column{ThemeSpuColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "theme_spu_theme_theme_spu",
				Columns:    []*schema.Column{ThemeSpuColumns[2]},
				RefColumns: []*schema.Column{ThemeColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ActivitiesTable,
		BannersTable,
		BannerItemsTable,
		CategoriesTable,
		ChargesTable,
		GridCategoryTable,
		RefundsTable,
		ThemeTable,
		ThemeSpuTable,
	}
)

func init() {
	BannerItemsTable.ForeignKeys[0].RefTable = BannersTable
	CategoriesTable.ForeignKeys[0].RefTable = CategoriesTable
	GridCategoryTable.Annotation = &entsql.Annotation{
		Table: "grid_category",
	}
	ThemeTable.Annotation = &entsql.Annotation{
		Table: "theme",
	}
	ThemeSpuTable.ForeignKeys[0].RefTable = ThemeTable
	ThemeSpuTable.Annotation = &entsql.Annotation{
		Table: "theme_spu",
	}
}
