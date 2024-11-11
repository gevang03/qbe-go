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

type CallInst struct {
	Dest     *Dest
	Env      *value.Temporary
	Variadic bool
	target   value.Value
	args     []arg
}

func (*CallInst) isInst() {}

func (inst CallInst) String() string {
	var builder strings.Builder
	if inst.Dest != nil {
		builder.WriteString(inst.Dest.Name.String())
		builder.WriteString(" =")
		builder.WriteString(fmt.Sprint(inst.Dest.Type))
		builder.WriteByte(' ')
	}
	builder.WriteString("call ")
	builder.WriteString(fmt.Sprint(inst.target))
	builder.WriteByte('(')
	if inst.Env != nil {
		builder.WriteString("env ")
		builder.WriteString(inst.Env.String())
		builder.WriteString(", ")
	}
	for _, arg := range inst.args {
		builder.WriteString(fmt.Sprint(arg.Type))
		builder.WriteByte(' ')
		builder.WriteString(fmt.Sprint(arg.Value))
		builder.WriteString(", ")
	}
	if inst.Variadic {
		builder.WriteString("...")
	}
	builder.WriteByte(')')
	return builder.String()
}

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

func (c *CallInst) InsertArg(type_ types.ABIType, value value.Value) {
	c.args = append(c.args, arg{type_, value})
}

func (b *Block) InsertVastart(ap value.Value) {
	inst := newSimpleInstNoDest(vastart, ap)
	b.insertInstruction(inst)
}

func (b *Block) InsertVaarg(dest value.Temporary, type_ types.BaseType, src value.Value) {
	inst := newSimpleInst(vaarg, dest, type_, src)
	b.insertInstruction(inst)
}
