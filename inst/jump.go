package inst

import (
	"fmt"

	"github.com/gevang03/qbe-go/value"
)

type jump interface{ isJump() }

type (
	jmp struct{ Target Label }

	jnz struct {
		value.Value
		NotZero Label
		Zero    Label
	}

	ret struct{ value.Value }

	hlt struct{}
)

func (jmp) isJump() {}
func (jnz) isJump() {}
func (ret) isJump() {}
func (hlt) isJump() {}

func (j jmp) String() string { return "jmp " + j.Target.String() }

func (j jnz) String() string {
	return fmt.Sprintf("jnz %v, %v, %v", j.Value, j.NotZero, j.Zero)
}

func (r ret) String() string {
	if r.Value != nil {
		return fmt.Sprintf("ret %v", r.Value)
	}
	return "ret"
}

func (hlt) String() string { return "hlt" }

// InsertJmp terminates b with an unconditional jump to the block that begins at label.
func (b *Block) InsertJmp(label Label) {
	b.jump = jmp{label}
}

// InsertJnz terminates b with a conditional jump based on value.
// If value evaluates to non-zero, control flow directs to notZero,
// otherwise to zero.
func (b *Block) InsertJnz(value value.Value, notZero, zero Label) {
	b.jump = jnz{value, notZero, zero}
}

// InsertRet terminates b with a return instruction, returning value.
// If value is nil no value is set as the operand of the instruction.
func (b *Block) InsertRet(value value.Value) {
	b.jump = ret{value}
}

// InsertHlt terminates b with a halt instruction.
func (b *Block) InsertHlt() {
	b.jump = hlt{}
}

// IsTerminated returns true iff there is a terminating instruction at the end of b.
func (b *Block) IsTerminated() bool {
	return b.jump != nil
}
