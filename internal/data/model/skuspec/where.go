// Code generated by entc, DO NOT EDIT.

package skuspec

import (
	"mall-go/internal/data/model/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.SkuSpec {
	return predicate.SkuSpec(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.SkuSpec {
	return predicate.SkuSpec(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.SkuSpec {
	return predicate.SkuSpec(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.SkuSpec {
	return predicate.SkuSpec(func(s *sql.Selector) {
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
func IDNotIn(ids ...int64) predicate.SkuSpec {
	return predicate.SkuSpec(func(s *sql.Selector) {
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
func IDGT(id int64) predicate.SkuSpec {
	return predicate.SkuSpec(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.SkuSpec {
	return predicate.SkuSpec(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.SkuSpec {
	return predicate.SkuSpec(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.SkuSpec {
	return predicate.SkuSpec(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// SpuID applies equality check predicate on the "spu_id" field. It's identical to SpuIDEQ.
func SpuID(v int64) predicate.SkuSpec {
	return predicate.SkuSpec(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSpuID), v))
	})
}

// SkuID applies equality check predicate on the "sku_id" field. It's identical to SkuIDEQ.
func SkuID(v int64) predicate.SkuSpec {
	return predicate.SkuSpec(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSkuID), v))
	})
}

// KeyID applies equality check predicate on the "key_id" field. It's identical to KeyIDEQ.
func KeyID(v int64) predicate.SkuSpec {
	return predicate.SkuSpec(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldKeyID), v))
	})
}

// ValueID applies equality check predicate on the "value_id" field. It's identical to ValueIDEQ.
func ValueID(v int64) predicate.SkuSpec {
	return predicate.SkuSpec(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValueID), v))
	})
}

// SpuIDEQ applies the EQ predicate on the "spu_id" field.
func SpuIDEQ(v int64) predicate.SkuSpec {
	return predicate.SkuSpec(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSpuID), v))
	})
}

// SpuIDNEQ applies the NEQ predicate on the "spu_id" field.
func SpuIDNEQ(v int64) predicate.SkuSpec {
	return predicate.SkuSpec(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSpuID), v))
	})
}

// SpuIDIn applies the In predicate on the "spu_id" field.
func SpuIDIn(vs ...int64) predicate.SkuSpec {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SkuSpec(func(s *sql.Selector) {
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
func SpuIDNotIn(vs ...int64) predicate.SkuSpec {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SkuSpec(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSpuID), v...))
	})
}

// SpuIDGT applies the GT predicate on the "spu_id" field.
func SpuIDGT(v int64) predicate.SkuSpec {
	return predicate.SkuSpec(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSpuID), v))
	})
}

// SpuIDGTE applies the GTE predicate on the "spu_id" field.
func SpuIDGTE(v int64) predicate.SkuSpec {
	return predicate.SkuSpec(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSpuID), v))
	})
}

// SpuIDLT applies the LT predicate on the "spu_id" field.
func SpuIDLT(v int64) predicate.SkuSpec {
	return predicate.SkuSpec(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSpuID), v))
	})
}

// SpuIDLTE applies the LTE predicate on the "spu_id" field.
func SpuIDLTE(v int64) predicate.SkuSpec {
	return predicate.SkuSpec(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSpuID), v))
	})
}

// SkuIDEQ applies the EQ predicate on the "sku_id" field.
func SkuIDEQ(v int64) predicate.SkuSpec {
	return predicate.SkuSpec(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSkuID), v))
	})
}

// SkuIDNEQ applies the NEQ predicate on the "sku_id" field.
func SkuIDNEQ(v int64) predicate.SkuSpec {
	return predicate.SkuSpec(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSkuID), v))
	})
}

// SkuIDIn applies the In predicate on the "sku_id" field.
func SkuIDIn(vs ...int64) predicate.SkuSpec {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SkuSpec(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSkuID), v...))
	})
}

// SkuIDNotIn applies the NotIn predicate on the "sku_id" field.
func SkuIDNotIn(vs ...int64) predicate.SkuSpec {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SkuSpec(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSkuID), v...))
	})
}

// SkuIDGT applies the GT predicate on the "sku_id" field.
func SkuIDGT(v int64) predicate.SkuSpec {
	return predicate.SkuSpec(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSkuID), v))
	})
}

