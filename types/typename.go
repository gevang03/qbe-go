package types

type TypeName string

func (TypeName) isSubType() {}
func (TypeName) isABIType() {}

func (tn TypeName) String() string { return ":" + string(tn) }
