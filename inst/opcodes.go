package inst

type opcode int

const (
	add opcode = iota
	sub
	div
	mul
	udiv
	rem
	urem

	neg

	or
	xor
	and
	sar
	shr
	shl

	stored
	stores
	storel
	storew
	storeh
	storeb

	loadd
	loads
	loadl
	loadsw
	loaduw
	loadsh
	loaduh
	loadsb
	loadub

	blit

	alloc4
	alloc8
	alloc16

	ceqd
	ceql
	ceqs
	ceqw
	cged
	cges
	cgtd
	cgts
	cled
	cles
	cltd
	clts
	cned
	cnel
	cnes
	cnew
	cod
	cos
	csgel
	csgew
	csgtl
	csgtw
	cslel
	cslew
	csltl
	csltw
	cugel
	cugew
	cugtl
	cugtw
	culel
	culew
	cultl
	cultw
	cuod
	cuos

	extsw
	extuw
	extsh
	extuh
	extsb
	extub
	exts
	truncd
	stosi
	stoui
	dtosi
	dtoui
	swtof
	uwtof
	sltof
	ultof

	cast
	copy

	// call

	vastart
	vaarg

	// hlt
	// jmp
	// jnz
	// ret
)

func (op opcode) String() string {
	return [...]string{
		add:     "add",
		sub:     "sub",
		div:     "div",
		mul:     "mul",
		udiv:    "udiv",
		rem:     "rem",
		urem:    "urem",
		neg:     "neg",
		or:      "or",
		xor:     "xor",
		and:     "and",
		sar:     "sar",
		shr:     "shr",
		shl:     "shl",
		stored:  "stored",
		stores:  "stores",
		storel:  "storel",
		storew:  "storew",
		storeh:  "storeh",
		storeb:  "storeb",
		loadd:   "loadd",
		loads:   "loads",
		loadl:   "loadl",
		loadsw:  "loadsw",
		loaduw:  "loaduw",
		loadsh:  "loadsh",
		loaduh:  "loaduh",
		loadsb:  "loadsb",
		loadub:  "loadub",
		blit:    "blit",
		alloc4:  "alloc4",
		alloc8:  "alloc8",
		alloc16: "alloc16",
		ceqd:    "ceqd",
		ceql:    "ceql",
		ceqs:    "ceqs",
		ceqw:    "ceqw",
		cged:    "cged",
		cges:    "cges",
		cgtd:    "cgtd",
		cgts:    "cgts",
		cled:    "cled",
		cles:    "cles",
		cltd:    "cltd",
		clts:    "clts",
		cned:    "cned",
		cnel:    "cnel",
		cnes:    "cnes",
		cnew:    "cnew",
		cod:     "cod",
		cos:     "cos",
		csgel:   "csgel",
		csgew:   "csgew",
		csgtl:   "csgtl",
		csgtw:   "csgtw",
		cslel:   "cslel",
		cslew:   "cslew",
		csltl:   "csltl",
		csltw:   "csltw",
		cugel:   "cugel",
		cugew:   "cugew",
		cugtl:   "cugtl",
		cugtw:   "cugtw",
		culel:   "culel",
		culew:   "culew",
		cultl:   "cultl",
		cultw:   "cultw",
		cuod:    "cuod",
		cuos:    "cuos",
		extsw:   "extsw",
		extuw:   "extuw",
		extsh:   "extsh",
		extuh:   "extuh",
		extsb:   "extsb",
		extub:   "extub",
		exts:    "exts",
		truncd:  "truncd",
		stosi:   "stosi",
		stoui:   "stoui",
		dtosi:   "dtosi",
		dtoui:   "dtoui",
		swtof:   "swtof",
		uwtof:   "uwtof",
		sltof:   "sltof",
		ultof:   "ultof",
		cast:    "cast",
		copy:    "copy",
		vastart: "vastart",
		vaarg:   "vaarg",
	}[op]
}
