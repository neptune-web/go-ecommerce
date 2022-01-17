// Code generated by entc, DO NOT EDIT.

package model

import (
	"fmt"
	"mall-go/app/app/service/internal/data/model/gridcategory"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// GridCategory is the model entity for the GridCategory schema.
type GridCategory struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// DeleteTime holds the value of the "delete_time" field.
	DeleteTime time.Time `json:"delete_time,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Img holds the value of the "img" field.
	Img string `json:"img,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// CategoryID holds the value of the "category_id" field.
	CategoryID int `json:"category_id,omitempty"`
	// RootCategoryID holds the value of the "root_category_id" field.
	RootCategoryID int `json:"root_category_id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GridCategory) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case gridcategory.FieldID, gridcategory.FieldCategoryID, gridcategory.FieldRootCategoryID:
			values[i] = new(sql.NullInt64)
		case gridcategory.FieldTitle, gridcategory.FieldImg, gridcategory.FieldName:
			values[i] = new(sql.NullString)
		case gridcategory.FieldCreateTime, gridcategory.FieldUpdateTime, gridcategory.FieldDeleteTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type GridCategory", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GridCategory fields.
func (gc *GridCategory) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case gridcategory.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			gc.ID = int64(value.Int64)
		case gridcategory.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				gc.CreateTime = value.Time
			}
		case gridcategory.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				gc.UpdateTime = value.Time
			}
		case gridcategory.FieldDeleteTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field delete_time", values[i])
			} else if value.Valid {
				gc.DeleteTime = value.Time
			}
		case gridcategory.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				gc.Title = value.String
			}
		case gridcategory.FieldImg:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field img", values[i])
			} else if value.Valid {
				gc.Img = value.String
			}
		case gridcategory.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				gc.Name = value.String
			}
		case gridcategory.FieldCategoryID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field category_id", values[i])
			} else if value.Valid {
				gc.CategoryID = int(value.Int64)
			}
		case gridcategory.FieldRootCategoryID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field root_category_id", values[i])
			} else if value.Valid {
				gc.RootCategoryID = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this GridCategory.
// Note that you need to call GridCategory.Unwrap() before calling this method if this GridCategory
// was returned from a transaction, and the transaction was committed or rolled back.
func (gc *GridCategory) Update() *GridCategoryUpdateOne {
	return (&GridCategoryClient{config: gc.config}).UpdateOne(gc)
}

// Unwrap unwraps the GridCategory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gc *GridCategory) Unwrap() *GridCategory {
	tx, ok := gc.config.driver.(*txDriver)
	if !ok {
		panic("model: GridCategory is not a transactional entity")
	}
	gc.config.driver = tx.drv
	return gc
}

// String implements the fmt.Stringer.
func (gc *GridCategory) String() string {
	var builder strings.Builder
	builder.WriteString("GridCategory(")
	builder.WriteString(fmt.Sprintf("id=%v", gc.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(gc.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(gc.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", delete_time=")
	builder.WriteString(gc.DeleteTime.Format(time.ANSIC))
	builder.WriteString(", title=")
	builder.WriteString(gc.Title)
	builder.WriteString(", img=")
	builder.WriteString(gc.Img)
	builder.WriteString(", name=")
	builder.WriteString(gc.Name)
	builder.WriteString(", category_id=")
	builder.WriteString(fmt.Sprintf("%v", gc.CategoryID))
	builder.WriteString(", root_category_id=")
	builder.WriteString(fmt.Sprintf("%v", gc.RootCategoryID))
	builder.WriteByte(')')
	return builder.String()
}

// GridCategories is a parsable slice of GridCategory.
type GridCategories []*GridCategory

func (gc GridCategories) config(cfg config) {
	for _i := range gc {
		gc[_i].config = cfg
	}
}
