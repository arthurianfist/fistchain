package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// module name
	ModuleName = "fistchain"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName
)

// HelloPrefix is a prefix for keys that store Hello messages
var HelloPrefix = []byte("hello-")

// Hello struct
type Hello struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	Name    string         `json:"name" yaml:"name"`
	Message string         `json:"message" yaml:"message"`
}

// Query Result Payload for a hello
type QueryHelloResult struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	Name    string         `json:"name" yaml:"name"`
	Message string         `json:"message" yaml:"message"`
}

// Query Result Payload for all hello
type QueryAllHelloResult []QueryHelloResult

// NewHello creates a new instance of Hello
func NewHello(creator sdk.AccAddress, name string, message string) Hello {
	return Hello{
		Creator: creator,
		Name:    name,
		Message: message,
	}
}

// Implement fmt.Stringer
func (h Hello) String() string {
	return string(h.Message)
}

// QuerySayHelloRequest is the request type for the SayHello query endpoint
type QuerySayHelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

// QuerySayHelloResponse is the response type for the SayHello query endpoint
type QuerySayHelloResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}
