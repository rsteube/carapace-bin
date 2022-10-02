package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var runner_profile_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Create or update a runner profile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runner_profile_setCmd).Standalone()

	runner_profile_setCmd.Flags().Bool("default", false, "Indicates that this on-demand runner should be used as default.")
	runner_profile_setCmd.Flags().String("env-vars", "", "Environment variable to expose to the on-demand runner.")
	runner_profile_setCmd.Flags().String("name", "", "The name of an existing runner profile to update.")
	runner_profile_setCmd.Flags().String("oci-url", "", "The url for the OCI image to launch for the on-demand runner.")
	runner_profile_setCmd.Flags().String("plugin-config", "", "Path to an hcl file that contains the configuration for the plugin.")
	runner_profile_setCmd.Flags().String("plugin-type", "", "The type of the plugin to launch for the on-demand runner.")

	addGlobalOptions(runner_profile_setCmd)

	runner_profileCmd.AddCommand(runner_profile_setCmd)

	carapace.Gen(runner_profile_setCmd).FlagCompletion(carapace.ActionMap{
		"plugin-config": carapace.ActionFiles(".hcl"),
	})
}
