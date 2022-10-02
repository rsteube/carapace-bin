package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Checks if configuration is valid",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(checkCmd).Standalone()

	checkCmd.Flags().StringP("config", "f", "", "Configuration file to check")
	checkCmd.Flags().Bool("debug", false, "Enable debug mode")
	checkCmd.Flags().BoolP("help", "h", false, "help for check")
	rootCmd.AddCommand(checkCmd)

	carapace.Gen(checkCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
	})
}
