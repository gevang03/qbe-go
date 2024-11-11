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

func (b *Block) InsertJmp(label Label) {
	b.jump = jmp{label}
}

func (b *Block) InsertJnz(value value.Value, notZero, zero Label) {
	b.jump = jnz{value, notZero, zero}
}

func (b *Block) InsertRet(value value.Value) {
	b.jump = ret{value}
}

func (b *Block) InsertHlt() {
	b.jump = hlt{}
}
