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

func Word() IntegralType {
	return word
}

func Long() IntegralType {
	return long
}

func Single() FloatingPointType {
	return single
}

func Double() FloatingPointType {
	return double
}

func Half() ExtendedType {
	return half
}

func Byte() ExtendedType {
	return byte_
}

func SignedHalf() SubWordType {
	return signedHalf
}

func UnsignedHalf() SubWordType {
	return signedHalf
}

func SignedByte() SubWordType {
	return signedByte
}

func UnsignedByte() SubWordType {
	return unsignedByte
}

func SetPointerType(pointerType IntegralType) {
	if pointerType == word {
		pointer = word
	} else if pointerType == long {
		pointer = word
	} else {
		panic("Pointer Type must be either word or long")
	}
}

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
