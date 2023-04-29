package keeper_test

import (
	"context"
	"testing"

	keepertest "fistchain/testutil/keeper"
	"fistchain/x/coins/keeper"
	"fistchain/x/coins/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CoinsKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
