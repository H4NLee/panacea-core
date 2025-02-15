package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/medibloc/panacea-core/v2/x/datadeal/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group datadeal queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdGetDeals())
	cmd.AddCommand(CmdGetDeal())
	cmd.AddCommand(CmdGetDataSale())
	cmd.AddCommand(CmdGetDataSales())
	cmd.AddCommand(CmdGetDataVerificationVote())
	cmd.AddCommand(CmdGetDataDeliveryVote())

	return cmd
}
