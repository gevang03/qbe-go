package qbe

// InsertOr adds an or instruction at the end of b, storing the bitwise or of
// src1 and src2 to dest with type type_.
func (b *Block) InsertOr(dest Temporary, type_ IntegralType, src1, src2 Value) {
	inst := newSimpleInst(or, dest, type_, src1, src2)
	b.insertInstruction(inst)
}

// InsertXor adds an xor instruction at the end of b, storing the bitwise xor of
// src1 and src2 to dest with type type_.
func (b *Block) InsertXor(dest Temporary, type_ IntegralType, src1, src2 Value) {
	inst := newSimpleInst(xor, dest, type_, src1, src2)
	b.insertInstruction(inst)
}

// InsertAnd adds an and instruction at the end of b, storing the bitwise and of
// src1 and src2 to dest with type type_.
func (b *Block) InsertAnd(dest Temporary, type_ IntegralType, src1, src2 Value) {
	inst := newSimpleInst(and, dest, type_, src1, src2)
	b.insertInstruction(inst)
}

// InsertSar adds a sar instruction at the end of b, storing the value of src1 shifted by src2 bits
// to the right and sign extended to dest with type type_.
func (b *Block) InsertSar(dest Temporary, type_ IntegralType, src1, src2 Value) {
	inst := newSimpleInst(sar, dest, type_, src1, src2)
	b.insertInstruction(inst)
}

// InsertShr adds a shr instruction at the end of b, storing the value of src1 shifted by src2 bits
// to the right and zero extended to dest with type type_.
func (b *Block) InsertShr(dest Temporary, type_ IntegralType, src1, src2 Value) {
	inst := newSimpleInst(shr, dest, type_, src1, src2)
	b.insertInstruction(inst)
}

// InsertShl adds a shl instruction at the end of b, storing the value of src1 shifted by src2 bits
// to the left to dest with type type_.
func (b *Block) InsertShl(dest Temporary, type_ IntegralType, src1, src2 Value) {
	inst := newSimpleInst(shl, dest, type_, src1, src2)
	b.insertInstruction(inst)
}
