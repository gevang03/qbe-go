package def

// Convention: temporaries of the form /%\.t\d+/ and labels of the form
// /@\.L\d+/ are reserved for automatic name generation.

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

// A Function represents a function definition in QBE IL.
type Function struct {
	Linkage                     // The linkage of the function, cannot be thread
	RetType  types.ABIType      // The return type of the function
	Name     value.GlobalSymbol // The symbol that references the function
	Env      *value.Temporary   // Parameter used to implement closures
	params   []param
	Variadic bool // Set if function is variadic.
	blocks   []*inst.Block
	labelGen uint
	tmpGen   uint
}

func (f *Function) isDefinition() {}

// NewFunction returns a new [Function] with function name name, private linkage,
// no return type, no env or any other parameters and it is not set as variadic.
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

// InsertParam inserts at the end of the parameter list of f a new parameter named name with type type_.
func (f *Function) InsertParam(type_ types.ABIType, name value.Temporary) {
	f.params = append(f.params, param{type_, name})
}

// InsertBlock inserts a new [inst.Block] at the end of the function body,
// with label as the name of its entry point. Returns a reference to that block.
func (f *Function) InsertBlock(label inst.Label) *inst.Block {
	f.blocks = append(f.blocks, inst.NewBlock(label))
	return f.blocks[len(f.blocks)-1]
}

// InsertBlock inserts a new [inst.Block] at the end of the function body,
// with an auto-generated label of the form /@\.L\d+/. Refrain from creating
// labels of this form and using this function to ensure uniqueness.
// Returns a reference to that block.
func (f *Function) InsertBlockAuto() *inst.Block {
	label := inst.Label(fmt.Sprintf(".L%v", f.labelGen))
	f.labelGen++
	return f.InsertBlock(label)
}

// NewTemporary returns a new [value.Temporary] of the form /%.t\d+/. Refrain
// from creating temporaries of this form and using this function to ensure uniqueness.
func (f *Function) NewTemporary() value.Temporary {
	tmp := value.Temporary(fmt.Sprintf(".t%v", f.tmpGen))
	f.tmpGen++
	return tmp
}

// String converts f to a string compatible with QBE code.
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
