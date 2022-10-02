package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new Gatsby project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(newCmd).Standalone()

	rootCmd.AddCommand(newCmd)

	carapace.Gen(newCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionDirectories(),
	)
}
