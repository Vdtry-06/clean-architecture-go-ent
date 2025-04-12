package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Book holds the schema definition for the Book entity.
type Book struct {
	ent.Schema
}

// Fields of the Book.
func (Book) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty(),
        field.String("author").NotEmpty(),
        field.Float("price").Min(0),
		field.Int("publish_year").Positive().Range(1000, time.Now().Year()),
	}
}
