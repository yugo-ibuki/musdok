package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Todo struct {
	ent.Schema
}

func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.String("title").Default(""),
		field.String("description").Default(""),
		field.Bool("is_done").Default(false),
	}
}
