package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "Display debug information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugCmd).Standalone()

	debugCmd.Flags().String("abi", "", "Only use wheels compatible with Python abi")
	debugCmd.Flags().String("implementation", "", "Only use wheels compatible with Python implementation")
	debugCmd.Flags().String("platform", "", "Only use wheels compatible with <platform>")
	debugCmd.Flags().String("python-version", "", "The Python interpreter version to use for wheel")
	rootCmd.AddCommand(debugCmd)
}
