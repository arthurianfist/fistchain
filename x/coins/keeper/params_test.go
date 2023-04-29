package keeper_test

import (
	"testing"

	testkeeper "fistchain/testutil/keeper"
	"fistchain/x/coins/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CoinsKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
