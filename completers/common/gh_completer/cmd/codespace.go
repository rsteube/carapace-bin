package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var codespaceCmd = &cobra.Command{
	Use:     "codespace",
	Short:   "Connect to and manage codespaces",
	Aliases: []string{"cs"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codespaceCmd).Standalone()
	rootCmd.AddCommand(codespaceCmd)
}
