package keeper_test

import (
	"testing"

	testkeeper "fistchain/testutil/keeper"
	"fistchain/x/addresses/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.AddressesKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
