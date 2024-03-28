package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var node_editCmd = &cobra.Command{
	Use:     "edit NODE",
	Short:   "[EXPERIMENTAL] Edit node(s).",
	Aliases: []string{"update"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(node_editCmd).Standalone()

	node_editCmd.Flags().StringSlice("port-add", []string{}, "[EXPERIMENTAL] (serverlb only!) Map ports from the node container to the host (Format: `[HOST:][HOSTPORT:]CONTAINERPORT[/PROTOCOL][@NODEFILTER]`)")
	nodeCmd.AddCommand(node_editCmd)
}
