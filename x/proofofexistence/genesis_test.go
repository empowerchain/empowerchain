package proofofexistence_test

import (
	"testing"
	"time"

	keepertest "github.com/empowerchain/empowerchain/testutil/keeper"
	"github.com/empowerchain/empowerchain/testutil/nullify"
	"github.com/empowerchain/empowerchain/x/proofofexistence"
	"github.com/empowerchain/empowerchain/x/proofofexistence/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ProofList: []types.Proof{
			{
				Hash:      "0",
				Timestamp: time.UnixMilli(42).UTC(),
				Reporter:  "empower10m8qakrveaxvcr5085c73de30ff5umdem2ua5u",
			},
			{
				Hash:      "1",
				Timestamp: time.UnixMilli(123).UTC(),
				Reporter:  "empower10m8qakrveaxvcr5085c73de30ff5umdem2ua5u",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ProofofexistenceKeeper(t)
	proofofexistence.InitGenesis(ctx, *k, genesisState)
	got := proofofexistence.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	for i, expected := range genesisState.ProofList {
		actual := got.ProofList[i]
		require.Equal(t, expected.Hash, actual.Hash)
		require.Equal(t, expected.Reporter, actual.Reporter)
		require.Equal(t, expected.Timestamp, actual.Timestamp)
	}

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ProofList, got.ProofList)
	// this line is used by starport scaffolding # genesis/test/assert
}
