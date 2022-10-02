package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compileCmd = &cobra.Command{
	Use:   "compile",
	Short: "Compile Dart to various formats",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compileCmd).Standalone()

	compileCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	rootCmd.AddCommand(compileCmd)

	carapace.Gen(compileCmd).PositionalCompletion(
		carapace.ActionFiles(".dart"),
	)
}
