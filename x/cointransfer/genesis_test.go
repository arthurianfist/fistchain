package cointransfer_test

import (
	"testing"

	keepertest "fistchain/testutil/keeper"
	"fistchain/testutil/nullify"
	"fistchain/x/cointransfer"
	"fistchain/x/cointransfer/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CointransferKeeper(t)
	cointransfer.InitGenesis(ctx, *k, genesisState)
	got := cointransfer.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
