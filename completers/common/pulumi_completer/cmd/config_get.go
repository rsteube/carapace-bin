package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/pulumi_completer/cmd/action"
	"github.com/spf13/cobra"
)

var config_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a single configuration value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_getCmd).Standalone()
	config_getCmd.Flags().BoolP("json", "j", false, "Emit output as JSON")
	config_getCmd.PersistentFlags().Bool("path", false, "The key contains a path to a property in a map or list to get")
	configCmd.AddCommand(config_getCmd)

	carapace.Gen(config_getCmd).PositionalCompletion(
		action.ActionConfigKeys(config_getCmd),
	)
}
