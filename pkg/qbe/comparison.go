package qbe

// InsertCeq adds a ceq instruction at the end of b, comparing values src1 and
// src2 of srcType type and storing the result to dest with destType type.
func (b *Block) InsertCeq(dest Temporary, destType IntegralType,
	src1, src2 Value, srcType BaseType) {
	var op opcode
	if srcType == DoubleType() {
		op = ceqd
	} else if srcType == SingleType() {
		op = ceqs
	} else if srcType == LongType() {
		op = ceql
	} else if srcType == WordType() {
		op = ceqw
	} else {
		panic("unreachable")
	}
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

// InsertCge adds a cge instruction at the end of b, comparing floating point
// values src1 and src2 of srcType type and storing the result to dest with
// destType type.
func (b *Block) InsertCge(dest Temporary, destType IntegralType,
	src1, src2 Value, srcType FloatingPointType) {
	var op opcode
	if srcType == DoubleType() {
		op = cged
	} else if srcType == SingleType() {
		op = cges
	} else {
		panic("unreachable")
	}
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

// InsertCgt adds a cgt instruction at the end of b, comparing floating point
// values src1 and src2 of srcType type and storing the result to dest with
// destType type.
func (b *Block) InsertCgt(dest Temporary, destType IntegralType,
	src1, src2 Value, srcType FloatingPointType) {
	var op opcode
	if srcType == DoubleType() {
		op = cgtd
	} else if srcType == SingleType() {
		op = cgts
	} else {
		panic("unreachable")
	}
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

// InsertCle adds a cle instruction at the end of b, comparing floating point
// values src1 and src2 of srcType type and storing the result to dest with
// destType type.
func (b *Block) InsertCle(dest Temporary, destType IntegralType,
	src1, src2 Value, srcType FloatingPointType) {
	var op opcode
	if srcType == DoubleType() {
		op = cled
	} else if srcType == SingleType() {
		op = cles
	} else {
		panic("unreachable")
	}
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

// InsertClt adds a clt instruction at the end of b, comparing floating point
// values src1 and src2 of srcType type and storing the result to dest with
// destType type.
func (b *Block) InsertClt(dest Temporary, destType IntegralType,
	src1, src2 Value, srcType FloatingPointType) {
	var op opcode
	if srcType == DoubleType() {
		op = cltd
	} else if srcType == SingleType() {
		op = clts
	} else {
		panic("unreachable")
	}
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

// InsertCne adds a cne instruction at the end of b, comparing values src1 and
// src2 of srcType type and storing the result to dest with destType type.
func (b *Block) InsertCne(dest Temporary, destType IntegralType,
	src1, src2 Value, srcType BaseType) {
	var op opcode
	if srcType == DoubleType() {
		op = cned
	} else if srcType == SingleType() {
		op = cnes
	} else if srcType == LongType() {
		op = cnel
	} else if srcType == WordType() {
		op = cnew
	} else {
		panic("unreachable")
	}
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

// InsertCo adds a co instruction at the end of b, comparing floating point
// values src1 and src2 of srcType type and storing the result to dest with
// destType type.
func (b *Block) InsertCo(dest Temporary, destType IntegralType,
	src1, src2 Value, srcType FloatingPointType) {
	var op opcode
	if srcType == DoubleType() {
		op = cod
	} else if srcType == SingleType() {
		op = cos
	} else {
		panic("unreachable")
	}
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

// InsertCsge adds a csge instruction at the end of b, comparing integral
// values src1 and src2 of srcType type and storing the result to dest with
// destType type.
func (b *Block) InsertCsge(dest Temporary, destType IntegralType,
	src1, src2 Value, srcType IntegralType) {
	var op opcode
	if srcType == LongType() {
		op = csgel
	} else if srcType == WordType() {
		op = csgew
	} else {
		panic("unreachable")
	}
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

// InsertCsgt adds a csgt instruction at the end of b, comparing integral
// values src1 and src2 of srcType type and storing the result to dest with
// destType type.
func (b *Block) InsertCsgt(dest Temporary, destType IntegralType,
	src1, src2 Value, srcType IntegralType) {
	var op opcode
	if srcType == LongType() {
		op = csgtl
	} else if srcType == WordType() {
		op = csgtw
	} else {
		panic("unreachable")
	}
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

// InsertCsle adds a csle instruction at the end of b, comparing integral
// values src1 and src2 of srcType type and storing the result to dest with
// destType type.
func (b *Block) InsertCsle(dest Temporary, destType IntegralType,
	src1, src2 Value, srcType IntegralType) {
	var op opcode
	if srcType == LongType() {
		op = cslel
	} else if srcType == WordType() {
		op = cslew
	} else {
		panic("unreachable")
	}
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

// InsertCslt adds a cslt instruction at the end of b, comparing integral
// values src1 and src2 of srcType type and storing the result to dest with
// destType type.
func (b *Block) InsertCslt(dest Temporary, destType IntegralType,
	src1, src2 Value, srcType IntegralType) {
	var op opcode
	if srcType == LongType() {
		op = csltl
	} else if srcType == WordType() {
		op = csltw
	} else {
		panic("unreachable")
	}
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

// InsertCuge adds a cuge instruction at the end of b, comparing integral
// values src1 and src2 of srcType type and storing the result to dest with
// destType type.
func (b *Block) InsertCuge(dest Temporary, destType IntegralType,
	src1, src2 Value, srcType IntegralType) {
	var op opcode
	if srcType == LongType() {
		op = cugel
	} else if srcType == WordType() {
		op = cugew
	} else {
		panic("unreachable")
	}
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

// InsertCugt adds a cugt instruction at the end of b, comparing integral
// values src1 and src2 of srcType type and storing the result to dest with
// destType type.
func (b *Block) InsertCugt(dest Temporary, destType IntegralType,
	src1, src2 Value, srcType IntegralType) {
	var op opcode
	if srcType == LongType() {
		op = cugtl
	} else if srcType == WordType() {
		op = cugtw
	} else {
		panic("unreachable")
	}
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

// InsertCule adds a cule instruction at the end of b, comparing integral
// values src1 and src2 of srcType type and storing the result to dest with
// destType type.
func (b *Block) InsertCule(dest Temporary, destType IntegralType,
	src1, src2 Value, srcType IntegralType) {
	var op opcode
	if srcType == LongType() {
		op = culel
	} else if srcType == WordType() {
		op = culew
	} else {
		panic("unreachable")
	}
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

// InsertCult adds a cult instruction at the end of b, comparing integral
// values src1 and src2 of srcType type and storing the result to dest with
// destType type.
func (b *Block) InsertCult(dest Temporary, destType IntegralType,
	src1, src2 Value, srcType IntegralType) {
	var op opcode
	if srcType == LongType() {
		op = cultl
	} else if srcType == WordType() {
		op = cultw
	} else {
		panic("unreachable")
	}
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}

// InsertCuo adds a cuo instruction at the end of b, comparing floating point
// values src1 and src2 of srcType type and storing the result to dest with
// destType type.
func (b *Block) InsertCuo(dest Temporary, destType IntegralType,
	src1, src2 Value, srcType FloatingPointType) {
	var op opcode
	if srcType == DoubleType() {
		op = cuod
	} else if srcType == SingleType() {
		op = cuos
	} else {
		panic("unreachable")
	}
	inst := newSimpleInst(op, dest, destType, src1, src2)
	b.insertInstruction(inst)
}
