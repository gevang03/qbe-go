package validation

import "regexp"

// QBE doc does not specify exactly what a name is, but this is consistent with the current implementation.
var validPattern = regexp.MustCompile(`^[a-zA-Z._][a-zA-Z$._0-9]*$`)

func Validate(s string) bool {
	return validPattern.MatchString(s)
}
