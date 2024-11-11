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

type Variant []field

type Struct struct {
	Name   types.TypeName
	Align  uint
	fields []field
}

type Union struct {
	Name     types.TypeName
	Align    uint
	variants []Variant
}

type Opaque struct {
	Name  types.TypeName
	Align uint
	Size  uint
}

func (*Struct) isDefinition() {}
func (*Union) isDefinition()  {}
func (*Opaque) isDefinition() {}

func NewStruct(name types.TypeName) *Struct {
	return &Struct{name, 0, nil}
}

func NewUnion(name types.TypeName) *Union {
	return &Union{name, 0, nil}
}

func NewOpaque(name types.TypeName, align uint, size uint) *Opaque {
	return &Opaque{name, align, size}
}

func (s *Struct) InsertField(type_ types.SubType, count uint) {
	s.fields = append(s.fields, field{type_, count})
}

func (u *Union) InsertVariant(variant Variant) {
	u.variants = append(u.variants, variant)
}

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

func (s *Struct) String() string {
	parts := []string{"type", s.Name.String(), "="}
	if s.Align != 0 {
		parts = append(parts, "align", fmt.Sprint(s.Align))
	}
	parts = stringifyFields(parts, s.fields)
	return strings.Join(parts, " ")
}

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

func (o *Opaque) String() string {
	return fmt.Sprintf("type %v = align %v { %v }", o.Name, o.Align, o.Size)
}
