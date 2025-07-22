package inst

import (
	"fmt"
	"strings"

	"github.com/gevang03/qbe-go/internal/validation"
)

// A Label represents the name of the entrypoint to a [Block].
type Label string

// String converts l to a string compatible with QBE code.
func (l Label) String() string {
	s := string(l)
	if !validation.Validate(s) {
		panic(fmt.Sprintf("Label @%q is not valid", s))
	}
	return "@" + s
}

// A Block represents a QBE basic block.
type Block struct {
	label Label
	phis  []phi
	insts []inst
	jump
}

// NewBlock creates a new [Block] with label at its entrypoint,
// no phi nodes or regular instructions and no terminating instruction, i.e
// control flow falls through to the next basic block.
func NewBlock(label Label) *Block {
	return &Block{
		label: label,
		phis:  nil,
		insts: nil,
		jump:  nil,
	}
}

// Label returns the label of b.
func (b *Block) Label() Label {
	return b.label
}

// String converts b to a string compatible with QBE code.
func (b *Block) String() string {
	var builder = strings.Builder{}
	builder.WriteString(b.label.String())
	builder.WriteByte('\n')
	for _, phi := range b.phis {
		builder.WriteByte('\t')
		builder.WriteString(phi.String())
		builder.WriteByte('\n')
	}
	for _, inst := range b.insts {
		builder.WriteByte('\t')
		builder.WriteString(fmt.Sprint(inst))
		builder.WriteByte('\n')
	}
	if b.jump != nil {
		builder.WriteByte('\t')
		builder.WriteString(fmt.Sprint(b.jump))
		builder.WriteByte('\n')
	}
	return builder.String()
}
