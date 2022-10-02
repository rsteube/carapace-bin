package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(testCmd).Standalone()
	testCmd.Flags().BoolP("help", "h", false, "display help for command")
	testCmd.Flags().String("schema", "", "validate against common schema")

	rootCmd.AddCommand(testCmd)

	carapace.Gen(testCmd).FlagCompletion(carapace.ActionMap{
		"schema": carapace.ActionFiles(),
	})

	carapace.Gen(testCmd).PositionalCompletion(
		carapace.ActionValues("init", "test", "export", "serve"),
	)
}
