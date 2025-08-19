package qbe

import "fmt"

// InsertFtoi adds an ftoi instruction at the end of b, converting src value
// with srcType type, to dest with type destType and signed based on signed.
func (b *Block) InsertFtoi(dest Temporary, destType IntegralType,
	src Value, srcType FloatingPointType, signed bool) {
	var op opcode
	switch srcType {
	case DoubleType():
		if signed {
			op = dtosi
		} else {
			op = dtoui
		}

	case SingleType():
		if signed {
			op = stosi
		} else {
			op = stoui
		}
	default:
		panic(fmt.Sprintf("unexpected source type %#v", srcType))
	}
	inst := newSimpleInst(op, dest, destType, src)
	b.insertInstruction(inst)
}

// InsertItof adds an itof instruction at the end of b, converting src value
// with srcType type and signed based on signed, to dest with type destType.
func (b *Block) InsertItof(dest Temporary, destType FloatingPointType,
	src Value, srcType IntegralType, signed bool) {
	var op opcode
	switch srcType {
	case LongType():
		if signed {
			op = sltof
		} else {
			op = ultof
		}
	case WordType():
		if signed {
			op = swtof
		} else {
			op = uwtof
		}
	default:
		panic(fmt.Sprintf("unexpected source type %#v", srcType))
	}
	inst := newSimpleInst(op, dest, destType, src)
	b.insertInstruction(inst)
}

// InsertExt adds an ext instruction at the end of b, storing the zero or
// signed (based on signed) extended value of src with srcType type to dest
// of destType. Valid srcType values are [WordType], [HalfType] and [ByteType].
func (b *Block) InsertExt(dest Temporary, destType IntegralType,
	src Value, srcType ExtendedType, signed bool) {
	var op opcode
	switch srcType {
	case WordType():
		if signed {
			op = extsw
		} else {
			op = extuw
		}
	case HalfType():
		if signed {
			op = extsh
		} else {
			op = extuh
		}
	case ByteType():
		if signed {
			op = extsb
		} else {
			op = extub
		}
	default:
		panic(fmt.Sprintf("unexpected source type %#v", srcType))
	}
	inst := newSimpleInst(op, dest, destType, src)
	b.insertInstruction(inst)
}

// InsertExts adds an exts instruction at the end of b, storing the value of
// src as [DoubleType] to dest.
func (b *Block) InsertExts(dest Temporary, src Value) {
	inst := newSimpleInst(exts, dest, DoubleType(), src)
	b.insertInstruction(inst)
}

// InsertTruncd adds a truncd instruction at the end of b, storing the value of
// src as [SingleType] to dest.
func (b *Block) InsertTruncd(dest Temporary, src Value) {
	inst := newSimpleInst(truncd, dest, SingleType(), src)
	b.insertInstruction(inst)
}

// InsertCast adds a cast instruction at the end of b, casting value src
// to type type_ and storing at dest.
func (b *Block) InsertCast(dest Temporary, type_ BaseType, src Value) {
	inst := newSimpleInst(cast, dest, type_, src)
	b.insertInstruction(inst)
}

// InsertCopy adds a copy instruction at the end of b, copying value src
// to type type_ and storing at dest.
func (b *Block) InsertCopy(dest Temporary, type_ BaseType, src Value) {
	inst := newSimpleInst(copy, dest, type_, src)
	b.insertInstruction(inst)
}
