package types

import (
	"encoding/base64"
	"fmt"
	"github.com/tendermint/tendermint/crypto"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/empowerchain/empowerchain/testutil/sample"
	"github.com/stretchr/testify/require"
)

const testData = "This is just some random test data to be hashed. 42."

func TestMsgCreateProof_ValidateBasic(t *testing.T) {
	hash := crypto.Sha256([]byte(testData))
	hashb64 := base64.StdEncoding.EncodeToString(hash)

	tests := []struct {
		name string
		msg  MsgCreateProof
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateProof{
				Reporter: "invalid_address",
				Hash:     hashb64,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address and valid hash",
			msg: MsgCreateProof{
				Reporter: sample.AccAddress(),
				Hash:     hashb64,
			},
		},
		{
			name: "invalid base64",
			msg: MsgCreateProof{
				Reporter: sample.AccAddress(),
				Hash:     "this is not base64!",
			},
			err: ErrInvalidHash,
		},
		{
			name: "invalid hash",
			msg: MsgCreateProof{
				Reporter: sample.AccAddress(),
				Hash:     base64.StdEncoding.EncodeToString([]byte("this is not a hash!")),
			},
			err: ErrInvalidHash,
		},
	}
	for _, tt := range tests {
		fmt.Println(tt.name, tt.msg.Hash)
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
