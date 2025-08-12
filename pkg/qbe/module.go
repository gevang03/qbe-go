// Package qbe implements an API to generate QBE IL files.
package qbe

import "fmt"

// A Module represents a single file of definitions
type Module struct {
	name        string
	typeDefs    map[TypeName]TypeDef
	definitions map[GlobalSymbol]Definition
}

// NewModule returns a new [Module] that can be written to file name
func NewModule(name string) *Module {
	return &Module{
		name,
		make(map[TypeName]TypeDef),
		make(map[GlobalSymbol]Definition),
	}
}

func (mod *Module) insertDef(name GlobalSymbol, def Definition) {
	if _, exists := mod.definitions[name]; exists {
		panic(fmt.Sprintf("duplicate definition of `%v'", name.String()))
	}
	mod.definitions[name] = def
}

func (mod *Module) insertTypeDef(name TypeName, def TypeDef) {
	if _, exists := mod.typeDefs[name]; exists {
		panic(fmt.Sprintf("duplicate type definition of `%v'", name.String()))
	}
	mod.typeDefs[name] = def
}

// DefineStruct inserts into mod a reference to [Struct] with typename name.
func (mod *Module) DefineStruct(name TypeName) *Struct {
	s := newStruct(name)
	mod.insertTypeDef(name, s)
	return s
}

// DefineUnion inserts into mod a reference to [Union] with typename name.
func (mod *Module) DefineUnion(name TypeName) *Union {
	u := newUnion(name)
	mod.insertTypeDef(name, u)
	return u
}

// DefineOpaque inserts into mod a reference to [Opaque] with typename name,
// alignment align and size equal to size.
func (mod *Module) DefineOpaque(name TypeName, align uint, size uint) *Opaque {
	o := newOpaque(name, align, size)
	mod.insertTypeDef(name, o)
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

// GetSymbol returns the [Definition] definition with symbol in mod iff ok is true.
func (mod *Module) GetSymbol(symbol GlobalSymbol) (def Definition, ok bool) {
	def, ok = mod.definitions[symbol]
	return
}

// GetTypeDef returns the [TypeDef] typeDef with name in mod iff ok is true.
func (mod *Module) GetTypeDef(name TypeName) (typeDef TypeDef, ok bool) {
	typeDef, ok = mod.typeDefs[name]
	return
}

// Name returns the name of mod.
func (mod *Module) Name() string {
	return mod.name
}
