package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/terraform_completer/cmd/action"
	"github.com/spf13/cobra"
)

var fmtCmd = &cobra.Command{
	Use:   "fmt",
	Short: "Reformat your configuration in the standard style",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fmtCmd).Standalone()

	fmtCmd.Flags().BoolS("check", "check", false, "Check if the input is formatted")
	fmtCmd.Flags().BoolS("diff", "diff", false, "Display diffs of formatting changes")
	fmtCmd.Flags().StringS("list", "list", "", "Don't list files whose formatting differs")
	fmtCmd.Flags().BoolS("no-color", "no-color", false, "If specified, output won't contain any color")
	fmtCmd.Flags().BoolS("recursive", "recursive", false, "Also process files in subdirectories")
	fmtCmd.Flags().StringS("write", "write", "", "Don't write to source files")
	rootCmd.AddCommand(fmtCmd)

	carapace.Gen(fmtCmd).FlagCompletion(carapace.ActionMap{
		"list":  action.ActionBool(),
		"write": action.ActionBool(),
	})

	carapace.Gen(fmtCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
