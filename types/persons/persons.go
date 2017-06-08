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
