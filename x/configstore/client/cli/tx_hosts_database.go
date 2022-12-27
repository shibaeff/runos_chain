package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"runos_chain/x/configstore/types"
)

func CmdCreateHostsDatabase() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-hosts-database [dpid] [mac] [inport]",
		Short: "Create a new hostsDatabase",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexDpid := args[0]
			indexMac := args[1]

			// Get value arguments
			argInport := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateHostsDatabase(
				clientCtx.GetFromAddress().String(),
				indexDpid,
				indexMac,
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

func CmdUpdateHostsDatabase() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-hosts-database [dpid] [mac] [inport]",
		Short: "Update a hostsDatabase",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexDpid := args[0]
			indexMac := args[1]

			// Get value arguments
			argInport := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateHostsDatabase(
				clientCtx.GetFromAddress().String(),
				indexDpid,
				indexMac,
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

func CmdDeleteHostsDatabase() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-hosts-database [dpid] [mac]",
		Short: "Delete a hostsDatabase",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexDpid := args[0]
			indexMac := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteHostsDatabase(
				clientCtx.GetFromAddress().String(),
				indexDpid,
				indexMac,
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
