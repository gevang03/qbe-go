package main

// This toy language consists of functions with simple imperative statements.
// All values are 32-bit untyped integers. No front end is implemented, ASTs
// can be created by hand.

type Program []Function

type Function struct {
	Name   Ident
	Params []Ident
	Body   Stmt
}

// We'll just have to make do with poor mans sum types.

type (
	Stmt interface{ isStmt() }
	Expr interface{ isExpr() }
)

// Statements

type (
	Block []Stmt

	Assign struct {
		Dest  Ident
		Value Expr
	}

	IfElse struct {
		Cond Expr
		Then Stmt
		Else Stmt
	}

	Return struct{ Value Expr }
)

// Expressions

type (
	Ident string

	Num int32

	Bin struct {
		Op    string
		Left  Expr
		Right Expr
	}

	Call struct {
		Callee Expr
		Args   []Expr
	}
)

func (Block) isStmt()  {}
func (Assign) isStmt() {}
func (IfElse) isStmt() {}
func (Return) isStmt() {}

func (Ident) isExpr() {}
func (Num) isExpr()   {}
func (Bin) isExpr()   {}
func (Call) isExpr()  {}
