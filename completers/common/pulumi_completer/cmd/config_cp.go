package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/pulumi_completer/cmd/action"
	"github.com/spf13/cobra"
)

var config_cpCmd = &cobra.Command{
	Use:   "cp",
	Short: "Copy config to another stack",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_cpCmd).Standalone()
	config_cpCmd.PersistentFlags().StringP("dest", "d", "", "The name of the new stack to copy the config to")
	config_cpCmd.PersistentFlags().Bool("path", false, "The key contains a path to a property in a map or list to set")
	configCmd.AddCommand(config_cpCmd)

	carapace.Gen(config_cpCmd).FlagCompletion(carapace.ActionMap{
		"dest": action.ActionStacks(config_cpCmd, action.StackOpts{}),
	})

	carapace.Gen(config_cpCmd).PositionalCompletion(
		action.ActionConfigKeys(config_cpCmd),
	)
}
