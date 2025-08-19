package qbe

import "fmt"

// InsertStore adds a store instruction at the end of b, storing src to the
// address at dest. The size and type of the instruction is determined by type_.
func (b *Block) InsertStore(type_ ExtendedType, src, dest Value) {
	var op opcode
	switch type_ {
	case DoubleType():
		op = stored
	case SingleType():
		op = stores
	case LongType():
		op = storel
	case WordType():
		op = storew
	case HalfType():
		op = storeh
	case ByteType():
		op = storeb
	default:
		panic(fmt.Sprintf("unexpected type %#v", type_))
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
	switch srcType {
	case DoubleType():
		op = loadd
	case SingleType():
		op = loads
	case LongType():
		op = loadl
	case WordType():
		if signed {
			op = loadsw
		} else {
			op = loaduw
		}
	case HalfType():
		if signed {
			op = loadsh
		} else {
			op = loaduh
		}
	case ByteType():
		if signed {
			op = loadsb
		} else {
			op = loadub
		}
	default:
		panic(fmt.Sprintf("unexpected source type %#v", srcType))

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
