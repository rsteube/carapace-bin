package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install the Waypoint server to Kubernetes, Nomad, ECS, or Docker",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().Bool("accept-tos", false, "Pass to accept the Terms of Service and Privacy Policy.")
	installCmd.Flags().String("context-create", "", "Create a context with connection information for this installation.")
	installCmd.Flags().Bool("context-set-default", false, "Set the newly installed server as the default CLI context.")
	installCmd.Flags().String("platform", "", "Platform to install the Waypoint server into.")

	addGlobalOptions(installCmd)
	addDockerOptions(installCmd)
	addEcsOptions(installCmd)
	addK8sOptions(installCmd)
	addNomadOptions(installCmd)

	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).FlagCompletion(carapace.ActionMap{
		"platform": carapace.ActionValues("kubernetes", "nomad", "ecs", "docker"),
	})
}
