package def

import (
	"fmt"
	"strings"

	"github.com/gevang03/qbe-go/types"
	"github.com/gevang03/qbe-go/value"
)

type Data struct {
	Linkage
	Name   value.GlobalSymbol
	Align  uint
	Fields []dataField
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

func (data *Data) InsertValue(type_ types.ExtendedType, item ...value.DataItem) {
	data.Fields = append(data.Fields, dataFieldValue{type_, item})
}

func (data *Data) InsertZeroes(count uint) {
	data.Fields = append(data.Fields, dataFieldZeroes(count))
}

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
	for _, field := range data.Fields {
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
			parts = append(parts, "z")
			parts = append(parts, fmt.Sprintf("%v,", f))
		default:
			panic("unreachable")
		}
	}
	parts = append(parts, "}")
	return strings.Join(parts, " ")
}
