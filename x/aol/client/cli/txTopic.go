package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/medibloc/panacea-core/x/aol/types"
	"github.com/spf13/cobra"
)

const (
	flagDescription = "description"
	flagMoniker     = "moniker"
	flagFeePayer    = "payer"
)

func CmdCreateTopic() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-topic [topicName]",
		Short: "Creates a new topic",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			description, err := cmd.Flags().GetString(flagDescription)
			if err != nil {
				description = ""
			}

			msg := types.NewMsgCreateTopic(args[0], description, clientCtx.GetFromAddress().String())
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(flagDescription, "", "description of topic")

	return cmd
}
