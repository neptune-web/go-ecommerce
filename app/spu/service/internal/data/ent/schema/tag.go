package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Tag struct {
	ent.Schema
}

func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Comment("中文限制6个，英文限制12个，由逻辑层控制"),
		field.String("description").Comment(""),

		field.Int("highlight").Comment(""),
		field.Int("type").Comment(""),
	}
}

func (Tag) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (Tag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("spu", Spu.Type).Ref("tag"),
	}
}
