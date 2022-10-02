package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var settings_listCmd = &cobra.Command{
	Use:   "list",
	Short: "list verbs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(settings_listCmd).Standalone()

	settingsCmd.AddCommand(settings_listCmd)
}
