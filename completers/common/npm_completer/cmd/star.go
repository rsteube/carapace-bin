package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var starCmd = &cobra.Command{
	Use:   "star",
	Short: "Mark your favorite packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(starCmd).Standalone()
	starCmd.Flags().Bool("unicode", false, "use unicode in output")

	rootCmd.AddCommand(starCmd)

	carapace.Gen(starCmd).PositionalAnyCompletion(
		action.ActionPackageNames(starCmd),
	)
}
