// package types exports types and type constructors used in QBE definitions and instructions.
package types

type (
	// An IntegralType is any of the values returned by [Word] or [Long].
	IntegralType interface {
		BaseType
		isIntegralType()
	}

	// The MemoryType represents the size of pointers. It is always one of [Word] or [Long].
	MemoryType interface {
		IntegralType
		isMemoryType()
	}

	// A FloatingPointType is any of the values returned by [Single] or [Double].
	FloatingPointType interface {
		BaseType
		isFloatingPointType()
	}

	// A BaseType is any type of [IntegralType] or [FloatingPointType].
	BaseType interface {
		ExtendedType
		isBaseType()
	}

	// An Extended type is any [BaseType] or the values returned by [Half] and [Byte].
	ExtendedType interface {
		ABIType
		SubType
		isExtendedType()
	}

	// A SubType is any [ExtendedType] or a [TypeName].
	SubType interface {
		ABIType
		isSubType()
	}

	// A SubWordType is any signed or unsigned integer type with width less or equal to 16-bits,
	// namely values returned by [SignedHalf], [UnsignedHalf], [SignedByte], [UnsignedByte].
	SubWordType interface{ isSubWordType() }

	// An ABIType is any [BaseType], [SubWordType] or [TypeName].
	ABIType interface{ isABIType() }
)
