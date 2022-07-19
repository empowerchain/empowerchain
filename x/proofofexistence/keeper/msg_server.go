package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/empowerchain/empowerchain/x/proofofexistence/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) CreateProof(goCtx context.Context, msg *types.MsgCreateProof) (*types.MsgCreateProofResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.CreateNewProof(ctx, msg.Hash, msg.Reporter); err != nil {
		return nil, err
	}

	return &types.MsgCreateProofResponse{}, nil
}
