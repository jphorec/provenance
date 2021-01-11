package cli

import (
	"fmt"
	"strings"

	"github.com/provenance-io/provenance/x/name/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// The flag for created restricted names
const flagRestricted = "restrict"

// NewTxCmd is the top-level command for name CLI transactions.
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Transaction commands for the name module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	txCmd.AddCommand(
		GetBindNameCmd(),
	)
	return txCmd
}

// The CLI command for binding a name to an address.
func GetBindNameCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bind [name] [address] [root]",
		Short: "Bind a name to an address under the given root name in the provenance blockchain",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Bind a name under an existing name in the provenance blockchain:

Example:
$ %s tx name bind sample pb1skjwj5whet0lpe65qaq4rpq03hjxlwd9nf39lk root.example
`,
				version.AppName,
			)),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			address, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}
			msg := types.NewMsgBindNameRequest(
				types.NewNameRecord(
					strings.ToLower(args[0]),
					address,
					viper.GetBool(flagRestricted),
				),
				types.NewNameRecord(
					strings.ToLower(args[2]),
					clientCtx.FromAddress,
					false,
				),
			)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().BoolP(flagRestricted, "r", true, "Restrict creation of child names to owner only")

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
