package inst

import (
	"fmt"
	"strings"

	"github.com/gevang03/qbe-go/types"
	"github.com/gevang03/qbe-go/value"
)

// A PhiSrc represents a source value from a basic block. Used as source to the phi instruction.
type PhiSrc struct {
	Label       // The label of the basic block the value originates.
	value.Value // The value from the incoming basic block.
}

type phi struct {
	Dest value.Temporary
	Type types.BaseType
	Srcs []PhiSrc
}

// InsertPhi adds a phi instruction to b to temporary dest with type_ type and sources as arguments.
func (b *Block) InsertPhi(dest value.Temporary, type_ types.BaseType, sources ...PhiSrc) {
	phi_ := phi{dest, type_, sources}
	b.phis = append(b.phis, phi_)
}

func (phi phi) String() string {
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
