package schema

import (
	"entgo.io/ent"

	"entgo.io/ent/schema/field"
)

type Coupon struct {
	ent.Schema
}

func (Coupon) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Comment(""),
		field.Time("start_time").Comment(""),
		field.Time("end_time").Comment(""),
		field.String("description").Comment(""),
		field.Float("full_money").Comment(""),
		field.Float("minus").Comment(""),
		field.Float("rate").Comment("国内多是打折，国外多是百分比 off"),
		field.Int("type").Comment("1. 满减券 2.折扣券 3.无门槛券 4.满金额折扣券"),

		field.Int("valitiy").Comment(""),
		field.Int64("activity_id").Optional().Comment(""),
		field.String("remark").Comment(""),
		field.Int("whole_store").Comment(""),
	}
}

func (Coupon) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (Coupon) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.To("category", Category.Type).StorageKey(
		//	edge.Table("coupon_category"), edge.Columns("coupon_id", "category_id")),
		//
		//edge.From("activity", Activity.Type).
		//	Ref("coupon"),
	}
}
