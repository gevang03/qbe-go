package qbe

import (
	"fmt"
	"strings"
)

// A Dest represents the destination to store the result of an instruction.
type Dest struct {
	Name Temporary // The name of the destination.
	Type ABIType   // The type of the destination.
}

type simpleInst struct {
	opcode
	*Dest
	srcs []Value
}

type inst interface{ isInst() }

func (simpleInst) isInst() {}

func (inst simpleInst) String() string {
	var parts []string
	if inst.Dest != nil {
		parts = append(parts, inst.Name.String())
		parts = append(parts, fmt.Sprintf("=%v", inst.Type))
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

func newSimpleInst(opcode opcode, name Temporary, type_ BaseType, srcs ...Value) simpleInst {
	return simpleInst{opcode, &Dest{name, type_}, srcs}
}

func newSimpleInstNoDest(opcode opcode, srcs ...Value) simpleInst {
	return simpleInst{opcode, nil, srcs}
}

func (b *Block) insertInstruction(inst inst) {
	b.insts = append(b.insts, inst)
}
