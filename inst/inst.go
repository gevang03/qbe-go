package inst

import (
	"fmt"
	"strings"

	"github.com/gevang03/qbe-go/types"
	"github.com/gevang03/qbe-go/value"
)

type Dest struct {
	Name value.Temporary
	Type types.ABIType
}

type simpleInst struct {
	opcode
	*Dest
	srcs []value.Value
}

type inst interface{ isInst() }

func (simpleInst) isInst() {}

func (inst simpleInst) String() string {
	var parts []string
	if inst.Dest != nil {
		parts = append(parts, inst.Dest.Name.String())
		parts = append(parts, fmt.Sprintf("=%v", inst.Dest.Type))
	}
	parts = append(parts, inst.opcode.String())
	for i, value := range inst.srcs {
		if i < len(inst.srcs)-1 {
			parts = append(parts, fmt.Sprintf("%v,", value))
		} else {
			parts = append(parts, fmt.Sprint(value))
		}
	}
	return strings.Join(parts, " ")
}

func newSimpleInst(opcode opcode, name value.Temporary, type_ types.BaseType, srcs ...value.Value) simpleInst {
	return simpleInst{opcode, &Dest{name, type_}, srcs}
}

func newSimpleInstNoDest(opcode opcode, srcs ...value.Value) simpleInst {
	return simpleInst{opcode, nil, srcs}
}

func (b *Block) insertInstruction(inst inst) {
	b.insts = append(b.insts, inst)
}
