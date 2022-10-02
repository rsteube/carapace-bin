package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Check whether the configuration is valid",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(validateCmd).Standalone()

	validateCmd.Flags().BoolS("json", "json", false, "Produce output in a machine-readable JSON format.")
	validateCmd.Flags().BoolS("no-color", "no-color", false, "If specified, output won't contain any color.")
	rootCmd.AddCommand(validateCmd)
}
