package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/medibloc/panacea-core/x/aol/types"
)

func CmdAddRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-record [ownerAddress] [topicName] [key] [value]",
		Short: "Add a new record",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			feePayerAddress, err := cmd.Flags().GetString(flagFeePayer)
			if err != nil {
				feePayerAddress = ""
			}

			msg := types.NewMsgAddRecord(args[1], []byte(args[2]), []byte(args[3]), clientCtx.GetFromAddress().String(), args[0], feePayerAddress)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(flagFeePayer, "", "optional address to pay for the fee")

	return cmd
}
