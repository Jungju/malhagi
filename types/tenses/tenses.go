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

//Ids ...
func Ids() []int {
	return []int{
		int(Past),
		int(Present),
		int(Future),
	}
}
