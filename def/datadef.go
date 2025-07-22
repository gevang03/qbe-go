package def

import (
	"fmt"
	"strings"

	"github.com/gevang03/qbe-go/types"
	"github.com/gevang03/qbe-go/value"
)

// A Data struct represent a data definition in QBE IL.
type Data struct {
	Linkage                    // The linkage of the data definition
	Name    value.GlobalSymbol // The symbol that references the data
	Align   uint               // The required alignment of the data. Set to zero for default alignment
	fields  []dataField        // Values contained into the data definition.
}

func (*Data) isDefinition() {}

type (
	dataField interface{ isDataField() }

	dataFieldValue struct {
		Type  types.ExtendedType
		Items []value.DataItem
	}

	dataFieldZeroes uint
)

func (dataFieldValue) isDataField()  {}
func (dataFieldZeroes) isDataField() {}

// NewData returns a new Data with symbol name, private linkage aligment 0, and no data entries.
func NewData(name value.GlobalSymbol) *Data {
	return &Data{Linkage: PrivateLinkage(), Name: name, Align: 0, fields: nil}
}

// InsertValue inserts to the end of the data entries the values of item... with type type_.
func (data *Data) InsertValue(type_ types.ExtendedType, item ...value.DataItem) {
	data.fields = append(data.fields, dataFieldValue{type_, item})
}

// InsertValue inserts count zeroes to the end of data.
func (data *Data) InsertZeroes(count uint) {
	data.fields = append(data.fields, dataFieldZeroes(count))
}

// String converts data to a string compatible with QBE code.
func (data *Data) String() string {
	linkage := data.Linkage.String()
	var parts []string
	if linkage != "" {
		parts = append(parts, linkage)
	}
	parts = append(parts, "data", data.Name.String(), "=")
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
