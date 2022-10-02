package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var binCmd = &cobra.Command{
	Use:   "bin",
	Short: "Get the path to the Turbo binary",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(binCmd).Standalone()

	rootCmd.AddCommand(binCmd)
}
