// Code generated by entc, DO NOT EDIT.

package ordersnap

import (
	"mall-go/app/order/service/internal/data/model/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
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
func IDNotIn(ids ...int64) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
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
func IDGT(id int64) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// SnapImg applies equality check predicate on the "snap_img" field. It's identical to SnapImgEQ.
func SnapImg(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSnapImg), v))
	})
}

// SnapTitle applies equality check predicate on the "snap_title" field. It's identical to SnapTitleEQ.
func SnapTitle(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSnapTitle), v))
	})
}

// SnapItems applies equality check predicate on the "snap_items" field. It's identical to SnapItemsEQ.
func SnapItems(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSnapItems), v))
	})
}

// SnapAddress applies equality check predicate on the "snap_address" field. It's identical to SnapAddressEQ.
func SnapAddress(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSnapAddress), v))
	})
}

// OrderID applies equality check predicate on the "order_id" field. It's identical to OrderIDEQ.
func OrderID(v int64) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderID), v))
	})
}

// SnapImgEQ applies the EQ predicate on the "snap_img" field.
func SnapImgEQ(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSnapImg), v))
	})
}

// SnapImgNEQ applies the NEQ predicate on the "snap_img" field.
func SnapImgNEQ(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSnapImg), v))
	})
}

// SnapImgIn applies the In predicate on the "snap_img" field.
func SnapImgIn(vs ...string) predicate.OrderSnap {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderSnap(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSnapImg), v...))
	})
}

// SnapImgNotIn applies the NotIn predicate on the "snap_img" field.
func SnapImgNotIn(vs ...string) predicate.OrderSnap {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderSnap(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSnapImg), v...))
	})
}

// SnapImgGT applies the GT predicate on the "snap_img" field.
func SnapImgGT(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSnapImg), v))
	})
}

// SnapImgGTE applies the GTE predicate on the "snap_img" field.
func SnapImgGTE(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSnapImg), v))
	})
}

// SnapImgLT applies the LT predicate on the "snap_img" field.
func SnapImgLT(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSnapImg), v))
	})
}

// SnapImgLTE applies the LTE predicate on the "snap_img" field.
func SnapImgLTE(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSnapImg), v))
	})
}

// SnapImgContains applies the Contains predicate on the "snap_img" field.
func SnapImgContains(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSnapImg), v))
	})
}

// SnapImgHasPrefix applies the HasPrefix predicate on the "snap_img" field.
func SnapImgHasPrefix(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSnapImg), v))
	})
}

// SnapImgHasSuffix applies the HasSuffix predicate on the "snap_img" field.
func SnapImgHasSuffix(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSnapImg), v))
	})
}

// SnapImgEqualFold applies the EqualFold predicate on the "snap_img" field.
func SnapImgEqualFold(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSnapImg), v))
	})
}

// SnapImgContainsFold applies the ContainsFold predicate on the "snap_img" field.
func SnapImgContainsFold(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSnapImg), v))
	})
}

// SnapTitleEQ applies the EQ predicate on the "snap_title" field.
func SnapTitleEQ(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSnapTitle), v))
	})
}

// SnapTitleNEQ applies the NEQ predicate on the "snap_title" field.
func SnapTitleNEQ(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSnapTitle), v))
	})
}

// SnapTitleIn applies the In predicate on the "snap_title" field.
func SnapTitleIn(vs ...string) predicate.OrderSnap {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderSnap(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSnapTitle), v...))
	})
}

// SnapTitleNotIn applies the NotIn predicate on the "snap_title" field.
func SnapTitleNotIn(vs ...string) predicate.OrderSnap {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderSnap(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSnapTitle), v...))
	})
}

// SnapTitleGT applies the GT predicate on the "snap_title" field.
func SnapTitleGT(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSnapTitle), v))
	})
}

// SnapTitleGTE applies the GTE predicate on the "snap_title" field.
func SnapTitleGTE(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSnapTitle), v))
	})
}

// SnapTitleLT applies the LT predicate on the "snap_title" field.
func SnapTitleLT(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSnapTitle), v))
	})
}

// SnapTitleLTE applies the LTE predicate on the "snap_title" field.
func SnapTitleLTE(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSnapTitle), v))
	})
}

// SnapTitleContains applies the Contains predicate on the "snap_title" field.
func SnapTitleContains(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSnapTitle), v))
	})
}

// SnapTitleHasPrefix applies the HasPrefix predicate on the "snap_title" field.
func SnapTitleHasPrefix(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSnapTitle), v))
	})
}

// SnapTitleHasSuffix applies the HasSuffix predicate on the "snap_title" field.
func SnapTitleHasSuffix(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSnapTitle), v))
	})
}

// SnapTitleEqualFold applies the EqualFold predicate on the "snap_title" field.
func SnapTitleEqualFold(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSnapTitle), v))
	})
}

// SnapTitleContainsFold applies the ContainsFold predicate on the "snap_title" field.
func SnapTitleContainsFold(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSnapTitle), v))
	})
}

// SnapItemsEQ applies the EQ predicate on the "snap_items" field.
func SnapItemsEQ(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSnapItems), v))
	})
}

// SnapItemsNEQ applies the NEQ predicate on the "snap_items" field.
func SnapItemsNEQ(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSnapItems), v))
	})
}

