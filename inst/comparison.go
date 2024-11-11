package inst

import (
	"github.com/gevang03/qbe-go/types"
	"github.com/gevang03/qbe-go/value"
)

func (b *Block) InsertCeq(dest value.Temporary, destType types.IntegralType,
	src1, src2 value.Value, srcType types.BaseType) {
	var op opcode
	if srcType == types.Double() {
		op = ceqd
	} else if srcType == types.Single() {
		op = ceqs
	} else if srcType == types.Long() {
		op = ceql
	} else if srcType == types.Word() {
		op = ceqw
	} else {
		panic("unreachable")
	}
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

func (b *Block) InsertCge(dest value.Temporary, destType types.IntegralType,
	src1, src2 value.Value, srcType types.FloatingPointType) {
	var op opcode
	if srcType == types.Double() {
		op = cged
	} else if srcType == types.Single() {
		op = cges
	} else {
		panic("unreachable")
	}
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

func (b *Block) InsertCgt(dest value.Temporary, destType types.IntegralType,
	src1, src2 value.Value, srcType types.FloatingPointType) {
	var op opcode
	if srcType == types.Double() {
		op = cgtd
	} else if srcType == types.Single() {
		op = cgts
	} else {
		panic("unreachable")
	}
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

func (b *Block) InsertCle(dest value.Temporary, destType types.IntegralType,
	src1, src2 value.Value, srcType types.FloatingPointType) {
	var op opcode
	if srcType == types.Double() {
		op = cled
	} else if srcType == types.Single() {
		op = cles
	} else {
		panic("unreachable")
	}
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

func (b *Block) InsertClt(dest value.Temporary, destType types.IntegralType,
	src1, src2 value.Value, srcType types.FloatingPointType) {
	var op opcode
	if srcType == types.Double() {
		op = cltd
	} else if srcType == types.Single() {
		op = clts
	} else {
		panic("unreachable")
	}
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

func (b *Block) InsertCne(dest value.Temporary, destType types.IntegralType,
	src1, src2 value.Value, srcType types.BaseType) {
	var op opcode
	if srcType == types.Double() {
		op = cned
	} else if srcType == types.Single() {
		op = cnes
	} else if srcType == types.Long() {
		op = cnel
	} else if srcType == types.Word() {
		op = cnew
	} else {
		panic("unreachable")
	}
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

func (b *Block) InsertCo(dest value.Temporary, destType types.IntegralType,
	src1, src2 value.Value, srcType types.FloatingPointType) {
	var op opcode
	if srcType == types.Double() {
		op = cod
	} else if srcType == types.Single() {
		op = cos
	} else {
		panic("unreachable")
	}
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

func (b *Block) InsertCsge(dest value.Temporary, destType types.IntegralType,
	src1, src2 value.Value, srcType types.IntegralType) {
	var op opcode
	if srcType == types.Long() {
		op = csgel
	} else if srcType == types.Word() {
		op = csgew
	} else {
		panic("unreachable")
	}
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

func (b *Block) InsertCsgt(dest value.Temporary, destType types.IntegralType,
	src1, src2 value.Value, srcType types.IntegralType) {
	var op opcode
	if srcType == types.Long() {
		op = csgtl
	} else if srcType == types.Word() {
		op = csgtw
	} else {
		panic("unreachable")
	}
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

func (b *Block) InsertCsle(dest value.Temporary, destType types.IntegralType,
	src1, src2 value.Value, srcType types.IntegralType) {
	var op opcode
	if srcType == types.Long() {
		op = cslel
	} else if srcType == types.Word() {
		op = cslew
	} else {
		panic("unreachable")
	}
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

func (b *Block) InsertCslt(dest value.Temporary, destType types.IntegralType,
	src1, src2 value.Value, srcType types.IntegralType) {
	var op opcode
	if srcType == types.Long() {
		op = csltl
	} else if srcType == types.Word() {
		op = csltw
	} else {
		panic("unreachable")
	}
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}
func (b *Block) InsertCuge(dest value.Temporary, destType types.IntegralType,
	src1, src2 value.Value, srcType types.IntegralType) {
	var op opcode
	if srcType == types.Long() {
		op = cugel
	} else if srcType == types.Word() {
		op = cugew
	} else {
		panic("unreachable")
	}
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

func (b *Block) InsertCugt(dest value.Temporary, destType types.IntegralType,
	src1, src2 value.Value, srcType types.IntegralType) {
	var op opcode
	if srcType == types.Long() {
		op = cugtl
	} else if srcType == types.Word() {
		op = cugtw
	} else {
		panic("unreachable")
	}
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

func (b *Block) InsertCule(dest value.Temporary, destType types.IntegralType,
	src1, src2 value.Value, srcType types.IntegralType) {
	var op opcode
	if srcType == types.Long() {
		op = culel
	} else if srcType == types.Word() {
		op = culew
	} else {
		panic("unreachable")
	}
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

func (b *Block) InsertCult(dest value.Temporary, destType types.IntegralType,
	src1, src2 value.Value, srcType types.IntegralType) {
	var op opcode
	if srcType == types.Long() {
		op = cultl
	} else if srcType == types.Word() {
		op = cultw
	} else {
		panic("unreachable")
	}
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

func (b *Block) InsertCuo(dest value.Temporary, destType types.IntegralType,
	src1, src2 value.Value, srcType types.FloatingPointType) {
	var op opcode
	if srcType == types.Double() {
		op = cuod
	} else if srcType == types.Single() {
		op = cuos
	} else {
		panic("unreachable")
	}
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}
