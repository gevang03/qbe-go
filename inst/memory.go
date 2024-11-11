package inst

import (
	"github.com/gevang03/qbe-go/types"
	"github.com/gevang03/qbe-go/value"
)

func (b *Block) InsertStore(type_ types.ExtendedType, src, dest value.Value) {
	var op opcode
	if type_ == types.Double() {
		op = stored
	} else if type_ == types.Single() {
		op = stores
	} else if type_ == types.Long() {
		op = storel
	} else if type_ == types.Word() {
		op = storew
	} else if type_ == types.Half() {
		op = storeh
	} else if type_ == types.Byte() {
		op = storeb
	} else {
		panic("unreachable")
	}
	inst := newSimpleInstNoDest(op, src, dest)
	b.insertInstruction(inst)
}

func (b *Block) InsertLoad(dest value.Temporary, destType types.BaseType,
	src value.Value, srcType types.ExtendedType, signed bool) {
	var op opcode
	if srcType == types.Double() {
		op = loadd
	} else if srcType == types.Single() {
		op = loads
	} else if srcType == types.Long() {
		op = loadl
	} else if srcType == types.Word() {
		if signed {
			op = loadsw
		} else {
			op = loaduw
		}
	} else if srcType == types.Half() {
		if signed {
			op = loadsh
		} else {
			op = loaduh
		}
	} else if srcType == types.Byte() {
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

func (b *Block) InsertBlit(src, dest value.Value, count uint) {
	inst := newSimpleInstNoDest(blit, src, dest, value.Integer(count))
	b.insertInstruction(inst)
}

func (b *Block) InsertAlloc(dest value.Temporary, align uint, count value.Value) {
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
	inst := newSimpleInst(op, dest, types.Pointer(), count)
	b.insertInstruction(inst)
}