// SkuIDGTE applies the GTE predicate on the "sku_id" field.
func SkuIDGTE(v int64) predicate.SkuSpec {
	return predicate.SkuSpec(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSkuID), v))
	})
}

// SkuIDLT applies the LT predicate on the "sku_id" field.
func SkuIDLT(v int64) predicate.SkuSpec {
	return predicate.SkuSpec(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSkuID), v))
	})
}

// SkuIDLTE applies the LTE predicate on the "sku_id" field.
func SkuIDLTE(v int64) predicate.SkuSpec {
	return predicate.SkuSpec(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSkuID), v))
	})
}

// KeyIDEQ applies the EQ predicate on the "key_id" field.
func KeyIDEQ(v int64) predicate.SkuSpec {
	return predicate.SkuSpec(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldKeyID), v))
	})
}

// KeyIDNEQ applies the NEQ predicate on the "key_id" field.
func KeyIDNEQ(v int64) predicate.SkuSpec {
	return predicate.SkuSpec(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldKeyID), v))
	})
}

// KeyIDIn applies the In predicate on the "key_id" field.
func KeyIDIn(vs ...int64) predicate.SkuSpec {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SkuSpec(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldKeyID), v...))
	})
}

// KeyIDNotIn applies the NotIn predicate on the "key_id" field.
func KeyIDNotIn(vs ...int64) predicate.SkuSpec {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SkuSpec(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldKeyID), v...))
	})
}

// KeyIDGT applies the GT predicate on the "key_id" field.
func KeyIDGT(v int64) predicate.SkuSpec {
	return predicate.SkuSpec(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldKeyID), v))
	})
}

// KeyIDGTE applies the GTE predicate on the "key_id" field.
func KeyIDGTE(v int64) predicate.SkuSpec {
	return predicate.SkuSpec(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldKeyID), v))
	})
}

// KeyIDLT applies the LT predicate on the "key_id" field.
func KeyIDLT(v int64) predicate.SkuSpec {
	return predicate.SkuSpec(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldKeyID), v))
	})
}

// KeyIDLTE applies the LTE predicate on the "key_id" field.
func KeyIDLTE(v int64) predicate.SkuSpec {
	return predicate.SkuSpec(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldKeyID), v))
	})
}

// ValueIDEQ applies the EQ predicate on the "value_id" field.
func ValueIDEQ(v int64) predicate.SkuSpec {
	return predicate.SkuSpec(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValueID), v))
	})
}

// ValueIDNEQ applies the NEQ predicate on the "value_id" field.
func ValueIDNEQ(v int64) predicate.SkuSpec {
	return predicate.SkuSpec(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldValueID), v))
	})
}

// ValueIDIn applies the In predicate on the "value_id" field.
func ValueIDIn(vs ...int64) predicate.SkuSpec {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SkuSpec(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldValueID), v...))
	})
}

// ValueIDNotIn applies the NotIn predicate on the "value_id" field.
func ValueIDNotIn(vs ...int64) predicate.SkuSpec {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SkuSpec(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldValueID), v...))
	})
}

// ValueIDGT applies the GT predicate on the "value_id" field.
func ValueIDGT(v int64) predicate.SkuSpec {
	return predicate.SkuSpec(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldValueID), v))
	})
}

// ValueIDGTE applies the GTE predicate on the "value_id" field.
func ValueIDGTE(v int64) predicate.SkuSpec {
	return predicate.SkuSpec(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldValueID), v))
	})
}

// ValueIDLT applies the LT predicate on the "value_id" field.
func ValueIDLT(v int64) predicate.SkuSpec {
	return predicate.SkuSpec(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldValueID), v))
	})
}

// ValueIDLTE applies the LTE predicate on the "value_id" field.
func ValueIDLTE(v int64) predicate.SkuSpec {
	return predicate.SkuSpec(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldValueID), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SkuSpec) predicate.SkuSpec {
	return predicate.SkuSpec(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SkuSpec) predicate.SkuSpec {
	return predicate.SkuSpec(func(s *sql.Selector) {
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
func Not(p predicate.SkuSpec) predicate.SkuSpec {
	return predicate.SkuSpec(func(s *sql.Selector) {
		p(s.Not())
	})
}