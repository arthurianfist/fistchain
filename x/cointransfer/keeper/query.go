package keeper

import (
	"fistchain/x/cointransfer/types"
)

var _ types.QueryServer = Keeper{}
