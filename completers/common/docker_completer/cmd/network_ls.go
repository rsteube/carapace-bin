package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var network_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List networks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_lsCmd).Standalone()

	network_lsCmd.Flags().StringP("filter", "f", "", "Provide filter values (e.g. 'driver=bridge')")
	network_lsCmd.Flags().String("format", "", "Pretty-print networks using a Go template")
	network_lsCmd.Flags().Bool("no-trunc", false, "Do not truncate the output")
	network_lsCmd.Flags().BoolP("quiet", "q", false, "Only display network IDs")
	networkCmd.AddCommand(network_lsCmd)
}
