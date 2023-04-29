package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// HelloPrefix is a prefix for hello message keys
const HelloPrefix = "hello-"

// Hello struct
type Hello struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	Name    string         `json:"name" yaml:"name"`
	Message string         `json:"message" yaml:"message"`
}

func (h Hello) String() string {
	return fmt.Sprintf(`Hello
	Creator: %s
	Name: %s
	Message: %s`, h.Creator.String(), h.Name, h.Message)
}

// Hellos slice of hello
type Hellos []Hello

// Query Result Payload for a hello query
type QueryResHello struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}
