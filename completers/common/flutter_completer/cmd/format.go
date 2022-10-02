package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var formatCmd = &cobra.Command{
	Use:   "format",
	Short: "Idiomatically format Dart source code",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(formatCmd).Standalone()

	formatCmd.Flags().Bool("fix", false, "Apply all style fixes.")
	formatCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	formatCmd.Flags().StringP("line-length", "l", "", "Wrap lines longer than this.")
	formatCmd.Flags().StringP("output", "o", "", "Set where to write formatted output.")
	formatCmd.Flags().Bool("set-exit-if-changed", false, "Return exit code 1 if there are any formatting changes.")
	formatCmd.Flags().BoolP("verbose", "v", false, "Show all options and flags with --help.")
	rootCmd.AddCommand(formatCmd)

	carapace.Gen(formatCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionValuesDescribed(
			"json", "Print code and selection as JSON.",
			"none", "Discard output.",
			"show", "Print code to terminal.",
			"write", "Overwrite formatted files on disk.",
		),
	})

	carapace.Gen(formatCmd).PositionalAnyCompletion(
		carapace.ActionFiles(".dart"),
	)
}
