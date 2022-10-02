package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var config_initCmd = &cobra.Command{
	Use:   "init",
	Short: "Shows a prompt to set basic glab configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_initCmd).Standalone()
	configCmd.AddCommand(config_initCmd)
}
