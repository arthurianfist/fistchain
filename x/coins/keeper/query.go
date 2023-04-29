package keeper

import (
	"fistchain/x/coins/types"
)

var _ types.QueryServer = Keeper{}
