package qbe

// Convention: temporaries of the form /%<name>.\d+/ and labels of the form
// /@<name>.\d+/ are reserved for automatic name generation.

import (
	"fmt"
	"strings"
)

type param struct {
	Type ABIType
	Name Temporary
}

// A Function represents a function definition in QBE IL.
type Function struct {
	Linkage               // The linkage of the function, cannot be thread
	RetType  ABIType      // The return type of the function
	name     GlobalSymbol // The symbol that references the function
	Env      *Temporary   // Parameter used to implement closures
	params   []param
	Variadic bool // Set if function is variadic.
	blocks   []*Block
	labelGen uint
	tmpGen   uint
}

func (f *Function) isDefinition() {}

// newFunction returns a new [Function] with function name name, private linkage,
// no return type, no env or any other parameters and it is not set as variadic.
func newFunction(name GlobalSymbol) *Function {
	return &Function{
		Linkage:  PrivateLinkage(),
		RetType:  nil,
		name:     name,
		Env:      nil,
		params:   nil,
		Variadic: false,
		blocks:   nil,
		labelGen: 0,
		tmpGen:   0,
	}
}

// Name returns the name of f.
func (f *Function) Name() GlobalSymbol {
	return f.name
}

// InsertParam inserts at the end of the parameter list of f a new parameter named name with type type_.
func (f *Function) InsertParam(type_ ABIType, name Temporary) {
	f.params = append(f.params, param{type_, name})
}

// InsertBlock inserts a new [Block] at the end of the function body,
// with label as the name of its entry point. Returns a reference to that block.
func (f *Function) InsertBlock(label Label) *Block {
	f.blocks = append(f.blocks, newBlock(label))
	return f.blocks[len(f.blocks)-1]
}

// InsertBlockAuto inserts a new [Block] at the end of the function body,
// with an auto-generated label of the form /@<name>\.\d+/. Refrain from creating
// labels of this form and using this function to ensure uniqueness.
// Returns a reference to that block.
func (f *Function) InsertBlockAuto(name string) *Block {
	label := Label(fmt.Sprintf("%v.%v", name, f.labelGen))
	f.labelGen++
	return f.InsertBlock(label)
}

// NewTemporary returns a new [Temporary] of the form /%<name>\.\d+/. Refrain
// from creating temporaries of this form and using this function to ensure uniqueness.
func (f *Function) NewTemporary(name string) Temporary {
	tmp := Temporary(fmt.Sprintf("%v.%v", name, f.tmpGen))
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
		builder.WriteString(f.RetType.Name())
		builder.WriteByte(' ')
	}
	builder.WriteString(f.name.String())
	builder.WriteByte('(')
	if f.Env != nil {
		builder.WriteString("env ")
		builder.WriteString(f.Env.String())
		builder.WriteString(", ")
	}
	for _, param := range f.params {
		builder.WriteString(param.Type.Name())
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
