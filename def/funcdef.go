package def

// Convention: temporaries of the form /\.t\d+/ and labels of the form
// /\.L\d+/ are reserved for automatic name generation.

import (
	"fmt"
	"strings"

	"github.com/gevang03/qbe-go/inst"
	"github.com/gevang03/qbe-go/types"
	"github.com/gevang03/qbe-go/value"
)

type param struct {
	Type types.ABIType
	Name value.Temporary
}

type Function struct {
	Linkage
	RetType  types.ABIType
	Name     value.GlobalSymbol
	Env      *value.Temporary
	params   []param
	Variadic bool
	blocks   []inst.Block
	labelGen uint
	tmpGen   uint
}

func (f *Function) isDefinition() {}

func NewFunction(name value.GlobalSymbol) *Function {
	return &Function{
		Linkage:  PrivateLinkage(),
		RetType:  nil,
		Name:     name,
		Env:      nil,
		params:   nil,
		Variadic: false,
		blocks:   nil,
		labelGen: 0,
		tmpGen:   0,
	}
}

func (f *Function) InsertParam(type_ types.ABIType, name value.Temporary) {
	f.params = append(f.params, param{type_, name})
}

func (f *Function) InsertBlock(label inst.Label) *inst.Block {
	f.blocks = append(f.blocks, *inst.NewBlock(label))
	return &f.blocks[len(f.blocks)-1]
}

func (f *Function) InsertBlockAuto() *inst.Block {
	label := inst.Label(fmt.Sprintf(".L%v", f.labelGen))
	f.labelGen++
	return f.InsertBlock(label)
}

func (f *Function) NewTemporary() value.Temporary {
	tmp := value.Temporary(fmt.Sprintf(".t%v", f.tmpGen))
	f.tmpGen++
	return tmp
}

func (f *Function) String() string {
	builder := strings.Builder{}
	linkage := f.Linkage.String()
	if linkage != "" {
		builder.WriteString(linkage)
		builder.WriteByte(' ')
	}
	builder.WriteString("function ")
	if f.RetType != nil {
		builder.WriteString(fmt.Sprint(f.RetType))
		builder.WriteByte(' ')
	}
	builder.WriteString(f.Name.String())
	builder.WriteByte('(')
	if f.Env != nil {
		builder.WriteString("env ")
		builder.WriteString(f.Env.String())
		builder.WriteString(", ")
	}
	for _, param := range f.params {
		builder.WriteString(fmt.Sprint(param.Type))
		builder.WriteByte(' ')
		builder.WriteString(param.Name.String())
		builder.WriteString(", ")
	}
	if f.Variadic {
		builder.WriteString("...")
	}
	builder.WriteString(") {\n")
	for _, block := range f.blocks {
		builder.WriteString(block.String())
	}
	builder.WriteString("}")
	return builder.String()
}
