package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var replCmd = &cobra.Command{
	Use:   "repl",
	Short: "Get a node repl with context of Gatsby environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(replCmd).Standalone()

	rootCmd.AddCommand(replCmd)
}
