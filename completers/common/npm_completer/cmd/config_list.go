package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var config_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show all the config settings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_listCmd).Standalone()
	config_listCmd.Flags().Bool("json", false, "output as json")

	configCmd.AddCommand(config_listCmd)
}
