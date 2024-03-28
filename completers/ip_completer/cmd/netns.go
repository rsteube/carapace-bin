package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var netnsCmd = &cobra.Command{
	Use:   "netns",
	Short: "manage network namespaces",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(netnsCmd).Standalone()

	rootCmd.AddCommand(netnsCmd)
}
