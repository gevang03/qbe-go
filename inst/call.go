package inst

import (
	"fmt"
	"strings"

	"github.com/gevang03/qbe-go/types"
	"github.com/gevang03/qbe-go/value"
)

type arg struct {
	Type types.ABIType
	value.Value
}

// A CallInst represents a call instruction.
type CallInst struct {
	Dest     *Dest            // Where the result of the call is stored, nil if no result.
	Env      *value.Temporary // The env parameter, nil if not passed.
	Variadic bool             // Set to true if the function called is variadic.
	target   value.Value
	args     []arg
}

func (*CallInst) isInst() {}

// String converts inst to a string compatible with QBE code.
func (c *CallInst) String() string {
	var builder strings.Builder
	if c.Dest != nil {
		builder.WriteString(c.Dest.Name.String())
		builder.WriteString(" =")
		builder.WriteString(fmt.Sprint(c.Dest.Type))
		builder.WriteByte(' ')
	}
	builder.WriteString("call ")
	builder.WriteString(fmt.Sprint(c.target))
	builder.WriteByte('(')
	if c.Env != nil {
		builder.WriteString("env ")
		builder.WriteString(c.Env.String())
		builder.WriteString(", ")
	}
	for _, arg := range c.args {
		builder.WriteString(fmt.Sprint(arg.Type))
		builder.WriteByte(' ')
		builder.WriteString(fmt.Sprint(arg.Value))
		builder.WriteString(", ")
	}
	if c.Variadic {
		builder.WriteString("...")
	}
	builder.WriteByte(')')
	return builder.String()
}

// InsertCall adds a call instruction to b with callee target and
// returns a pointer to the generated CallInst.
func (b *Block) InsertCall(target value.Value) *CallInst {
	inst := &CallInst{
		Dest:     nil,
		Env:      nil,
		Variadic: false,
		target:   target,
		args:     nil,
	}
	b.insertInstruction(inst)
	return b.insts[len(b.insts)-1].(*CallInst)
}

// InsertArg adds argument value with type type_ to the end of the argument list of c.
func (c *CallInst) InsertArg(type_ types.ABIType, value value.Value) {
	c.args = append(c.args, arg{type_, value})
}

// InsertVastart adds a vastart at the end of b with argument ap.
func (b *Block) InsertVastart(ap value.Value) {
	inst := newSimpleInstNoDest(vastart, ap)
	b.insertInstruction(inst)
}

// InsertVaarg adds a vaarg at the end of b with source ap and storing the argument to dest with type type_.
func (b *Block) InsertVaarg(dest value.Temporary, type_ types.BaseType, ap value.Value) {
	inst := newSimpleInst(vaarg, dest, type_, ap)
	b.insertInstruction(inst)
}
