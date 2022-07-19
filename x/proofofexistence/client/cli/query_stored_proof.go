package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/empowerchain/empowerchain/x/proofofexistence/types"
	"github.com/spf13/cobra"
)

func CmdShowProof() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-stored-proof [hash]",
		Short: "shows a Proof",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argHash := args[0]

			params := &types.QueryGetProofRequest{
				Hash: argHash,
			}

			res, err := queryClient.Proof(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
