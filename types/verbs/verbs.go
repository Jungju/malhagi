package verbs

import "strings"

// Type is verbs's type
type Type int

// const
const (
	None Type = iota
	Be
	General
)

func (t Type) String() string {
	return strings.ToLower(_TypeValueToName[t])
}

//Ids ...
func Ids() []int {
	return []int{
		int(Be),
		int(General),
	}
}
