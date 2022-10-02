package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var stdlibCmd = &cobra.Command{
	Use:   "stdlib",
	Short: "Displays the stdlib available in the .envrc execution context",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stdlibCmd).Standalone()

	rootCmd.AddCommand(stdlibCmd)
}
