package main

func main() {
	program := Program{
		{
			Name:   "fma",
			Params: []Ident{"x", "y", "z"},
			Body: Return{
				Bin{
					Op: "+",
					Left: Bin{
						Op:    "*",
						Left:  Ident("x"),
						Right: Ident("y"),
					},
					Right: Ident("z")},
			},
		},

		{
			Name:   "min",
			Params: []Ident{"a", "b"},
			Body: Block{
				Assign{Dest: "m", Value: Num(0)},
				IfElse{
					Cond: Bin{"<", Ident("a"), Ident("b")},
					Then: Assign{Dest: "m", Value: Ident("a")},
					Else: Assign{Dest: "m", Value: Ident("b")},
				},
				Return{Ident("m")},
			},
		},

		{
			Name:   "min3",
			Params: []Ident{"x", "y", "z"},
			Body: Return{
				Call{
					Callee: Ident("min"),
					Args: []Expr{
						Ident("x"),
						Call{
							Callee: Ident("min"),
							Args:   []Expr{Ident("y"), Ident("z")},
						},
					},
				},
			},
		},

		{
			Name:   "cmp",
			Params: []Ident{"x", "y"},
			Body: IfElse{
				Cond: Bin{"<", Ident("x"), Ident("y")},
				Then: Return{Num(-1)},
				Else: IfElse{
					Cond: Bin{">", Ident("x"), Ident("y")},
					Then: Return{Num(1)},
					Else: Return{Num(0)},
				},
			},
		},

		{
			Name:   "fct",
			Params: []Ident{"n"},
			Body: IfElse{
				Cond: Bin{"<", Ident("n"), Num(1)},
				Then: Return{Num(1)},
				Else: Return{
					Bin{
						Op:   "*",
						Left: Ident("n"),
						Right: Call{
							Callee: Ident("fct"),
							Args:   []Expr{Bin{"-", Ident("n"), Num(1)}},
						},
					},
				},
			},
		},
	}

	gen := NewGenerator("basic")
	if err := gen.Emit(program); err != nil {
		panic(err)
	}
}
