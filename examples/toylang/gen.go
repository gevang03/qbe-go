package main

// The code generator expects valid programs as input. Some simple bad cases are
// caught and handled with panic. Execution paths with missing returns run into
// the hlt instruction.

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gevang03/qbe-go/pkg/qbe"
)

type symtab struct {
	entries map[Ident]qbe.Value
	outer   *symtab
}

func newSymtab(outer *symtab) *symtab {
	return &symtab{make(map[Ident]qbe.Value), outer}
}

func (s *symtab) Lookup(name Ident) (qbe.Value, bool) {
	if value, ok := s.entries[name]; ok {
		return value, true
	}
	if s.outer == nil {
		return nil, false
	}
	return s.outer.Lookup(name)
}

func (s *symtab) Define(ident Ident, value qbe.Value) {
	s.entries[ident] = value
}

type Generator struct {
	mod             *qbe.Module
	currentFunction *qbe.Function
	currentBlock    *qbe.Block
	symtab          *symtab
}

func NewGenerator(name string) *Generator {
	return &Generator{
		mod:             qbe.NewModule(name),
		currentFunction: nil,
		currentBlock:    nil,
		symtab:          newSymtab(nil),
	}
}

func (gen *Generator) push() {
	gen.symtab = newSymtab(gen.symtab)
}

func (gen *Generator) pop() {
	gen.symtab = gen.symtab.outer
}

func (gen *Generator) Emit(program Program) error {
	for _, function := range program {
		gen.compileFunction(function)
	}
	w := bufio.NewWriter(os.Stdout)
	if _, err := gen.mod.ToIL(w); err != nil {
		return err
	}
	if err := w.Flush(); err != nil {
		return err
	}
	return nil
}

func (gen *Generator) compileFunction(function Function) {
	symbol := qbe.GlobalSymbol(function.Name)
	f := gen.mod.DefineFunction(symbol, qbe.WordType())
	gen.symtab.Define(function.Name, symbol)
	gen.push()
	gen.currentFunction = f
	gen.currentBlock = f.InsertBlock("start")
	for _, param := range function.Params {
		tmp := f.NewTemporary(string(param))
		gen.symtab.Define(param, tmp)
		f.InsertParam(qbe.WordType(), tmp)
	}
	gen.compileStmt(function.Body)
	gen.pop()
	if !gen.currentBlock.IsTerminated() {
		gen.currentBlock.InsertHlt()
	}
}

func (gen *Generator) compileBlock(block Block) {
	gen.push()
	for _, stmt := range block {
		gen.compileStmt(stmt)
	}
	gen.pop()
}

func (gen *Generator) compileStmt(stmt Stmt) {
	switch s := stmt.(type) {
	case Block:
		gen.compileBlock(s)
	case Assign:
		gen.compileAssign(s)
	case IfElse:
		gen.compileIfElse(s)
	case Return:
		gen.compileReturn(s)
	}
}

func (gen *Generator) compileAssign(assign Assign) {
	tmp, ok := gen.symtab.Lookup(assign.Dest)
	if !ok {
		tmp = gen.currentFunction.NewTemporary(string(assign.Dest))
		gen.symtab.Define(assign.Dest, tmp)
	}
	value := gen.compileExpr(assign.Value)
	if tmp, ok := tmp.(qbe.Temporary); ok {
		gen.currentBlock.InsertCopy(tmp, qbe.WordType(), value)
	} else {
		panic(fmt.Sprintf("cannot assign to function %v", assign.Dest))
	}
}

func (gen *Generator) compileIfElse(ifElse IfElse) {
	cond := gen.compileExpr(ifElse.Cond)
	condBlock := gen.currentBlock

	thenEntry := gen.currentFunction.InsertBlockAuto(".then")
	gen.currentBlock = thenEntry
	gen.compileStmt(ifElse.Then)
	thenEnd := gen.currentBlock

	elseEntry := gen.currentFunction.InsertBlockAuto(".else")
	gen.currentBlock = elseEntry
	gen.compileStmt(ifElse.Else)

	condBlock.InsertJnz(cond, thenEntry.Label(), elseEntry.Label())

	endIfBlock := gen.currentFunction.InsertBlockAuto(".endif")
	gen.currentBlock = endIfBlock

	if !thenEnd.IsTerminated() {
		thenEnd.InsertJmp(endIfBlock.Label())
	}
	// last block of else falls through
}

func (gen *Generator) compileReturn(s Return) {
	value := gen.compileExpr(s.Value)
	gen.currentBlock.InsertRet(value)
}

func (gen *Generator) compileExpr(expr Expr) qbe.Value {
	switch e := expr.(type) {
	case Bin:
		return gen.compileBin(e)
	case Ident:
		return gen.compileIdent(e)
	case Num:
		return qbe.Integer(e)
	case Call:
		return gen.compileCall(e)
	default:
		panic(fmt.Sprintf("unexpected main.Expr: %v", e))
	}
}

func (gen *Generator) compileIdent(ident Ident) qbe.Value {
	tmp, ok := gen.symtab.Lookup(ident)
	if !ok {
		panic(fmt.Sprintf("variable %v is not defined", string(ident)))
	}
	return tmp
}

func (gen *Generator) compileBin(bin Bin) qbe.Value {
	left := gen.compileExpr(bin.Left)
	right := gen.compileExpr(bin.Right)
	dest := gen.currentFunction.NewTemporary(".bin")
	switch bin.Op {
	case "+":
		gen.currentBlock.InsertAdd(dest, qbe.WordType(), left, right)
	case "-":
		gen.currentBlock.InsertSub(dest, qbe.WordType(), left, right)
	case "*":
		gen.currentBlock.InsertMul(dest, qbe.WordType(), left, right)
	case "/":
		gen.currentBlock.InsertDiv(dest, qbe.WordType(), left, right)
	case "==":
		gen.currentBlock.InsertCeq(dest, qbe.WordType(), left, right, qbe.WordType())
	case "!=":
		gen.currentBlock.InsertCne(dest, qbe.WordType(), left, right, qbe.WordType())
	case "<":
		gen.currentBlock.InsertCslt(dest, qbe.WordType(), left, right, qbe.WordType())
	case "<=":
		gen.currentBlock.InsertCsle(dest, qbe.WordType(), left, right, qbe.WordType())
	case ">":
		gen.currentBlock.InsertCsgt(dest, qbe.WordType(), left, right, qbe.WordType())
	case ">=":
		gen.currentBlock.InsertCsge(dest, qbe.WordType(), left, right, qbe.WordType())
	default:
		panic(fmt.Sprintf("unexpected main.Bin.Op: %#v", bin.Op))
	}
	return dest
}

func (gen *Generator) compileCall(call Call) qbe.Value {
	callee := gen.compileExpr(call.Callee)
	args := make([]qbe.Value, 0, len(call.Args))
	for _, arg := range call.Args {
		args = append(args, gen.compileExpr(arg))
	}
	callInst := gen.currentBlock.InsertCall(callee)
	dest := gen.currentFunction.NewTemporary(".call")
	callInst.SetDest(dest, qbe.WordType())
	for _, arg := range args {
		callInst.InsertArg(qbe.WordType(), arg)
	}
	return dest
}
