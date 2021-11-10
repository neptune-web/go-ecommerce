package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type UserCoupon struct {
	ent.Schema
}

func (UserCoupon) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user_coupon"},
	}
}
func (UserCoupon) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Comment("").Unique(),
		field.Int("coupon_id").Comment(""),
		field.Int8("status").Comment("1:未使用，2：已使用， 3：已过期"),

		field.Int("order_id").Comment(""),
	}
}
func (UserCoupon) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
