package qbe

import (
	"fmt"
	"maps"
	"strings"
)

var _ TypeDef = &Struct{}
var _ TypeDef = &Union{}
var _ TypeDef = &Opaque{}

// A TypeName represents the name given to an aggregate definition [Struct], [Union] or [Opaque].
type TypeName string

// String converts tn to a string compatible with QBE code.
func (tn TypeName) String() string {
	s := string(tn)
	if !validateName(s) {
		panic(fmt.Sprintf("Type name :%q is not valid", s))
	}
	return ":" + s
}

type field struct {
	Type  SubType
	Count uint
}

// A FieldList represents the inner structure of [Struct] and [Union] types.
type FieldList struct {
	fields []field
	size   uint
	align  uint
}

// A Struct represents the definition of a struct type.
type Struct struct {
	name      TypeName // Name of the struct
	Align     uint     // Alignment of the struct
	fieldList FieldList
}

// A Union represents the definition of a union type.
type Union struct {
	name     TypeName // Name of the union
	Align    uint     // Alignment of the union
	variants []FieldList
	size     uint
}

// An Opaque represents the definition of an opaque type.
type Opaque struct {
	name  TypeName // Name of the opaque type
	Align uint     // Alignment of the opaque type
	size  uint     // Size of the opaque type
}

func (*Struct) isSubType() {}
func (*Union) isSubType()  {}
func (*Opaque) isSubType() {}

func (*Struct) isABIType() {}
func (*Union) isABIType()  {}
func (*Opaque) isABIType() {}

func (*Struct) isRetType() {}
func (*Union) isRetType()  {}
func (*Opaque) isRetType() {}

func (s *Struct) Name() string {
	return s.name.String()
}

func (u *Union) Name() string {
	return u.name.String()
}

func (o *Opaque) Name() string {
	return o.name.String()
}

func (s *Struct) getFields() map[TypeDef]struct{} {
	return s.fieldList.fieldSet()
}

func (u *Union) getFields() map[TypeDef]struct{} {
	fieldSet := make(map[TypeDef]struct{})
	for _, variant := range u.variants {
		maps.Copy(fieldSet, variant.fieldSet())
	}
	return fieldSet
}

func (o *Opaque) getFields() map[TypeDef]struct{} {
	return nil
}

func newStruct(name TypeName) *Struct {
	return &Struct{name, 0, FieldList{}}
}

// newUnion returns a new [Union] with typename name, alignment zero and no variants.
func newUnion(name TypeName) *Union {
	return &Union{name, 0, nil, 0}
}

// newOpaque returns a new [Opaque] with typename name,
// alignment align and size equal to size.
func newOpaque(name TypeName, align uint, size uint) *Opaque {
	return &Opaque{name, align, size}
}

func isPowerOf2(n uint) bool {
	return n > 0 && n&(n-1) == 0
}

func alignedSize(size, align uint) uint {
	if align == 0 {
		return size
	}
	if !isPowerOf2(align) {
		panic(fmt.Sprintf("align %#v is not a power of two", align))
	}
	return (size + align - 1) & ^(align - 1)
}

func (s *Struct) SizeOf() uint {
	return alignedSize(s.fieldList.size, s.AlignOf())
}

func (u *Union) SizeOf() uint {
	return alignedSize(u.size, u.AlignOf())
}

func (o *Opaque) SizeOf() uint {
	return alignedSize(o.size, o.AlignOf())
}

func (s *Struct) AlignOf() uint {
	if s.Align != 0 {
		return s.Align
	}
	return s.fieldList.align
}

func (u *Union) AlignOf() uint {
	if u.Align != 0 {
		return u.Align
	}
	alignment := uint(0)
	for _, variant := range u.variants {
		alignment = max(alignment, variant.align)
	}
	return alignment
}

func (o *Opaque) AlignOf() uint {
	return o.Align
}

// InsertField appends field of type type_ and count elements to fl.
// Never insert types that have not been completed to avoid inconsistencies. Returns
// then field's index in fl.
func (fl *FieldList) InsertField(type_ SubType, count uint) int {
	if type_ == nil {
		panic("field type cannot be nil")
	}
	index := len(fl.fields)
	fl.fields = append(fl.fields, field{type_, count})
	fl.size += type_.SizeOf() * count
	fl.align = max(fl.align, type_.AlignOf())
	return index
}

// InsertFieldAligned appends field of type type_ and count elements to fl after
// aligning to type_'s alignment. Returns the field's index in fl.
func (fl *FieldList) InsertFieldAligned(type_ SubType, count uint) int {
	if type_ == nil {
		panic("field type cannot be nil")
	}
	alignment := type_.AlignOf()
	padded := alignedSize(fl.size, alignment)
	if padded > fl.size {
		fl.fields = append(fl.fields, field{ByteType(), padded - fl.size})
		fl.size = padded
	}
	return fl.InsertField(type_, count)
}

func (fl *FieldList) fieldSet() map[TypeDef]struct{} {
	fieldSet := make(map[TypeDef]struct{})
	for _, field := range fl.fields {
		t := field.Type
		if t, ok := t.(TypeDef); ok {
			fieldSet[t] = struct{}{}
		}
	}
	return fieldSet
}

func (fl *FieldList) offsetOf(index int) uint {
	bytes := uint(0)
	for i, field := range fl.fields {
		if i == index {
			return bytes
		}
		bytes += field.Count * field.Type.SizeOf()
	}
	panic(fmt.Sprintf("Field offset %v out of bounds", index))
}

// NewFieldList returns a new [FieldList].
func NewFieldList() FieldList {
	return FieldList{}
}

// OffsetOf returns the byte offset of the index field of s.
func (s *Struct) OffsetOf(index int) uint {
	return s.fieldList.offsetOf(index)
}

// OffsetOf returns the byte offset of the index field of variantIndex variant of u.
func (u *Union) OffsetOf(variantIndex int, index int) uint {
	return u.variants[variantIndex].offsetOf(index)
}

// SetFields sets fl as the fields of s.
func (s *Struct) SetFields(fl FieldList) {
	s.fieldList = fl
}

// InsertVariant adds fl as a variant of u.
func (u *Union) InsertVariant(fl FieldList) {
	u.variants = append(u.variants, fl)
	u.size = max(fl.size, u.size)
}

func stringifyFields(parts []string, fields []field) []string {
	parts = append(parts, "{")
	for _, field := range fields {
		if field.Count == 1 {
			parts = append(parts, fmt.Sprintf("%v,", field.Type.Name()))
		} else {
			parts = append(parts, fmt.Sprintf("%v %v,", field.Type.Name(), field.Count))
		}
	}
	parts = append(parts, "}")
	return parts
}

// String converts s to a string compatible with QBE code.
func (s *Struct) String() string {
	parts := []string{"type", s.name.String(), "="}
	if s.Align != 0 {
		parts = append(parts, "align", fmt.Sprint(s.Align))
	}
	parts = stringifyFields(parts, s.fieldList.fields)
	return strings.Join(parts, " ")
}

// String converts u to a string compatible with QBE code.
func (u *Union) String() string {
	parts := []string{"type", u.name.String(), "="}
	if u.Align != 0 {
		parts = append(parts, "align", fmt.Sprint(u.Align))
	}
	parts = append(parts, "{")
	for _, variant := range u.variants {
		parts = stringifyFields(parts, variant.fields)
	}
	parts = append(parts, "}")
	return strings.Join(parts, " ")
}

// String converts o to a string compatible with QBE code.
func (o *Opaque) String() string {
	return fmt.Sprintf("type %v = align %v { %v }", o.name, o.Align, o.size)
}
