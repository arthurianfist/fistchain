package keeper

import (
	"fistchain/x/addresses/types"
)

var _ types.QueryServer = Keeper{}
