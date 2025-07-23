package types

type primitiveType int

const (
	word primitiveType = iota
	long
	single
	double

	half
	byte_

	signedByte
	unsignedByte
	signedHalf
	unsignedHalf
)

var pointer MemoryType = long

// Word returns the type for 32-bit integer.
func Word() IntegralType {
	return word
}

// Long returns the type for 64-bit integer.
func Long() IntegralType {
	return long
}

// Single returns the type for 32-bit floating point.
func Single() FloatingPointType {
	return single
}

// Double returns the type for 64-bit floating point.
func Double() FloatingPointType {
	return double
}

// Half returns the type for 16-bit integer.
func Half() ExtendedType {
	return half
}

// Byte returns the type for 8-bit integer.
func Byte() ExtendedType {
	return byte_
}

// SignedHalf returns the type for signed 16-bit integer.
func SignedHalf() SubWordType {
	return signedHalf
}

// UnsignedHalf returns the type for unsigned 16-bit integer.
func UnsignedHalf() SubWordType {
	return signedHalf
}

// SignedByte returns the type for signed 8-bit integer.
func SignedByte() SubWordType {
	return signedByte
}

// UnsignedByte returns the type unsigned 8-bit integer.
func UnsignedByte() SubWordType {
	return unsignedByte
}

// SetPointerType sets the type of pointer to pointerType. pointerType must be either word or long.
func SetPointerType(pointerType IntegralType) {
	switch pointerType {
	case word:
		pointer = word
	case long:
		pointer = long
	default:
		panic("Pointer Type must be either word or long")
	}
}

// Pointer returns the type of Pointer used. Equivalent to one of [Word] or [Long].
func Pointer() MemoryType {
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
		word:         "w",
		long:         "l",
		single:       "s",
		double:       "d",
		half:         "h",
		byte_:        "b",
		signedByte:   "sb",
		unsignedByte: "ub",
		signedHalf:   "sh",
		unsignedHalf: "uh",
	}[p]
}