// SnapItemsIn applies the In predicate on the "snap_items" field.
func SnapItemsIn(vs ...string) predicate.OrderSnap {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderSnap(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSnapItems), v...))
	})
}

// SnapItemsNotIn applies the NotIn predicate on the "snap_items" field.
func SnapItemsNotIn(vs ...string) predicate.OrderSnap {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderSnap(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSnapItems), v...))
	})
}

// SnapItemsGT applies the GT predicate on the "snap_items" field.
func SnapItemsGT(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSnapItems), v))
	})
}

// SnapItemsGTE applies the GTE predicate on the "snap_items" field.
func SnapItemsGTE(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSnapItems), v))
	})
}

// SnapItemsLT applies the LT predicate on the "snap_items" field.
func SnapItemsLT(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSnapItems), v))
	})
}

// SnapItemsLTE applies the LTE predicate on the "snap_items" field.
func SnapItemsLTE(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSnapItems), v))
	})
}

// SnapItemsContains applies the Contains predicate on the "snap_items" field.
func SnapItemsContains(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSnapItems), v))
	})
}

// SnapItemsHasPrefix applies the HasPrefix predicate on the "snap_items" field.
func SnapItemsHasPrefix(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSnapItems), v))
	})
}

// SnapItemsHasSuffix applies the HasSuffix predicate on the "snap_items" field.
func SnapItemsHasSuffix(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSnapItems), v))
	})
}

// SnapItemsEqualFold applies the EqualFold predicate on the "snap_items" field.
func SnapItemsEqualFold(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSnapItems), v))
	})
}

// SnapItemsContainsFold applies the ContainsFold predicate on the "snap_items" field.
func SnapItemsContainsFold(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSnapItems), v))
	})
}

// SnapAddressEQ applies the EQ predicate on the "snap_address" field.
func SnapAddressEQ(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSnapAddress), v))
	})
}

// SnapAddressNEQ applies the NEQ predicate on the "snap_address" field.
func SnapAddressNEQ(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSnapAddress), v))
	})
}

// SnapAddressIn applies the In predicate on the "snap_address" field.
func SnapAddressIn(vs ...string) predicate.OrderSnap {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderSnap(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSnapAddress), v...))
	})
}

// SnapAddressNotIn applies the NotIn predicate on the "snap_address" field.
func SnapAddressNotIn(vs ...string) predicate.OrderSnap {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderSnap(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSnapAddress), v...))
	})
}

// SnapAddressGT applies the GT predicate on the "snap_address" field.
func SnapAddressGT(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSnapAddress), v))
	})
}

// SnapAddressGTE applies the GTE predicate on the "snap_address" field.
func SnapAddressGTE(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSnapAddress), v))
	})
}

// SnapAddressLT applies the LT predicate on the "snap_address" field.
func SnapAddressLT(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSnapAddress), v))
	})
}

// SnapAddressLTE applies the LTE predicate on the "snap_address" field.
func SnapAddressLTE(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSnapAddress), v))
	})
}

// SnapAddressContains applies the Contains predicate on the "snap_address" field.
func SnapAddressContains(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSnapAddress), v))
	})
}

// SnapAddressHasPrefix applies the HasPrefix predicate on the "snap_address" field.
func SnapAddressHasPrefix(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSnapAddress), v))
	})
}

// SnapAddressHasSuffix applies the HasSuffix predicate on the "snap_address" field.
func SnapAddressHasSuffix(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSnapAddress), v))
	})
}

// SnapAddressEqualFold applies the EqualFold predicate on the "snap_address" field.
func SnapAddressEqualFold(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSnapAddress), v))
	})
}

// SnapAddressContainsFold applies the ContainsFold predicate on the "snap_address" field.
func SnapAddressContainsFold(v string) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSnapAddress), v))
	})
}

// OrderIDEQ applies the EQ predicate on the "order_id" field.
func OrderIDEQ(v int64) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderID), v))
	})
}

// OrderIDNEQ applies the NEQ predicate on the "order_id" field.
func OrderIDNEQ(v int64) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOrderID), v))
	})
}

// OrderIDIn applies the In predicate on the "order_id" field.
func OrderIDIn(vs ...int64) predicate.OrderSnap {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderSnap(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOrderID), v...))
	})
}

// OrderIDNotIn applies the NotIn predicate on the "order_id" field.
func OrderIDNotIn(vs ...int64) predicate.OrderSnap {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderSnap(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOrderID), v...))
	})
}

// OrderIDIsNil applies the IsNil predicate on the "order_id" field.
func OrderIDIsNil() predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOrderID)))
	})
}

// OrderIDNotNil applies the NotNil predicate on the "order_id" field.
func OrderIDNotNil() predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOrderID)))
	})
}

// HasOrder applies the HasEdge predicate on the "order" edge.
func HasOrder() predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OrderTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OrderTable, OrderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrderWith applies the HasEdge predicate on the "order" edge with a given conditions (other predicates).
func HasOrderWith(preds ...predicate.Order) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OrderInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OrderTable, OrderColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OrderSnap) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OrderSnap) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
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
func Not(p predicate.OrderSnap) predicate.OrderSnap {
	return predicate.OrderSnap(func(s *sql.Selector) {
		p(s.Not())
	})
}