package keeper

import (
	"fistchain/x/fistchain/types"
)

var _ types.QueryServer = Keeper{}
