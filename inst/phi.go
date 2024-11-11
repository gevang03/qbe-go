package inst

import (
	"fmt"
	"strings"

	"github.com/gevang03/qbe-go/types"
	"github.com/gevang03/qbe-go/value"
)

type PhiSrc struct {
	Label
	value.Value
}

type Phi struct {
	Dest value.Temporary
	Type types.BaseType
	Srcs []PhiSrc
}

func (b *Block) InsertPhi(dest value.Temporary, type_ types.BaseType, values ...PhiSrc) {
	phi := Phi{dest, type_, values}
	b.phis = append(b.phis, phi)
}

func (phi Phi) String() string {
	parts := []string{phi.Dest.String(), fmt.Sprintf("=%v", phi.Type), "phi"}
	for i, value := range phi.Srcs {
		if i == len(phi.Srcs)-1 {
			parts = append(parts, fmt.Sprintf("%v %v", value.Label, value.Value))
		} else {
			parts = append(parts, fmt.Sprintf("%v %v,", value.Label, value.Value))
		}
	}
	return strings.Join(parts, " ")
}
