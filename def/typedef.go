package def

import (
	"fmt"
	"strings"

	"github.com/gevang03/qbe-go/types"
)

type field struct {
	Type  types.SubType
	Count uint
}

// A Variant represents a variant of a [Union] type.
type Variant []field

// A Struct represents the definition of a struct type.
type Struct struct {
	Name   types.TypeName // Name of the struct
	Align  uint           // Alignment of the struct
	fields []field
}

// A Union represents the definition of a union type.
type Union struct {
	Name     types.TypeName // Name of the union
	Align    uint           // Alignment of the union
	variants []Variant
}

// An Opaque represents the definition of an opaque type.
type Opaque struct {
	Name  types.TypeName // Name of the opaque type
	Align uint           // Alignment of the opaque type
	Size  uint           // Size of the opaque type
}

func (*Struct) isDefinition() {}
func (*Union) isDefinition()  {}
func (*Opaque) isDefinition() {}

// NewStruct returns a new [Struct] with typename name, alignment zero and no fields.
func NewStruct(name types.TypeName) *Struct {
	return &Struct{name, 0, nil}
}

// NewUnion returns a new [Union] with typename name, alignment zero and no variants.
func NewUnion(name types.TypeName) *Union {
	return &Union{name, 0, nil}
}

// NewUnion returns a new [Opaque] with typename name,
// alignment align and size equal to size.
func NewOpaque(name types.TypeName, align uint, size uint) *Opaque {
	return &Opaque{name, align, size}
}

// InsertField appends count fields of type type_ to s.
func (s *Struct) InsertField(type_ types.SubType, count uint) {
	s.fields = append(s.fields, field{type_, count})
}

// InsertVariant adds a variant to u.
func (u *Union) InsertVariant(variant Variant) {
	u.variants = append(u.variants, variant)
}

// InsertField appends count fields of type type_ to v.
func (v *Variant) InsertField(type_ types.SubType, count uint) {
	*v = append(*v, field{type_, count})
}

func stringifyFields(parts []string, fields []field) []string {
	parts = append(parts, "{")
	for _, field := range fields {
		if field.Count == 1 {
			parts = append(parts, fmt.Sprintf("%v,", field.Type))
		} else {
			parts = append(parts, fmt.Sprintf("%v %v,", field.Type, field.Count))
		}
	}
	parts = append(parts, "}")
	return parts
}

// String converts s to a string compatible with QBE code.
func (s *Struct) String() string {
	parts := []string{"type", s.Name.String(), "="}
	if s.Align != 0 {
		parts = append(parts, "align", fmt.Sprint(s.Align))
	}
	parts = stringifyFields(parts, s.fields)
	return strings.Join(parts, " ")
}

// String converts u to a string compatible with QBE code.
func (u *Union) String() string {
	parts := []string{"type", u.Name.String(), "="}
	if u.Align != 0 {
		parts = append(parts, "align", fmt.Sprint(u.Align))
	}
	parts = append(parts, "{")
	for _, variant := range u.variants {
		parts = stringifyFields(parts, variant)
	}
	parts = append(parts, "}")
	return strings.Join(parts, " ")
}

// String converts o to a string compatible with QBE code.
func (o *Opaque) String() string {
	return fmt.Sprintf("type %v = align %v { %v }", o.Name, o.Align, o.Size)
}
