package formats

import "strings"

// Type is format's type
type Type int

// const
const (
	None Type = iota
	Plain
	Negative
	Future
	Question
)

func (t Type) String() string {
	return strings.ToLower(_TypeValueToName[t])
}

//Ids ...
func Ids() []int {
	return []int{
		int(Plain),
		int(Negative),
		int(Future),
		int(Question),
	}
}
