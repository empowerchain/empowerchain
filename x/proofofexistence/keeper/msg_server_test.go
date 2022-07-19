package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/empowerchain/empowerchain/testutil/keeper"
	"github.com/empowerchain/empowerchain/x/proofofexistence/keeper"
	"github.com/empowerchain/empowerchain/x/proofofexistence/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := keepertest.ProofofexistenceKeeper(t)
	return keeper.NewMsgServerImpl(*k), *k, sdk.WrapSDKContext(ctx)
}

func TestCreateProof(t *testing.T) {
	msgServer, _, ctx := setupMsgServer(t)
	_, err := msgServer.CreateProof(ctx, &types.MsgCreateProof{
		Reporter: "empower10m8qakrveaxvcr5085c73de30ff5umdem2ua5u",
		Hash:     "cZiMTY4IA7pFGfCyhkwTMcFKGJC/hpTiUTeRd7/ttcM=",
	})
	require.NoError(t, err)
}
