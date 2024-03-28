package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var config_initCmd = &cobra.Command{
	Use:     "init",
	Short:   "",
	Aliases: []string{"create"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_initCmd).Standalone()

	config_initCmd.Flags().BoolP("force", "f", false, "Force overwrite of target file")
	config_initCmd.Flags().StringP("output", "o", "", "Write a default k3d config")
	configCmd.AddCommand(config_initCmd)
}
