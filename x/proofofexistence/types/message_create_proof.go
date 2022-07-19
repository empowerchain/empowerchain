package types

import (
	"encoding/base64"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateProof = "create_proof"

var _ sdk.Msg = &MsgCreateProof{}

func NewMsgCreateProof(reporter string, hash string) *MsgCreateProof {
	return &MsgCreateProof{
		Reporter: reporter,
		Hash:     hash,
	}
}

func (msg *MsgCreateProof) Route() string {
	return RouterKey
}

func (msg *MsgCreateProof) Type() string {
	return TypeMsgCreateProof
}

func (msg *MsgCreateProof) GetSigners() []sdk.AccAddress {
	reporter, err := sdk.AccAddressFromBech32(msg.Reporter)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{reporter}
}

func (msg *MsgCreateProof) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateProof) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Reporter)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid reporter address (%s)", err)
	}

	hashBytes, err := base64.StdEncoding.DecodeString(msg.Hash)
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalidHash, "Hash %s is not base64 encoded", msg.Hash)
	}

	if len(hashBytes) != 32 {
		return sdkerrors.Wrapf(ErrInvalidHash, "Base 64-decoded hash %s is not SHA-256", msg.Hash)
	}

	return nil
}
