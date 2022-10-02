package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var presetAllCmd = &cobra.Command{
	Use:   "preset-all",
	Short: "Enable/disable all unit files based on preset configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(presetAllCmd).Standalone()

	rootCmd.AddCommand(presetAllCmd)
}
