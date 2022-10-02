package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "Start a local development server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devCmd).Standalone()
	devCmd.Flags().StringP("listen", "l", "0.0.0.0:3000", "Specify a URI endpoint on which to listen")

	rootCmd.AddCommand(devCmd)

	carapace.Gen(devCmd).FlagCompletion(carapace.ActionMap{
		"listen": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValues("0.0.0.0", "127.0.0.1").Invoke(c).Suffix(":").ToA()
			case 1:
				return net.ActionPorts()
			default:
				return carapace.ActionValues()
			}
		}),
	})
}
