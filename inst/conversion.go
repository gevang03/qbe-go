package inst

import (
	"github.com/gevang03/qbe-go/types"
	"github.com/gevang03/qbe-go/value"
)

// InsertFtoi adds an ftoi instruction at the end of b, converting src value
// with srcType type, to dest with type destType and signed based on signed.
func (b *Block) InsertFtoi(dest value.Temporary, destType types.IntegralType,
	src value.Value, srcType types.FloatingPointType, signed bool) {
	var op opcode
	if srcType == types.Double() {
		if signed {
			op = dtosi
		} else {
			op = dtoui
		}
	} else if srcType == types.Single() {
		if signed {
			op = stosi
		} else {
			op = stoui
		}
	} else {
		panic("unreachable")
	}
	inst := newSimpleInst(op, dest, destType, src)
	b.insertInstruction(inst)
}

// InsertItof adds an itof instruction at the end of b, converting src value
// with srcType type and signed based on signed, to dest with type destType.
func (b *Block) InsertItof(dest value.Temporary, destType types.FloatingPointType,
	src value.Value, srcType types.IntegralType, signed bool) {
	var op opcode
	if srcType == types.Long() {
		if signed {
			op = sltof
		} else {
			op = ultof
		}
	} else if srcType == types.Word() {
		if signed {
			op = swtof
		} else {
			op = uwtof
		}
	} else {
		panic("unreachable")
	}
	inst := newSimpleInst(op, dest, destType, src)
	b.insertInstruction(inst)
}

// InsertExt adds an ext instruction at the end of b, storing the zero or
// signed (based on signed) extended value of src with srcType type to dest
// of destType. Valid srcType values are [types.Word](), [types.Half]()
// and [types.Byte]().
func (b *Block) InsertExt(dest value.Temporary, destType types.IntegralType,
	src value.Value, srcType types.ExtendedType, signed bool) {
	var op opcode
	if srcType == types.Word() {
		if signed {
			op = extsw
		} else {
			op = extuw
		}
	} else if srcType == types.Half() {
		if signed {
			op = extsh
		} else {
			op = extuh
		}
	} else if srcType == types.Byte() {
		if signed {
			op = extsb
		} else {
			op = extub
		}
	} else {
		panic("invalid source type")
	}
	inst := newSimpleInst(op, dest, destType, src)
	b.insertInstruction(inst)
}

// InsertExts adds an exts instruction at the end of b, storing the value of
// src as [types.Double]() to dest.
func (b *Block) InsertExts(dest value.Temporary, src value.Value) {
	inst := newSimpleInst(exts, dest, types.Double(), src)
	b.insertInstruction(inst)
}

// InsertTruncd adds a truncd instruction at the end of b, storing the value of
// src as [types.Single]() to dest.
func (b *Block) InsertTruncd(dest value.Temporary, src value.Value) {
	inst := newSimpleInst(truncd, dest, types.Single(), src)
	b.insertInstruction(inst)
}

// InsertCast adds a cast instruction at the end of b, casting value src
// to type type_ and storing at dest.
func (b *Block) InsertCast(dest value.Temporary, type_ types.BaseType, src value.Value) {
	inst := newSimpleInst(cast, dest, type_, src)
	b.insertInstruction(inst)
}

// InsertCopy adds a copy instruction at the end of b, copying value src
// to type type_ and storing at dest.
func (b *Block) InsertCopy(dest value.Temporary, type_ types.BaseType, src value.Value) {
	inst := newSimpleInst(copy, dest, type_, src)
	b.insertInstruction(inst)
}
