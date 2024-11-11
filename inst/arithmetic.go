package inst

import (
	"github.com/gevang03/qbe-go/types"
	"github.com/gevang03/qbe-go/value"
)

func (b *Block) InsertAdd(dest value.Temporary, type_ types.BaseType, src1, src2 value.Value) {
	inst := newSimpleInst(add, dest, type_, src1, src2)
	b.insertInstruction(inst)
}

func (b *Block) InsertSub(dest value.Temporary, type_ types.BaseType, src1, src2 value.Value) {
	inst := newSimpleInst(sub, dest, type_, src1, src2)
	b.insertInstruction(inst)
}

func (b *Block) InsertMul(dest value.Temporary, type_ types.BaseType, src1, src2 value.Value) {
	inst := newSimpleInst(mul, dest, type_, src1, src2)
	b.insertInstruction(inst)
}

func (b *Block) InsertDiv(dest value.Temporary, type_ types.BaseType, src1, src2 value.Value) {
	inst := newSimpleInst(div, dest, type_, src1, src2)
	b.insertInstruction(inst)
}

func (b *Block) InsertNeg(dest value.Temporary, type_ types.BaseType, src value.Value) {
	inst := newSimpleInst(add, dest, type_, src)
	b.insertInstruction(inst)
}

func (b *Block) InsertUdiv(dest value.Temporary, type_ types.IntegralType, src1, src2 value.Value) {
	inst := newSimpleInst(udiv, dest, type_, src1, src2)
	b.insertInstruction(inst)
}

func (b *Block) InsertRem(dest value.Temporary, type_ types.IntegralType, src1, src2 value.Value) {
	inst := newSimpleInst(rem, dest, type_, src1, src2)
	b.insertInstruction(inst)
}

func (b *Block) InsertUrem(dest value.Temporary, type_ types.IntegralType, src1, src2 value.Value) {
	inst := newSimpleInst(urem, dest, type_, src1, src2)
	b.insertInstruction(inst)
}
