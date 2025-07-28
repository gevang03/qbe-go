package qbe

import "regexp"

// QBE doc does not specify exactly what a name is, but this is consistent with the current implementation.
var validNamePattern = regexp.MustCompile(`^[a-zA-Z._][a-zA-Z$._0-9]*$`)

func validateName(s string) bool {
	return validNamePattern.MatchString(s)
}
