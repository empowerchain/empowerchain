package keeper_test

import (
	"testing"

	testkeeper "github.com/empowerchain/empowerchain/testutil/keeper"
	"github.com/empowerchain/empowerchain/x/proofofexistence/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ProofofexistenceKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
