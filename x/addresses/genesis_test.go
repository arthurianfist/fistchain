package addresses_test

import (
	"testing"

	keepertest "fistchain/testutil/keeper"
	"fistchain/testutil/nullify"
	"fistchain/x/addresses"
	"fistchain/x/addresses/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AddressesKeeper(t)
	addresses.InitGenesis(ctx, *k, genesisState)
	got := addresses.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
