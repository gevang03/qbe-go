package inst

import (
	"fmt"
	"strings"
)

type Label string

func (l Label) String() string { return "@" + string(l) }

type Block struct {
	label Label
	phis  []Phi
	insts []inst
	jump
}

func NewBlock(label Label) *Block {
	return &Block{
		label: label,
		phis:  nil,
		insts: nil,
		jump:  nil,
	}
}

func (b *Block) Label() Label {
	return b.label
}

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
