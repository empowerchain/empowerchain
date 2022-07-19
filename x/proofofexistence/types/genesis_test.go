package types_test

import (
	"testing"

	"github.com/empowerchain/empowerchain/x/proofofexistence/types"
	"github.com/stretchr/testify/require"
)

func TestDefaultGenesisIsCorrect(t *testing.T) {
	genesis := types.DefaultGenesis()
	require.Equal(t, &types.GenesisState{
		Params:    types.Params{},
		ProofList: []types.Proof{},
	}, genesis)
}

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				ProofList: []types.Proof{
					{
						Hash: "0",
					},
					{
						Hash: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated proof",
			genState: &types.GenesisState{
				ProofList: []types.Proof{
					{
						Hash: "0",
					},
					{
						Hash: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
