package qbe

import (
	"bufio"
	"fmt"

	"github.com/gevang03/qbe-go/def"
	"github.com/gevang03/qbe-go/types"
	"github.com/gevang03/qbe-go/value"
)

type Module struct {
	name        string
	definitions map[string]def.Definition
}

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

func (mod *Module) DefineStruct(name types.TypeName) *def.Struct {
	s := def.NewStruct(name)
	mod.insertDef(name, s)
	return s
}

func (mod *Module) DefineUnion(name types.TypeName) *def.Union {
	u := def.NewUnion(name)
	mod.insertDef(name, u)
	return u
}

func (mod *Module) DefineOpaque(name types.TypeName, align uint, size uint) *def.Opaque {
	o := def.NewOpaque(name, align, size)
	mod.insertDef(name, o)
	return o
}

func (mod *Module) DefineData(name value.GlobalSymbol) *def.Data {
	d := &def.Data{Linkage: def.PrivateLinkage(), Name: name, Align: 0, Fields: nil}
	mod.insertDef(name, d)
	return d
}

func (mod *Module) DefineFunction(name value.GlobalSymbol) *def.Function {
	f := def.NewFunction(name)
	mod.insertDef(name, f)
	return f
}

func (mod *Module) ToIL(w *bufio.Writer) (int, error) {
	written := 0
	for _, def := range mod.definitions {
		count, err := w.WriteString(fmt.Sprint(def))
		written += count
		if err != nil {
			return written, err
		}
		if err = w.WriteByte('\n'); err != nil {
			return written, err
		}
	}
	return written, nil
}
