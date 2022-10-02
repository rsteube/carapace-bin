package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var config_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a config variable",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_setCmd).Standalone()

	config_setCmd.Flags().String("label-scope", "", "Only set if the deployment or operation has a matching label set.")
	config_setCmd.Flags().Bool("runner", false, "Expose this configuration on runners.")
	config_setCmd.Flags().String("scope", "", "The scope for this configuration.")
	config_setCmd.Flags().String("workspace-scope", "", "Specify that the configuration is only available within a specific workspace.")

	addGlobalOptions(config_setCmd)

	configCmd.AddCommand(config_setCmd)
}
