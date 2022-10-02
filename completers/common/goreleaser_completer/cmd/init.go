package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Generates a .goreleaser.yml file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().StringP("config", "f", "", "Load configuration from file (default \".goreleaser.yml\")")
	initCmd.Flags().Bool("debug", false, "Enable debug mode")
	initCmd.Flags().BoolP("help", "h", false, "help for init")
	rootCmd.AddCommand(initCmd)

	carapace.Gen(initCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
	})
}
