package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"runos_chain/x/configstore/types"
)

var _ = strconv.Itoa(0)

func CmdGetPort() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-port [dpid] [mac]",
		Short: "Broadcast message get-port",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDpid, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argMac := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgGetPort(
				clientCtx.GetFromAddress().String(),
				argDpid,
				argMac,
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
