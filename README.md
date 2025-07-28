# QBE Bindings for Go

This library aims to simplify generating [QBE](https://c9x.me/compile/) v1.2 IL files in go.

## Basic Usage

```go
package main

import (
	"bufio"
	"os"

	"github.com/gevang03/qbe-go/pkg/qbe"
)

func main() {
	// Create module
	mod := qbe.NewModule("hello")

	// Define string data
	cstr := mod.DefineData("message")
	cstr.SectionName = ".rodata"
	cstr.InsertValue(qbe.ByteType(), qbe.DataString("Hello, world!"), qbe.Integer(0))

	// Define main function
	mainFn := mod.DefineFunction("main")
	mainFn.Export = true
	mainFn.RetType = qbe.WordType()

	block := mainFn.InsertBlockAuto()

	// Call puts function
	putsCall := block.InsertCall(qbe.GlobalSymbol("puts"))
	putsCall.InsertArg(qbe.PointerType(), cstr.Name)

	// Return 0 from main
	block.InsertRet(qbe.Integer(0))

	// Write IL to stdout
	w := bufio.NewWriter(os.Stdout)
	if _, err := mod.ToIL(w); err != nil {
		panic(err)
	}
	w.Flush()
}
```