package cli

import (
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/medibloc/panacea-core/v2/x/datadeal/types"
	"github.com/spf13/cobra"
)

func CmdSellData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sell-data [path]",
		Short: "Sell data",
		Long:  "[data-hash] is a hex-encoded string obtained by hashing the original data through the SHA256 hash function",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg, err := readMsgSellDataFrom(args[0])
			if err != nil {
				return err
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func readMsgSellDataFrom(path string) (*types.MsgSellData, error) {
	var msg types.MsgSellData

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if err := jsonpb.Unmarshal(file, &msg); err != nil {
		return nil, err
	}

	return &msg, nil
}
