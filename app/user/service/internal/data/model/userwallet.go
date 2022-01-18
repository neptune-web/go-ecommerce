// Code generated by entc, DO NOT EDIT.

package model

import (
	"fmt"
	"mall-go/app/user/service/internal/data/model/userwallet"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// UserWallet is the model entity for the UserWallet schema.
type UserWallet struct {
	config `json:"-"`
	// ID of the ent.
	// 用户ID
	ID int64 `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// DeleteTime holds the value of the "delete_time" field.
	DeleteTime time.Time `json:"delete_time,omitempty"`
	// Value holds the value of the "value" field.
	// 余额
	Value int `json:"value,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserWallet) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case userwallet.FieldID, userwallet.FieldValue:
			values[i] = new(sql.NullInt64)
		case userwallet.FieldCreateTime, userwallet.FieldUpdateTime, userwallet.FieldDeleteTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type UserWallet", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserWallet fields.
func (uw *UserWallet) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userwallet.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			uw.ID = int64(value.Int64)
		case userwallet.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				uw.CreateTime = value.Time
			}
		case userwallet.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				uw.UpdateTime = value.Time
			}
		case userwallet.FieldDeleteTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field delete_time", values[i])
			} else if value.Valid {
				uw.DeleteTime = value.Time
			}
		case userwallet.FieldValue:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				uw.Value = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this UserWallet.
// Note that you need to call UserWallet.Unwrap() before calling this method if this UserWallet
// was returned from a transaction, and the transaction was committed or rolled back.
func (uw *UserWallet) Update() *UserWalletUpdateOne {
	return (&UserWalletClient{config: uw.config}).UpdateOne(uw)
}

// Unwrap unwraps the UserWallet entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (uw *UserWallet) Unwrap() *UserWallet {
	tx, ok := uw.config.driver.(*txDriver)
	if !ok {
		panic("model: UserWallet is not a transactional entity")
	}
	uw.config.driver = tx.drv
	return uw
}

// String implements the fmt.Stringer.
func (uw *UserWallet) String() string {
	var builder strings.Builder
	builder.WriteString("UserWallet(")
	builder.WriteString(fmt.Sprintf("id=%v", uw.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(uw.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(uw.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", delete_time=")
	builder.WriteString(uw.DeleteTime.Format(time.ANSIC))
	builder.WriteString(", value=")
	builder.WriteString(fmt.Sprintf("%v", uw.Value))
	builder.WriteByte(')')
	return builder.String()
}

// UserWallets is a parsable slice of UserWallet.
type UserWallets []*UserWallet

func (uw UserWallets) config(cfg config) {
	for _i := range uw {
		uw[_i].config = cfg
	}
}