package inst

import (
	"github.com/gevang03/qbe-go/types"
	"github.com/gevang03/qbe-go/value"
)

func (b *Block) InsertOr(dest value.Temporary, type_ types.IntegralType, src1, src2 value.Value) {
	inst := newSimpleInst(or, dest, type_, src1, src2)
	b.insertInstruction(inst)
}

func (b *Block) InsertXor(dest value.Temporary, type_ types.IntegralType, src1, src2 value.Value) {
	inst := newSimpleInst(xor, dest, type_, src1, src2)
	b.insertInstruction(inst)
}

func (b *Block) InsertAnd(dest value.Temporary, type_ types.IntegralType, src1, src2 value.Value) {
	inst := newSimpleInst(and, dest, type_, src1, src2)
	b.insertInstruction(inst)
}

func (b *Block) InsertSar(dest value.Temporary, type_ types.IntegralType, src1, src2 value.Value) {
	inst := newSimpleInst(sar, dest, type_, src1, src2)
	b.insertInstruction(inst)
}

func (b *Block) InsertShr(dest value.Temporary, type_ types.IntegralType, src1, src2 value.Value) {
	inst := newSimpleInst(shr, dest, type_, src1, src2)
	b.insertInstruction(inst)
}

func (b *Block) InsertShl(dest value.Temporary, type_ types.IntegralType, src1, src2 value.Value) {
	inst := newSimpleInst(shl, dest, type_, src1, src2)
	b.insertInstruction(inst)
}
