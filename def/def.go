package def

import (
	"fmt"
	"strings"
)

type Definition interface {
	isDefinition()
}

type Linkage struct {
	Export       bool
	Thread       bool
	SectionName  string
	SectionFlags string
}

func PrivateLinkage() Linkage {
	return Linkage{false, false, "", ""}
}

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
