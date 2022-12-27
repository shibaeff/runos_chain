package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"runos_chain/x/configstore/types"
)

var _ = strconv.Itoa(0)

func CmdGetPort() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-port [dpid] [mac]",
		Short: "Query getPort",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDpid := args[0]
			reqMac := args[1]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetPortRequest{

				Dpid: reqDpid,
				Mac:  reqMac,
			}

			res, err := queryClient.GetPort(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
