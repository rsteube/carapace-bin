package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var config_editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit the configuration file in an editor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_editCmd).Standalone()

	config_editCmd.Flags().String("editor", "", "Editor to use to edit the file")
	configCmd.AddCommand(config_editCmd)

	carapace.Gen(config_editCmd).FlagCompletion(carapace.ActionMap{
		"editor": carapace.ActionFiles(),
	})
}
