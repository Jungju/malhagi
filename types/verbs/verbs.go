package verbs

import "strings"

// Type is verbs's type
type Type int

// const
const (
	None Type = iota
	BeVerb
	GeneralVerb
)

func (t Type) String() string {
	return strings.ToLower(_TypeValueToName[t])
}
