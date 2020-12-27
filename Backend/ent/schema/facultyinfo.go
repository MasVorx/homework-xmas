package schema

import "github.com/facebookincubator/ent"

// Facultyinfo holds the schema definition for the Facultyinfo entity.
type Facultyinfo struct {
	ent.Schema
}

// Fields of the Facultyinfo.
func (Facultyinfo) Fields() []ent.Field {
	return nil
}

// Edges of the Facultyinfo.
func (Facultyinfo) Edges() []ent.Edge {
	return nil
}
