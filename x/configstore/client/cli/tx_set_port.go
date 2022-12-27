package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"runos_chain/x/configstore/types"
)

var _ = strconv.Itoa(0)

func CmdSetPort() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-port [dpid] [mac] [inport]",
		Short: "Broadcast message setPort",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDpid := args[0]
			argMac := args[1]
			argInport := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetPort(
				clientCtx.GetFromAddress().String(),
				argDpid,
				argMac,
				argInport,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
