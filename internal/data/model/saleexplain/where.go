// Code generated by entc, DO NOT EDIT.

package saleexplain

import (
	"mall-go/internal/data/model/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// DeleteTime applies equality check predicate on the "delete_time" field. It's identical to DeleteTimeEQ.
func DeleteTime(v time.Time) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteTime), v))
	})
}

// Fixed applies equality check predicate on the "fixed" field. It's identical to FixedEQ.
func Fixed(v int) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFixed), v))
	})
}

// Text applies equality check predicate on the "text" field. It's identical to TextEQ.
func Text(v string) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldText), v))
	})
}

// SpuID applies equality check predicate on the "spu_id" field. It's identical to SpuIDEQ.
func SpuID(v int64) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSpuID), v))
	})
}

// Index applies equality check predicate on the "index" field. It's identical to IndexEQ.
func Index(v int) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIndex), v))
	})
}

// ReplaceID applies equality check predicate on the "replace_id" field. It's identical to ReplaceIDEQ.
func ReplaceID(v int64) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReplaceID), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.SaleExplain {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SaleExplain(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.SaleExplain {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SaleExplain(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.SaleExplain {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SaleExplain(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.SaleExplain {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SaleExplain(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// DeleteTimeEQ applies the EQ predicate on the "delete_time" field.
func DeleteTimeEQ(v time.Time) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeNEQ applies the NEQ predicate on the "delete_time" field.
func DeleteTimeNEQ(v time.Time) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeIn applies the In predicate on the "delete_time" field.
func DeleteTimeIn(vs ...time.Time) predicate.SaleExplain {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SaleExplain(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeleteTime), v...))
	})
}

// DeleteTimeNotIn applies the NotIn predicate on the "delete_time" field.
func DeleteTimeNotIn(vs ...time.Time) predicate.SaleExplain {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SaleExplain(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeleteTime), v...))
	})
}

// DeleteTimeGT applies the GT predicate on the "delete_time" field.
func DeleteTimeGT(v time.Time) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeGTE applies the GTE predicate on the "delete_time" field.
func DeleteTimeGTE(v time.Time) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeLT applies the LT predicate on the "delete_time" field.
func DeleteTimeLT(v time.Time) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeLTE applies the LTE predicate on the "delete_time" field.
func DeleteTimeLTE(v time.Time) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeIsNil applies the IsNil predicate on the "delete_time" field.
func DeleteTimeIsNil() predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeleteTime)))
	})
}

// DeleteTimeNotNil applies the NotNil predicate on the "delete_time" field.
func DeleteTimeNotNil() predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeleteTime)))
	})
}

// FixedEQ applies the EQ predicate on the "fixed" field.
func FixedEQ(v int) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFixed), v))
	})
}

// FixedNEQ applies the NEQ predicate on the "fixed" field.
func FixedNEQ(v int) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFixed), v))
	})
}

// FixedIn applies the In predicate on the "fixed" field.
func FixedIn(vs ...int) predicate.SaleExplain {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SaleExplain(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFixed), v...))
	})
}

// FixedNotIn applies the NotIn predicate on the "fixed" field.
func FixedNotIn(vs ...int) predicate.SaleExplain {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SaleExplain(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFixed), v...))
	})
}

// FixedGT applies the GT predicate on the "fixed" field.
func FixedGT(v int) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFixed), v))
	})
}

// FixedGTE applies the GTE predicate on the "fixed" field.
func FixedGTE(v int) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFixed), v))
	})
}

// FixedLT applies the LT predicate on the "fixed" field.
func FixedLT(v int) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFixed), v))
	})
}

// FixedLTE applies the LTE predicate on the "fixed" field.
func FixedLTE(v int) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFixed), v))
	})
}

// TextEQ applies the EQ predicate on the "text" field.
func TextEQ(v string) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldText), v))
	})
}

// TextNEQ applies the NEQ predicate on the "text" field.
func TextNEQ(v string) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldText), v))
	})
}

// TextIn applies the In predicate on the "text" field.
func TextIn(vs ...string) predicate.SaleExplain {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SaleExplain(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldText), v...))
	})
}

// TextNotIn applies the NotIn predicate on the "text" field.
func TextNotIn(vs ...string) predicate.SaleExplain {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SaleExplain(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldText), v...))
	})
}

// TextGT applies the GT predicate on the "text" field.
func TextGT(v string) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldText), v))
	})
}

// TextGTE applies the GTE predicate on the "text" field.
func TextGTE(v string) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldText), v))
	})
}

// TextLT applies the LT predicate on the "text" field.
func TextLT(v string) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldText), v))
	})
}

