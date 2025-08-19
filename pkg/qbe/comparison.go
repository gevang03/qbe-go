package qbe

import "fmt"

func dispatchIntegral(type_ IntegralType, opWord, opLong opcode) opcode {
	switch type_ {
	case WordType():
		return opWord
	case LongType():
		return opLong
	default:
		panic(fmt.Sprintf("unexpected type %#v", type_))
	}
}

func dispatchFloating(type_ FloatingPointType, opSingle, opDouble opcode) opcode {
	switch type_ {
	case SingleType():
		return opSingle
	case DoubleType():
		return opDouble
	default:
		panic(fmt.Sprintf("unexpected type %#v", type_))
	}
}

func dispatchBase(type_ BaseType, opWord, opLong, opSingle, opDouble opcode) opcode {
	switch type_ {
	case WordType():
		return opWord
	case LongType():
		return opLong
	case SingleType():
		return opSingle
	case DoubleType():
		return opDouble
	default:
		panic(fmt.Sprintf("unexpected type %#v", type_))
	}
}

// InsertCeq adds a ceq instruction at the end of b, comparing values src1 and
// src2 of srcType type and storing the result to dest with destType type.
func (b *Block) InsertCeq(dest Temporary, destType IntegralType,
	src1, src2 Value, srcType BaseType) {
	op := dispatchBase(srcType, ceqw, ceql, ceqs, ceqd)
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

// InsertCge adds a cge instruction at the end of b, comparing floating point
// values src1 and src2 of srcType type and storing the result to dest with
// destType type.
func (b *Block) InsertCge(dest Temporary, destType IntegralType,
	src1, src2 Value, srcType FloatingPointType) {
	op := dispatchFloating(srcType, cges, cged)
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

// InsertCgt adds a cgt instruction at the end of b, comparing floating point
// values src1 and src2 of srcType type and storing the result to dest with
// destType type.
func (b *Block) InsertCgt(dest Temporary, destType IntegralType,
	src1, src2 Value, srcType FloatingPointType) {
	op := dispatchFloating(srcType, cgts, cgtd)
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

// InsertCle adds a cle instruction at the end of b, comparing floating point
// values src1 and src2 of srcType type and storing the result to dest with
// destType type.
func (b *Block) InsertCle(dest Temporary, destType IntegralType,
	src1, src2 Value, srcType FloatingPointType) {
	op := dispatchFloating(srcType, cles, cled)
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

// InsertClt adds a clt instruction at the end of b, comparing floating point
// values src1 and src2 of srcType type and storing the result to dest with
// destType type.
func (b *Block) InsertClt(dest Temporary, destType IntegralType,
	src1, src2 Value, srcType FloatingPointType) {
	op := dispatchFloating(srcType, clts, cltd)
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

// InsertCne adds a cne instruction at the end of b, comparing values src1 and
// src2 of srcType type and storing the result to dest with destType type.
func (b *Block) InsertCne(dest Temporary, destType IntegralType,
	src1, src2 Value, srcType BaseType) {
	op := dispatchBase(srcType, cnew, cnel, cnes, cned)
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

// InsertCo adds a co instruction at the end of b, comparing floating point
// values src1 and src2 of srcType type and storing the result to dest with
// destType type.
func (b *Block) InsertCo(dest Temporary, destType IntegralType,
	src1, src2 Value, srcType FloatingPointType) {
	op := dispatchFloating(srcType, cos, cod)
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

// InsertCsge adds a csge instruction at the end of b, comparing integral
// values src1 and src2 of srcType type and storing the result to dest with
// destType type.
func (b *Block) InsertCsge(dest Temporary, destType IntegralType,
	src1, src2 Value, srcType IntegralType) {
	op := dispatchIntegral(srcType, csgew, csgel)
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

// InsertCsgt adds a csgt instruction at the end of b, comparing integral
// values src1 and src2 of srcType type and storing the result to dest with
// destType type.
func (b *Block) InsertCsgt(dest Temporary, destType IntegralType,
	src1, src2 Value, srcType IntegralType) {
	op := dispatchIntegral(srcType, csgtw, csgtl)
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

// InsertCsle adds a csle instruction at the end of b, comparing integral
// values src1 and src2 of srcType type and storing the result to dest with
// destType type.
func (b *Block) InsertCsle(dest Temporary, destType IntegralType,
	src1, src2 Value, srcType IntegralType) {
	op := dispatchIntegral(srcType, cslew, cslel)
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

// InsertCslt adds a cslt instruction at the end of b, comparing integral
// values src1 and src2 of srcType type and storing the result to dest with
// destType type.
func (b *Block) InsertCslt(dest Temporary, destType IntegralType,
	src1, src2 Value, srcType IntegralType) {
	op := dispatchIntegral(srcType, csltw, csltl)
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

// InsertCuge adds a cuge instruction at the end of b, comparing integral
// values src1 and src2 of srcType type and storing the result to dest with
// destType type.
func (b *Block) InsertCuge(dest Temporary, destType IntegralType,
	src1, src2 Value, srcType IntegralType) {
	op := dispatchIntegral(srcType, cugew, cugel)
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

// InsertCugt adds a cugt instruction at the end of b, comparing integral
// values src1 and src2 of srcType type and storing the result to dest with
// destType type.
func (b *Block) InsertCugt(dest Temporary, destType IntegralType,
	src1, src2 Value, srcType IntegralType) {
	op := dispatchIntegral(srcType, cugtw, cugtl)
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

// InsertCule adds a cule instruction at the end of b, comparing integral
// values src1 and src2 of srcType type and storing the result to dest with
// destType type.
func (b *Block) InsertCule(dest Temporary, destType IntegralType,
	src1, src2 Value, srcType IntegralType) {
	op := dispatchIntegral(srcType, culew, culel)
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

// InsertCult adds a cult instruction at the end of b, comparing integral
// values src1 and src2 of srcType type and storing the result to dest with
// destType type.
func (b *Block) InsertCult(dest Temporary, destType IntegralType,
	src1, src2 Value, srcType IntegralType) {
	op := dispatchIntegral(srcType, cultw, cultl)
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

// InsertCuo adds a cuo instruction at the end of b, comparing floating point
// values src1 and src2 of srcType type and storing the result to dest with
// destType type.
func (b *Block) InsertCuo(dest Temporary, destType IntegralType,
	src1, src2 Value, srcType FloatingPointType) {
	op := dispatchFloating(srcType, cuos, cuod)
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}
