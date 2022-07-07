package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/empowerchain/empowerchain/testutil/keeper"
	"github.com/empowerchain/empowerchain/testutil/nullify"
	"github.com/empowerchain/empowerchain/x/proofofexistence/keeper"
	"github.com/empowerchain/empowerchain/x/proofofexistence/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNStoredProof(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.StoredProof {
	items := make([]types.StoredProof, n)
	for i := range items {
		items[i].Hash = strconv.Itoa(i)

		keeper.SetStoredProof(ctx, items[i])
	}
	return items
}

func TestStoredProofGet(t *testing.T) {
	keeper, ctx := keepertest.ProofofexistenceKeeper(t)
	items := createNStoredProof(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetStoredProof(ctx,
			item.Hash,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}