package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var unstarCmd = &cobra.Command{
	Use:   "unstar",
	Short: "Remove an item from your favorite packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unstarCmd).Standalone()
	unstarCmd.Flags().String("otp", "", "one-time password")
	unstarCmd.Flags().Bool("unicode", false, "use unicode in output")

	rootCmd.AddCommand(unstarCmd)

	carapace.Gen(unstarCmd).PositionalAnyCompletion(
		action.ActionPackageNames(unstarCmd), // TODO use list of favorites
	)
}
