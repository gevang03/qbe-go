package qbe

// InsertStore adds a store instruction at the end of b, storing src to the
// address at dest. The size and type of the instruction is determined by type_.
func (b *Block) InsertStore(type_ ExtendedType, src, dest Value) {
	var op opcode
	if type_ == DoubleType() {
		op = stored
	} else if type_ == SingleType() {
		op = stores
	} else if type_ == LongType() {
		op = storel
	} else if type_ == WordType() {
		op = storew
	} else if type_ == HalfType() {
		op = storeh
	} else if type_ == ByteType() {
		op = storeb
	} else {
		panic("unreachable")
	}
	inst := newSimpleInstNoDest(op, src, dest)
	b.insertInstruction(inst)
}

// InsertLoad adds a load instruction at the end of b. The value at src with
// srcType is loaded at dest with destType. Integer values may be zero or sign
// extended based on the value of signed, whereas signed is irrelevant with
// floating point source
func (b *Block) InsertLoad(dest Temporary, destType BaseType,
	src Value, srcType ExtendedType, signed bool) {
	var op opcode
	if srcType == DoubleType() {
		op = loadd
	} else if srcType == SingleType() {
		op = loads
	} else if srcType == LongType() {
		op = loadl
	} else if srcType == WordType() {
		if signed {
			op = loadsw
		} else {
			op = loaduw
		}
	} else if srcType == HalfType() {
		if signed {
			op = loadsh
		} else {
			op = loaduh
		}
	} else if srcType == ByteType() {
		if signed {
			op = loadsb
		} else {
			op = loadub
		}
	} else {
		panic("unreachable")
	}
	inst := newSimpleInst(op, dest, destType, src)
	b.insertInstruction(inst)
}

// InsertBlit adds a blit instruction at the end of b, where count bytes are
// copied from dest to src.
func (b *Block) InsertBlit(src, dest Value, count uint) {
	inst := newSimpleInstNoDest(blit, src, dest, Integer(count))
	b.insertInstruction(inst)
}

// InsertAlloc adds a allocN instruction at the end of b with valid align values
// 4, 8 or 16, of count bytes on the stack.
func (b *Block) InsertAlloc(dest Temporary, align uint, count Value) {
	var op opcode
	switch align {
	case 4:
		op = alloc4
	case 8:
		op = alloc8
	case 16:
		op = alloc16
	default:
		panic("Invalid alloc alignment")
	}
	inst := newSimpleInst(op, dest, PointerType(), count)
	b.insertInstruction(inst)
}
