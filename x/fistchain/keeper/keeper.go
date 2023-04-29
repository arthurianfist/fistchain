package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/fistchain/fistchain/x/fistchain/types"
)

// Keeper of the fistchain store
type Keeper struct {
	storeKey sdk.StoreKey
	cdc      codec.Marshaler
}

// NewKeeper creates a fistchain keeper
func NewKeeper(cdc codec.Marshaler, key sdk.StoreKey) *Keeper {
	return &Keeper{
		storeKey: key,
		cdc:      cdc,
	}
}

// SayHello returns the greeting message
func (k Keeper) SayHello(ctx context.Context, req *types.QuerySayHelloRequest) (*types.QuerySayHelloResponse, error) {
	return &types.QuerySayHelloResponse{Message: "Hello, " + req.Name + "!"}, nil
}
