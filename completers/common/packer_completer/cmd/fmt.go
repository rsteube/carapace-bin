package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var fmtCmd = &cobra.Command{
	Use:   "fmt",
	Short: "Rewrites HCL2 config files to canonical format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fmtCmd).Standalone()

	fmtCmd.Flags().BoolS("check", "check", false, "Check if the input is formatted.")
	fmtCmd.Flags().BoolS("diff", "diff", false, "Display diffs of formatting change")
	fmtCmd.Flags().BoolS("recursive", "recursive", false, "Also process files in subdirectories.")
	fmtCmd.Flags().BoolS("write", "write", false, "Don't write to source files")
	rootCmd.AddCommand(fmtCmd)

	carapace.Gen(fmtCmd).PositionalCompletion(
		carapace.ActionFiles(".json", ".hcl"),
	)
}
