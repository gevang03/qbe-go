package qbe

import "fmt"

type primitiveType int

const (
	wordType primitiveType = iota
	longType
	singleType
	doubleType

	halfType
	byteType

	signedByteType
	unsignedByteType
	signedHalfType
	unsignedHalfType
)

// WordType returns the type for 32-bit integer.
func WordType() IntegralType {
	return wordType
}

// LongType returns the type for 64-bit integer.
func LongType() IntegralType {
	return longType
}

// SingleType returns the type for 32-bit floating point.
func SingleType() FloatingPointType {
	return singleType
}

// DoubleType returns the type for 64-bit floating point.
func DoubleType() FloatingPointType {
	return doubleType
}

// HalfType returns the type for 16-bit integer.
func HalfType() ExtendedType {
	return halfType
}

// ByteType returns the type for 8-bit integer.
func ByteType() ExtendedType {
	return byteType
}

// SignedHalfType returns the type for signed 16-bit integer.
func SignedHalfType() SubWordType {
	return signedHalfType
}

// UnsignedHalfType returns the type for unsigned 16-bit integer.
func UnsignedHalfType() SubWordType {
	return signedHalfType
}

// SignedByteType returns the type for signed 8-bit integer.
func SignedByteType() SubWordType {
	return signedByteType
}

// UnsignedByteType returns the type unsigned 8-bit integer.
func UnsignedByteType() SubWordType {
	return unsignedByteType
}

// PointerType returns the type for pointers. Equivalent to [LongType] as
// QBE currently supports 64-bit platforms only.
func PointerType() MemoryType {
	return longType
}

func (primitiveType) isMemoryType()        {}
func (primitiveType) isIntegralType()      {}
func (primitiveType) isFloatingPointType() {}
func (primitiveType) isBaseType()          {}
func (primitiveType) isExtendedType()      {}
func (primitiveType) isSubType()           {}
func (primitiveType) isABIType()           {}
func (primitiveType) isSubWordType()       {}

func (p primitiveType) SizeOf() uint {
	switch p {
	case byteType:
		return 1
	case doubleType:
		return 8
	case halfType:
		return 2
	case longType:
		return 8
	case signedByteType:
		return 1
	case signedHalfType:
		return 2
	case singleType:
		return 4
	case unsignedByteType:
		return 1
	case unsignedHalfType:
		return 2
	case wordType:
		return 4
	default:
		panic(fmt.Sprintf("unexpected qbe.primitiveType: %#v", p))
	}
}

func (p primitiveType) AlignOf() uint {
	return p.SizeOf()
}

func (p primitiveType) String() string {
	return [...]string{
		wordType:         "w",
		longType:         "l",
		singleType:       "s",
		doubleType:       "d",
		halfType:         "h",
		byteType:         "b",
		signedByteType:   "sb",
		unsignedByteType: "ub",
		signedHalfType:   "sh",
		unsignedHalfType: "uh",
	}[p]
}

func (p primitiveType) Name() string {
	return p.String()
}
