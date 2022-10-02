package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var config_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the active configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_listCmd).Standalone()

	configCmd.AddCommand(config_listCmd)
}