// TextLTE applies the LTE predicate on the "text" field.
func TextLTE(v string) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldText), v))
	})
}

// TextContains applies the Contains predicate on the "text" field.
func TextContains(v string) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldText), v))
	})
}

// TextHasPrefix applies the HasPrefix predicate on the "text" field.
func TextHasPrefix(v string) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldText), v))
	})
}

// TextHasSuffix applies the HasSuffix predicate on the "text" field.
func TextHasSuffix(v string) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldText), v))
	})
}

// TextEqualFold applies the EqualFold predicate on the "text" field.
func TextEqualFold(v string) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldText), v))
	})
}

// TextContainsFold applies the ContainsFold predicate on the "text" field.
func TextContainsFold(v string) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldText), v))
	})
}

// SpuIDEQ applies the EQ predicate on the "spu_id" field.
func SpuIDEQ(v int64) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSpuID), v))
	})
}

// SpuIDNEQ applies the NEQ predicate on the "spu_id" field.
func SpuIDNEQ(v int64) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSpuID), v))
	})
}

// SpuIDIn applies the In predicate on the "spu_id" field.
func SpuIDIn(vs ...int64) predicate.SaleExplain {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SaleExplain(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSpuID), v...))
	})
}

// SpuIDNotIn applies the NotIn predicate on the "spu_id" field.
func SpuIDNotIn(vs ...int64) predicate.SaleExplain {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SaleExplain(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSpuID), v...))
	})
}

// SpuIDIsNil applies the IsNil predicate on the "spu_id" field.
func SpuIDIsNil() predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSpuID)))
	})
}

// SpuIDNotNil applies the NotNil predicate on the "spu_id" field.
func SpuIDNotNil() predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSpuID)))
	})
}

// IndexEQ applies the EQ predicate on the "index" field.
func IndexEQ(v int) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIndex), v))
	})
}

// IndexNEQ applies the NEQ predicate on the "index" field.
func IndexNEQ(v int) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIndex), v))
	})
}

// IndexIn applies the In predicate on the "index" field.
func IndexIn(vs ...int) predicate.SaleExplain {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SaleExplain(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldIndex), v...))
	})
}

// IndexNotIn applies the NotIn predicate on the "index" field.
func IndexNotIn(vs ...int) predicate.SaleExplain {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SaleExplain(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldIndex), v...))
	})
}

// IndexGT applies the GT predicate on the "index" field.
func IndexGT(v int) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIndex), v))
	})
}

// IndexGTE applies the GTE predicate on the "index" field.
func IndexGTE(v int) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIndex), v))
	})
}

// IndexLT applies the LT predicate on the "index" field.
func IndexLT(v int) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIndex), v))
	})
}

// IndexLTE applies the LTE predicate on the "index" field.
func IndexLTE(v int) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIndex), v))
	})
}

// ReplaceIDEQ applies the EQ predicate on the "replace_id" field.
func ReplaceIDEQ(v int64) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReplaceID), v))
	})
}

// ReplaceIDNEQ applies the NEQ predicate on the "replace_id" field.
func ReplaceIDNEQ(v int64) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldReplaceID), v))
	})
}

// ReplaceIDIn applies the In predicate on the "replace_id" field.
func ReplaceIDIn(vs ...int64) predicate.SaleExplain {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SaleExplain(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldReplaceID), v...))
	})
}

// ReplaceIDNotIn applies the NotIn predicate on the "replace_id" field.
func ReplaceIDNotIn(vs ...int64) predicate.SaleExplain {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SaleExplain(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldReplaceID), v...))
	})
}

// ReplaceIDGT applies the GT predicate on the "replace_id" field.
func ReplaceIDGT(v int64) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldReplaceID), v))
	})
}

// ReplaceIDGTE applies the GTE predicate on the "replace_id" field.
func ReplaceIDGTE(v int64) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldReplaceID), v))
	})
}

// ReplaceIDLT applies the LT predicate on the "replace_id" field.
func ReplaceIDLT(v int64) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldReplaceID), v))
	})
}

// ReplaceIDLTE applies the LTE predicate on the "replace_id" field.
func ReplaceIDLTE(v int64) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldReplaceID), v))
	})
}

// HasSpu applies the HasEdge predicate on the "spu" edge.
func HasSpu() predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SpuTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SpuTable, SpuColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSpuWith applies the HasEdge predicate on the "spu" edge with a given conditions (other predicates).
func HasSpuWith(preds ...predicate.Spu) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SpuInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SpuTable, SpuColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SaleExplain) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SaleExplain) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SaleExplain) predicate.SaleExplain {
	return predicate.SaleExplain(func(s *sql.Selector) {
		p(s.Not())
	})
}