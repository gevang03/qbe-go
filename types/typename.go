package types

// A TypeName represents the name given to an aggregate definition [def.Struct], [def.Union] or [def.Opaque].
type TypeName string

func (TypeName) isSubType() {}
func (TypeName) isABIType() {}

// String converts tn to a string compatible with QBE code.
func (tn TypeName) String() string { return ":" + string(tn) }
