package fistchain_test

import (
	"testing"

	keepertest "fistchain/testutil/keeper"
	"fistchain/testutil/nullify"
	"fistchain/x/fistchain"
	"fistchain/x/fistchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FistchainKeeper(t)
	fistchain.InitGenesis(ctx, *k, genesisState)
	got := fistchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
