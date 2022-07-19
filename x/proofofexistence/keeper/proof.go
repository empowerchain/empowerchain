package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/empowerchain/empowerchain/x/proofofexistence/types"
)

// CreateNewProof create a new proof from its hash
func (k Keeper) CreateNewProof(ctx sdk.Context, hash string, reporter string) error {
	proof := types.Proof{
		Hash:      hash,
		Timestamp: ctx.BlockTime(),
		Reporter:  reporter,
	}

	return k.SetProof(ctx, proof)
}

// SetProof persists a specific proof. Usually, you want to use CreateNewProof
// SetProof is more useful for things like genesis setup
func (k Keeper) SetProof(ctx sdk.Context, proof types.Proof) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProofKeyPrefix))
	key := types.ProofKey(
		proof.Hash,
	)

	if store.Has(key) {
		return sdkerrors.Wrap(types.ErrHashExists, proof.Hash)
	}

	b := k.cdc.MustMarshal(&proof)
	store.Set(key, b)

	return nil
}

// GetProof returns a proof from its 'index' (hash in this case)
func (k Keeper) GetProof(
	ctx sdk.Context,
	hash string,

) (val types.Proof, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProofKeyPrefix))

	b := store.Get(types.ProofKey(
		hash,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllProof returns all proofs
func (k Keeper) GetAllProof(ctx sdk.Context) (list []types.Proof) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProofKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Proof
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
