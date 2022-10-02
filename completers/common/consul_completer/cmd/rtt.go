package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/consul_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rttCmd = &cobra.Command{
	Use:   "rtt",
	Short: "Estimates network round trip time between nodes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rttCmd).Standalone()
	addClientFlags(rttCmd)

	rttCmd.Flags().Bool("wan", false, "Use WAN coordinates instead of LAN coordinates.")
	rootCmd.AddCommand(rttCmd)

	carapace.Gen(rttCmd).PositionalCompletion(
		action.ActionNodes(rttCmd),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionNodes(rttCmd).Invoke(c).Filter(c.Args[:1]).ToA()
		}),
	)
}
