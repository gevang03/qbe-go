// Package value exports types used as [def.Data] members
// and as instruction operands in the [inst] package.
package value

import (
	"fmt"

	"github.com/gevang03/qbe-go/internal/validation"
)

type (
	// A Value is one of the following:
	// [Integer], [Single], [Double], [GlobalSymbol], [ThreadLocalSymbol], [Temporary].
	Value interface{ isValue() }

	// A DataItem is any [Value] or one of the following:
	// [DataString], [Address]
	DataItem interface{ isDataItem() }
)

type (
	// An Integer represents any integer value (signed/unsigned, 8/16/32/64 bits) used in QBE.
	Integer int64
	// A Single represents single precision floating point values in QBE.
	Single float32
	// A Double represents double precision floating point values in QBE.
	Double float64
	// A GlobalSymbol represents symbols used to reference functions or data in QBE.
	GlobalSymbol string
	// A ThreadLocalSymbol represents symbols with thread local storage in QBE.
	ThreadLocalSymbol string
	// A Temporary represents temporary values used in QBE.
	Temporary string
	// A DataString is a string literal used as a member in [def.Data]
	DataString string
	// An Address can be used as a member in [def.Data].
	Address struct {
		GlobalSymbol       // Base of address used.
		Offset       int64 // Number of bytes after base symbol. If offset is zero a [GlobalSymbol] can be used directly.
	}
)

func (Integer) isValue()           {}
func (Single) isValue()            {}
func (Double) isValue()            {}
func (GlobalSymbol) isValue()      {}
func (ThreadLocalSymbol) isValue() {}
func (Temporary) isValue()         {}

func (Integer) isDataItem()      {}
func (Single) isDataItem()       {}
func (Double) isDataItem()       {}
func (GlobalSymbol) isDataItem() {}
func (DataString) isDataItem()   {}
func (Address) isDataItem()      {}

// String converts i to a string compatible with QBE code.
func (i Integer) String() string { return fmt.Sprint(int64(i)) }

// String converts s to a string compatible with QBE code.
func (s Single) String() string { return fmt.Sprintf("s_%v", float32(s)) }

// String converts d to a string compatible with QBE code.
func (d Double) String() string { return fmt.Sprintf("d_%v", float64(d)) }

// String converts sym to a string compatible with QBE code.
func (sym GlobalSymbol) String() string {
	s := string(sym)
	if !validation.Validate(s) {
		panic(fmt.Sprintf("Global symbol $%q is not valid", s))
	}
	return "$" + s
}

// String converts tls to a string compatible with QBE code.
func (tls ThreadLocalSymbol) String() string {
	s := string(tls)
	if !validation.Validate(s) {
		panic(fmt.Sprintf("Thread local symbol thread $%q is not valid", s))
	}
	return "thread $" + s
}

// String converts str to a string compatible with QBE code.
func (str DataString) String() string { return fmt.Sprintf("%q", string(str)) }

// String converts addr to a string compatible with QBE code.
func (addr Address) String() string { return fmt.Sprintf("%v + %v", addr.GlobalSymbol, addr.Offset) }

// String converts tmp to a string compatible with QBE code.
func (tmp Temporary) String() string {
	s := string(tmp)
	if !validation.Validate(s) {
		panic(fmt.Sprintf("Temporary %%%q is not valid", s))
	}
	return "%" + s
}
