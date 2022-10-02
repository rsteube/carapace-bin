package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate Go files by processing source",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generateCmd).Standalone()

	generateCmd.Flags().BoolS("n", "n", false, "print commands that would be executed")
	generateCmd.Flags().StringS("run", "run", "", "specifies a regular expression to select matching directives")
	generateCmd.Flags().BoolS("v", "v", false, "print the names of packages and files as they are processed")
	generateCmd.Flags().BoolS("x", "x", false, "print commands as they are executed")
	rootCmd.AddCommand(generateCmd)

	carapace.Gen(generateCmd).PositionalAnyCompletion(carapace.ActionFiles(".go"))
}
