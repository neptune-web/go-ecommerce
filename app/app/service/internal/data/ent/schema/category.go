package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Category struct {
	ent.Schema
}

func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment(""),
		field.String("description").Comment(""),

		field.Int("is_root").Comment(""),
		field.Int64("parent_id").Optional().Comment(""),
		field.String("img").Comment(""),
		field.Int("index").Comment(""),
		field.Int("online").Comment(""),
		field.Int("level").Comment(""),
	}
}

func (Category) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (Category) Edges() []ent.Edge {
	return []ent.Edge{
		// edge.From("coupon", Coupon.Type).Ref("category"),
		edge.To("children", Category.Type).From("parent").
			Unique().Field("parent_id"),
	}
}
