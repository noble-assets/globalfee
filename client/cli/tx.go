package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noble-assets/globalfee/types"
	"github.com/spf13/cobra"
)

func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Transactions commands for the %s module", types.ModuleName),
		DisableFlagParsing:         false,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(TxUpdateGasPrices())

	return cmd
}

func TxUpdateGasPrices() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "update-gas-prices [gas-prices ...]",
		Short:   "Update the minimum required gas prices for non-bypassed messages",
		Example: "update-gas-prices 0.1uusdc 0.09ueure",
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			var gasPrices sdk.DecCoins
			for _, arg := range args {
				gasPrice, err := sdk.ParseDecCoin(arg)
				if err != nil {
					return err
				}

				gasPrices = append(gasPrices, gasPrice)
			}

			msg := &types.MsgUpdateGasPrices{
				Signer:    clientCtx.GetFromAddress().String(),
				GasPrices: gasPrices,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
