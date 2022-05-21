package kind

import "fmt"

func (k Kind) MarshalJSON() ([]byte, error) {
	// should have done this earlier!
	return []byte(fmt.Sprintf("%q", k.StringValue())), nil
}
