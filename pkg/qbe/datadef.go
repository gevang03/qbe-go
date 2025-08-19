package qbe

import (
	"fmt"
	"strings"
)

// A Data struct represent a data definition in QBE IL.
type Data struct {
	Linkage              // The linkage of the data definition
	name    GlobalSymbol // The symbol that references the data
	Align   uint         // The required alignment of the data. Set to zero for default alignment
	fields  []dataField  // Values contained into the data definition.
	size    uint
}

func (*Data) isDefinition() {}

type (
	dataField interface{ isDataField() }

	dataFieldValue struct {
		Type  ExtendedType
		Items []DataItem
	}

	dataFieldZeroes uint
)

func (dataFieldValue) isDataField()  {}
func (dataFieldZeroes) isDataField() {}

// newData returns a new Data with symbol name, private linkage aligment 0, and no data entries.
func newData(name GlobalSymbol) *Data {
	return &Data{Linkage: PrivateLinkage(), name: name, Align: 0, fields: nil, size: 0}
}

// Name returns the name of data.
func (data *Data) Name() GlobalSymbol {
	return data.name
}

// InsertValue inserts to the end of the data entries the values of items with type type_.
func (data *Data) InsertValue(type_ ExtendedType, items ...DataItem) {
	if type_ == nil {
		panic("data entry type cannot be nil")
	}
	for _, item := range items {
		if item == nil {
			panic("data entry cannot be nil")
		}
	}
	if type_ != ByteType() {
		data.size += type_.SizeOf() * uint(len(items))
	} else {
		for _, item := range items {
			if str, ok := item.(DataString); ok {
				data.size += uint(len(str))
			} else {
				data.size += 1
			}
		}
	}
	data.fields = append(data.fields, dataFieldValue{type_, items})
}

// InsertZeroes inserts count zeroes to the end of data.
func (data *Data) InsertZeroes(count uint) {
	data.size += count
	data.fields = append(data.fields, dataFieldZeroes(count))
}

// SizeOf returns the number of bytes needed for data.
func (data *Data) SizeOf() uint {
	return data.size
}

// String converts data to a string compatible with QBE code.
func (data *Data) String() string {
	linkage := data.Linkage.String()
	var parts []string
	if linkage != "" {
		parts = append(parts, linkage)
	}
	parts = append(parts, "data", data.name.String(), "=")
	if data.Align != 0 {
		parts = append(parts, fmt.Sprint(data.Align))
	}
	parts = append(parts, "{")
	for _, field := range data.fields {
		switch f := field.(type) {
		case dataFieldValue:
			parts = append(parts, fmt.Sprint(f.Type))
			for i, item := range f.Items {
				if i < len(f.Items)-1 {
					parts = append(parts, fmt.Sprint(item))
				} else {
					parts = append(parts, fmt.Sprintf("%v,", item))
				}
			}
		case dataFieldZeroes:
			parts = append(parts, fmt.Sprintf("z %v,", f))
		default:
			panic("unreachable")
		}
	}
	parts = append(parts, "}")
	return strings.Join(parts, " ")
}
