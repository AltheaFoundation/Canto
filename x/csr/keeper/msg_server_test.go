package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/Canto-Network/Canto/v2/x/csr/types"
    "github.com/Canto-Network/Canto/v2/x/csr/keeper"
    keepertest "github.com/Canto-Network/Canto/v2/testutil/keeper"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CsrKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
