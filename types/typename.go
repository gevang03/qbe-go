package types

import (
	"fmt"

	"github.com/gevang03/qbe-go/internal/validation"
)

// A TypeName represents the name given to an aggregate definition [def.Struct], [def.Union] or [def.Opaque].
type TypeName string

func (TypeName) isSubType() {}
func (TypeName) isABIType() {}

// String converts tn to a string compatible with QBE code.
func (tn TypeName) String() string {
	s := string(tn)
	if !validation.Validate(s) {
		panic(fmt.Sprintf("Type name :%q is not valid", s))
	}
	return ":" + s
}
