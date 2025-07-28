// Package qbe implements an API to generate QBE IL files.
package qbe

import "fmt"

// A Module represents a single file of definitions
type Module struct {
	name        string
	definitions map[string]Definition
}

// NewModule returns a new [Module] that can be written to file name
func NewModule(name string) *Module {
	return &Module{
		name,
		make(map[string]Definition),
	}
}

func (mod *Module) insertDef(name fmt.Stringer, def Definition) {
	key := name.String()
	if _, exists := mod.definitions[key]; exists {
		panic(fmt.Sprintf("duplicate definition of `%v'", key))
	}
	mod.definitions[key] = def
}

// DefineStruct inserts into mod a reference to [Struct] with typename name.
func (mod *Module) DefineStruct(name TypeName) *Struct {
	s := newStruct(name)
	mod.insertDef(name, s)
	return s
}

// DefineUnion inserts into mod a reference to [Union] with typename name.
func (mod *Module) DefineUnion(name TypeName) *Union {
	u := newUnion(name)
	mod.insertDef(name, u)
	return u
}

// DefineOpaque inserts into mod a reference to [Opaque] with typename name,
// alignment align and size equal to size.
func (mod *Module) DefineOpaque(name TypeName, align uint, size uint) *Opaque {
	o := newOpaque(name, align, size)
	mod.insertDef(name, o)
	return o
}

// DefineData inserts into mod a reference to [Data] with symbol name.
func (mod *Module) DefineData(name GlobalSymbol) *Data {
	d := newData(name)
	mod.insertDef(name, d)
	return d
}

// DefineFunction inserts into mod a reference to [Function] with symbol name.
func (mod *Module) DefineFunction(name GlobalSymbol) *Function {
	f := newFunction(name)
	mod.insertDef(name, f)
	return f
}
