package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var server_installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install the Waypoint server to Kubernetes, Nomad, ECS, or Docker",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_installCmd).Standalone()

	server_installCmd.Flags().Bool("accept-tos", false, "Pass to accept the Terms of Service and Privacy Policy.")
	server_installCmd.Flags().String("context-create", "", "Create a context with connection information for this installation.")
	server_installCmd.Flags().Bool("context-set-default", false, "Set the newly installed server as the default CLI context.")
	server_installCmd.Flags().String("platform", "", "Platform to install the Waypoint server into.")

	addGlobalOptions(server_installCmd)
	addDockerOptions(server_installCmd)
	addEcsOptions(server_installCmd)
	addK8sOptions(server_installCmd)
	addNomadOptions(server_installCmd)

	serverCmd.AddCommand(server_installCmd)

	carapace.Gen(server_installCmd).FlagCompletion(carapace.ActionMap{
		"platform": carapace.ActionValues("kubernetes", "nomad", "ecs", "docker"),
	})
}
