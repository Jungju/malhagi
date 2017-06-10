package persons

import "strings"

// Type is person's type
type Type int

// const
const (
	None Type = iota
	I
	We
	You
	They
	He
	She
	It
	Special
)

func (t Type) String() string {
	return strings.ToLower(_TypeValueToName[t])
}

//Ids ...
func Ids() []int {
	return []int{
		int(I),
		int(We),
		int(You),
		int(They),
		int(He),
		int(She),
		int(It),
		int(Special),
	}
}
