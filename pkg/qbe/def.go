package qbe

import (
	"fmt"
	"strings"
)

type Definition interface {
	isDefinition()
}

// Linkage represents the linkage of a definition ([Function] or [Data])
type Linkage struct {
	Export       bool   // Definition is visible
	Thread       bool   // Data definition has thread local storage
	SectionName  string // Section to store definition
	SectionFlags string // Other flags specified for a section
}

// PrivateLinkage returns a [Linkage] for a non exported, non thread local storage
// and no section information specified.
func PrivateLinkage() Linkage {
	return Linkage{false, false, "", ""}
}

// String converts l to a string compatible with QBE code.
func (l *Linkage) String() string {
	var parts []string
	if l.Export {
		parts = append(parts, "export")
	}
	if l.Thread {
		parts = append(parts, "thread")
	}
	if l.SectionName != "" {
		parts = append(parts, fmt.Sprintf("section %q", l.SectionName))
		if l.SectionFlags != "" {
			parts = append(parts, fmt.Sprintf("%q", l.SectionFlags))
		}
	} else {
		if l.SectionFlags != "" {
			panic("section flags set without a section name")
		}
	}
	return strings.Join(parts, " ")
}
