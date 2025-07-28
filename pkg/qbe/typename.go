package qbe

import "fmt"

// A TypeName represents the name given to an aggregate definition [Struct], [Union] or [Opaque].
type TypeName string

func (TypeName) isSubType() {}
func (TypeName) isABIType() {}

// String converts tn to a string compatible with QBE code.
func (tn TypeName) String() string {
	s := string(tn)
	if !validateName(s) {
		panic(fmt.Sprintf("Type name :%q is not valid", s))
	}
	return ":" + s
}
