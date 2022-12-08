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

func CmdSetPort() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-port [dpid] [mac] [in-port]",
		Short: "Broadcast message set-port",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDpid, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argMac := args[1]
			argInPort, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetPort(
				clientCtx.GetFromAddress().String(),
				argDpid,
				argMac,
				argInPort,
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
