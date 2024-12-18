package inst

import (
	"github.com/gevang03/qbe-go/types"
	"github.com/gevang03/qbe-go/value"
)

// InsertAdd adds an add instruction at the end of b, storing the sum of
// src1 and src2 to dest with type type_.
func (b *Block) InsertAdd(dest value.Temporary, type_ types.BaseType, src1, src2 value.Value) {
	inst := newSimpleInst(add, dest, type_, src1, src2)
	b.insertInstruction(inst)
}

// InsertSub adds a sub instruction at the end of b, storing the difference of
// src1 and src2 to dest with type type_.
func (b *Block) InsertSub(dest value.Temporary, type_ types.BaseType, src1, src2 value.Value) {
	inst := newSimpleInst(sub, dest, type_, src1, src2)
	b.insertInstruction(inst)
}

// InsertMul adds a mul instruction at the end of b, storing the product of
// src1 and src2 to dest with type type_.
func (b *Block) InsertMul(dest value.Temporary, type_ types.BaseType, src1, src2 value.Value) {
	inst := newSimpleInst(mul, dest, type_, src1, src2)
	b.insertInstruction(inst)
}

// InsertDiv adds a div instruction at the end of b, storing the quotient of
// src1 and src2 to dest with type type_.
func (b *Block) InsertDiv(dest value.Temporary, type_ types.BaseType, src1, src2 value.Value) {
	inst := newSimpleInst(div, dest, type_, src1, src2)
	b.insertInstruction(inst)
}

// InsertNeg adds a neg instruction at the end of b, storing the negative of
// src to dest with type type_.
func (b *Block) InsertNeg(dest value.Temporary, type_ types.BaseType, src value.Value) {
	inst := newSimpleInst(neg, dest, type_, src)
	b.insertInstruction(inst)
}

// InsertUdiv adds a udiv instruction at the end of b, storing the quotient of
// unsigned src1 and src2 to dest with type type_.
func (b *Block) InsertUdiv(dest value.Temporary, type_ types.IntegralType, src1, src2 value.Value) {
	inst := newSimpleInst(udiv, dest, type_, src1, src2)
	b.insertInstruction(inst)
}

// InsertRem adds a rem instruction at the end of b, storing the remainder of
// src1 and src2 to dest with type type_.
func (b *Block) InsertRem(dest value.Temporary, type_ types.IntegralType, src1, src2 value.Value) {
	inst := newSimpleInst(rem, dest, type_, src1, src2)
	b.insertInstruction(inst)
}

// InsertUrem adds a urem instruction at the end of b, storing the remainder of
// unsigned src1 and src2 to dest with type type_.
func (b *Block) InsertUrem(dest value.Temporary, type_ types.IntegralType, src1, src2 value.Value) {
	inst := newSimpleInst(urem, dest, type_, src1, src2)
	b.insertInstruction(inst)
}
