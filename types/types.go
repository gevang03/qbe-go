package types

type (
	IntegralType interface {
		BaseType
		isIntegralType()
	}

	MemoryType interface {
		IntegralType
		isMemoryType()
	}

	FloatingPointType interface {
		BaseType
		isFloatingPointType()
	}

	BaseType interface {
		ExtendedType
		isBaseType()
	}

	ExtendedType interface {
		ABIType
		SubType
		isExtendedType()
	}

	SubType interface {
		ABIType
		isSubType()
	}

	SubWordType interface{ isSubWordType() }

	ABIType interface{ isABIType() }
)
