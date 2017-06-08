package tenses

import "strings"

// Type is tenses type
type Type int

// const
const (
	None Type = iota
	Past
	Present
	Future
)

func (t Type) String() string {
	return strings.ToLower(_TypeValueToName[t])
}
