package keeper_test

import (
	"encoding/base64"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/empowerchain/empowerchain/testutil/sample"
	"github.com/empowerchain/empowerchain/x/proofofexistence/keeper"
	"github.com/empowerchain/empowerchain/x/proofofexistence/types"
	"github.com/tendermint/tendermint/crypto"
	"strconv"
	"testing"
	"time"

	keepertest "github.com/empowerchain/empowerchain/testutil/keeper"
	"github.com/stretchr/testify/require"
)

func createNProof(k *keeper.Keeper, ctx sdk.Context, n int) []types.Proof {
	items := make([]types.Proof, n)
	for i := range items {
		hash, reporter := createTestProofData(strconv.Itoa(i))
		items[i].Hash = hash
		items[i].Reporter = reporter
		items[i].Timestamp = time.Unix(1657454750+int64(i*25), 0).UTC()

		if err := k.SetProof(ctx, items[i]); err != nil {
			panic(err)
		}
	}
	return items
}

func createTestProofData(data string) (hash string, reporter string) {
	hashb := crypto.Sha256([]byte(data))
	hash = base64.StdEncoding.EncodeToString(hashb)
	reporter = sample.AccAddress()

	return hash, reporter
}

func TestCreateNewProof(t *testing.T) {
	k, ctx := keepertest.ProofofexistenceKeeper(t)
	hash, reporter := createTestProofData("my test data")

	err := k.CreateNewProof(ctx, hash, reporter)
	require.NoError(t, err)

	proof, found := k.GetProof(ctx, hash)
	require.True(t, found)
	require.Equal(t, reporter, proof.Reporter)
	require.Equal(t, hash, proof.Hash)
	require.Equal(t, ctx.BlockTime(), proof.Timestamp)
}

func TestCreateNewProofWithExistingHash(t *testing.T) {
	k, ctx := keepertest.ProofofexistenceKeeper(t)
	hash, reporter := createTestProofData("42")

	err := k.CreateNewProof(ctx, hash, reporter)
	require.NoError(t, err)

	err = k.CreateNewProof(ctx, hash, reporter)
	require.ErrorIs(t, err, types.ErrHashExists)
}

func TestGetProof(t *testing.T) {
	k, ctx := keepertest.ProofofexistenceKeeper(t)
	hash, reporter := createTestProofData("69")

	_, found := k.GetProof(ctx, hash)
	require.False(t, found)

	err := k.CreateNewProof(ctx, hash, reporter)
	require.NoError(t, err)

	proof, found := k.GetProof(ctx, hash)
	require.True(t, found)
	require.Equal(t, reporter, proof.Reporter)
	require.Equal(t, hash, proof.Hash)
	require.Equal(t, ctx.BlockTime(), proof.Timestamp)
}

func TestGetAllProof(t *testing.T) {
	k, ctx := keepertest.ProofofexistenceKeeper(t)
	proofs := createNProof(k, ctx, 36)
	got := k.GetAllProof(ctx)

	require.Equal(t, len(proofs), len(got))

	for _, p := range proofs {
		found := false
		for _, g := range got {
			if p.Hash == g.Hash && p.Reporter == g.Reporter && p.Timestamp == g.Timestamp {
				found = true
				break
			}
		}
		require.Truef(t, found, "Proof %v was not found", p)
	}
}
