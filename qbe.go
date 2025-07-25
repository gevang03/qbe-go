// Package qbe implements an API to generate QBE IL files.
package qbe

import (
	"fmt"

	"github.com/gevang03/qbe-go/def"
	"github.com/gevang03/qbe-go/types"
	"github.com/gevang03/qbe-go/value"
)

// A Module represents a single file of definitions
type Module struct {
	name        string
	definitions map[string]def.Definition
}

// NewModule returns a new [Module] that can be written to file name
func NewModule(name string) *Module {
	return &Module{
		name,
		make(map[string]def.Definition),
	}
}

func (mod *Module) insertDef(name fmt.Stringer, def def.Definition) {
	key := name.String()
	if _, exists := mod.definitions[key]; exists {
		panic(fmt.Sprintf("duplicate definition of `%v'", key))
	}
	mod.definitions[key] = def
}

// DefineStruct inserts into mod a reference to [def.Struct] with typename name.
func (mod *Module) DefineStruct(name types.TypeName) *def.Struct {
	s := def.NewStruct(name)
	mod.insertDef(name, s)
	return s
}

// DefineUnion inserts into mod a reference to [def.Union] with typename name.
func (mod *Module) DefineUnion(name types.TypeName) *def.Union {
	u := def.NewUnion(name)
	mod.insertDef(name, u)
	return u
}

// DefineOpaque inserts into mod a reference to [def.Opaque] with typename name,
// alignment align and size equal to size.
func (mod *Module) DefineOpaque(name types.TypeName, align uint, size uint) *def.Opaque {
	o := def.NewOpaque(name, align, size)
	mod.insertDef(name, o)
	return o
}

// DefineData inserts into mod a reference to [def.Data] with symbol name.
func (mod *Module) DefineData(name value.GlobalSymbol) *def.Data {
	d := def.NewData(name)
	mod.insertDef(name, d)
	return d
}

// DefineFunction inserts into mod a reference to [def.Function] with symbol name.
func (mod *Module) DefineFunction(name value.GlobalSymbol) *def.Function {
	f := def.NewFunction(name)
	mod.insertDef(name, f)
	return f
}
