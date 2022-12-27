package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"runos_chain/x/configstore/types"
)

func CmdListHostsDatabase() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-hosts-database",
		Short: "list all hostsDatabase",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllHostsDatabaseRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.HostsDatabaseAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowHostsDatabase() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-hosts-database [dpid] [mac]",
		Short: "shows a hostsDatabase",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argDpid := args[0]
			argMac := args[1]

			params := &types.QueryGetHostsDatabaseRequest{
				Dpid: argDpid,
				Mac:  argMac,
			}

			res, err := queryClient.HostsDatabase(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
