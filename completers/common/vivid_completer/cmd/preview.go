package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/vivid_completer/cmd/action"
	"github.com/spf13/cobra"
)

var previewCmd = &cobra.Command{
	Use:   "preview",
	Short: "Preview a given theme",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(previewCmd).Standalone()

	previewCmd.Flags().BoolP("help", "h", false, "Prints help information")
	rootCmd.AddCommand(previewCmd)

	carapace.Gen(previewCmd).PositionalCompletion(
		action.ActionThemes(),
	)
}
