package qbe

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

var pointer MemoryType = longType

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

// SetPointerType sets the type of pointer to pointerType. pointerType must be either word or long.
func SetPointerType(pointerType IntegralType) {
	switch pointerType {
	case wordType:
		pointer = wordType
	case longType:
		pointer = longType
	default:
		panic("Pointer Type must be either word or long")
	}
}

// PointerType returns the type of PointerType used. Equivalent to one of [WordType] or [LongType].
func PointerType() MemoryType {
	return pointer
}

func (primitiveType) isMemoryType()        {}
func (primitiveType) isIntegralType()      {}
func (primitiveType) isFloatingPointType() {}
func (primitiveType) isBaseType()          {}
func (primitiveType) isExtendedType()      {}
func (primitiveType) isSubType()           {}
func (primitiveType) isABIType()           {}
func (primitiveType) isSubWordType()       {}

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
