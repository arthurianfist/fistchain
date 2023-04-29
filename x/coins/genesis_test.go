package coins_test

import (
	"testing"

	keepertest "fistchain/testutil/keeper"
	"fistchain/testutil/nullify"
	"fistchain/x/coins"
	"fistchain/x/coins/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CoinsKeeper(t)
	coins.InitGenesis(ctx, *k, genesisState)
	got := coins.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
