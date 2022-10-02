package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var apps_listCmd = &cobra.Command{
	Use:   "list",
	Short: "list verbs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apps_listCmd).Standalone()

	appsCmd.AddCommand(apps_listCmd)
}
