package value

import "fmt"

type (
	Value interface{ isValue() }

	DataItem interface{ isDataItem() }
)

type (
	Integer      int64
	Single       float32
	Double       float64
	GlobalSymbol string

	ThreadLocalSymbol string

	Temporary string

	DataString string

	Address struct {
		GlobalSymbol
		Offset int64
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

func (i Integer) String() string             { return fmt.Sprint(int64(i)) }
func (s Single) String() string              { return fmt.Sprintf("s_%v", float32(s)) }
func (d Double) String() string              { return fmt.Sprintf("d_%v", float64(d)) }
func (sym GlobalSymbol) String() string      { return "$" + string(sym) }
func (tls ThreadLocalSymbol) String() string { return "thread $" + string(tls) }
func (str DataString) String() string        { return fmt.Sprintf("%q", string(str)) }
func (addr Address) String() string          { return fmt.Sprintf("%v + %v", addr.GlobalSymbol, addr.Offset) }
func (tmp Temporary) String() string         { return "%" + string(tmp) }
