package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var namespaceCmd = &cobra.Command{
	Use:   "namespace",
	Short: "Operate on namespaces",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(namespaceCmd).Standalone()
	rootCmd.AddCommand(namespaceCmd)
}
