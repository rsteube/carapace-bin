package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/vi_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vi",
	Short: "screen oriented (visual) display editor based on ex",
	Long:  "https://en.wikipedia.org/wiki/Vi",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("L", "L", "", "temporary file to open in recovery mode")
	rootCmd.Flags().BoolS("R", "R", false, "Files are opened read-only when this option is given")
	rootCmd.Flags().BoolS("V", "V", false, "Echo command input to standard error, unless it originates from a terminal")
	rootCmd.Flags().StringS("c", "c", "", "Execute command when editing begins")
	rootCmd.Flags().BoolS("l", "l", false, "Start in a special mode useful for the Lisp programming language")
	rootCmd.Flags().StringS("r", "r", "", "temporary file to open in recovery mode")
	rootCmd.Flags().BoolS("s", "s", false, "Script mode, all feedback for interactive editing is disabled")
	rootCmd.Flags().StringS("t", "t", "", "Read the tags file, then choose the file and position specified by tagstring for editing")
	rootCmd.Flags().StringS("w", "w", "", "Specify the size of the editing window for visual mode")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"L": action.ActionTemporaryFiles(),
		"r": action.ActionTemporaryFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
