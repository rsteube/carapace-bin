package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/waypoint_completer/cmd/action"
	"github.com/spf13/cobra"
)

var server_upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrades Waypoint server in the current context to the latest version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_upgradeCmd).Standalone()

	server_upgradeCmd.Flags().Bool("auto-approve", false, "Confirm server upgrade.")
	server_upgradeCmd.Flags().String("context-name", "", "Waypoint server context to upgrade.")
	server_upgradeCmd.Flags().String("docker-odr-image", "", "Docker image for the Waypoint On-Demand Runners.")
	server_upgradeCmd.Flags().String("docker-server-image", "", "Docker image for the Waypoint server.")
	server_upgradeCmd.Flags().String("ecs-cluster", "", "Configures the Cluster to upgrade.")
	server_upgradeCmd.Flags().String("ecs-cpu", "", "Configures the requested CPU amount for the Waypoint server task in ECS.")
	server_upgradeCmd.Flags().String("ecs-mem", "", "Configures the requested memory amount for the Waypoint server task in ECS.")
	server_upgradeCmd.Flags().String("ecs-odr-cpu", "", "Configures the requested CPU amount for the Waypoint On-Demand runner in ECS.")
	server_upgradeCmd.Flags().String("ecs-odr-image", "", "Docker image for the Waypoint On-Demand Runners.")
	server_upgradeCmd.Flags().String("ecs-odr-mem", "", "Configures the requested memory amount for the Waypoint On-Demand runner in ECS.")
	server_upgradeCmd.Flags().String("ecs-region", "", "Configures which AWS region to install into.")
	server_upgradeCmd.Flags().String("ecs-server-image", "", "Docker image for the Waypoint server.")
	server_upgradeCmd.Flags().String("ecs-task-role-name", "", "IAM Execution Role to assign to the on-demand runner.")
	server_upgradeCmd.Flags().Bool("k8s-advertise-internal", false, "Advertise the internal service address rather than the external.")
	server_upgradeCmd.Flags().String("k8s-context", "", "The Kubernetes context to upgrade the Waypoint server to.")
	server_upgradeCmd.Flags().String("k8s-namespace", "", "Namespace to install the Waypoint server into for Kubernetes.")
	server_upgradeCmd.Flags().String("k8s-odr-image", "", "Docker image for the Waypoint On-Demand Runners")
	server_upgradeCmd.Flags().Bool("k8s-openshift", false, "Enables installing the Waypoint server on Kubernetes on Red Hat OpenShift.")
	server_upgradeCmd.Flags().String("k8s-runner-service-account", "", "Service account to assign to the on-demand runner.")
	server_upgradeCmd.Flags().Bool("k8s-runner-service-account-init", false, "Create the service account if it does not exist.")
	server_upgradeCmd.Flags().String("k8s-server-image", "", "Docker image for the Waypoint server.")
	server_upgradeCmd.Flags().String("nomad-annotate-service", "", "Annotations for the Service generated.")
	server_upgradeCmd.Flags().Bool("nomad-auth-soft-fail", false, "Don't fail the Nomad task on an auth failure obtaining server image container.")
	server_upgradeCmd.Flags().String("nomad-consul-datacenter", "", "The datacenter where Consul is located.")
	server_upgradeCmd.Flags().String("nomad-consul-domain", "", "The domain where Consul is located. The default is consul.")
	server_upgradeCmd.Flags().Bool("nomad-consul-service", false, "Create service for Waypoint UI and Server in Consul.")
	server_upgradeCmd.Flags().String("nomad-consul-service-backend-tags", "", "Tags for the Waypoint backend service generated in Consul.")
	server_upgradeCmd.Flags().String("nomad-consul-service-hostname", "", "If set, will use this hostname for Consul DNS rather than the default.")
	server_upgradeCmd.Flags().String("nomad-consul-service-ui-tags", "", "Tags for the Waypoint UI service generated in Consul.")
	server_upgradeCmd.Flags().String("nomad-dc", "", "Datacenters to install to for Nomad. The default is dc1.")
	server_upgradeCmd.Flags().String("nomad-host", "", "Hostname of the Nomad server to use.")
	server_upgradeCmd.Flags().String("nomad-host-volume", "", "Nomad host volume name.")
	server_upgradeCmd.Flags().String("nomad-namespace", "", "Namespace to install the Waypoint server into for Nomad.")
	server_upgradeCmd.Flags().String("nomad-odr-image", "", "Docker image for the on-demand runners.")
	server_upgradeCmd.Flags().Bool("nomad-policy-override", false, "Override the Nomad sentinel policy for enterprise Nomad.")
	server_upgradeCmd.Flags().String("nomad-region", "", "Region to install to for Nomad.")
	server_upgradeCmd.Flags().String("nomad-runner-cpu", "", "CPU required to run this task in MHz.")
	server_upgradeCmd.Flags().String("nomad-runner-memory", "", "MB of Memory to allocate to the runner job task.")
	server_upgradeCmd.Flags().String("nomad-server-cpu", "", "CPU required to run this task in MHz.")
	server_upgradeCmd.Flags().String("nomad-server-image", "", "Docker image for the Waypoint server.")
	server_upgradeCmd.Flags().String("nomad-server-memory", "", "MB of Memory to allocate to the server job task.")
	server_upgradeCmd.Flags().String("platform", "", "Platform to upgrade the Waypoint server from.")
	server_upgradeCmd.Flags().Bool("snapshot", false, "Enable or disable taking a snapshot of Waypoint server prior to upgrades.")
	server_upgradeCmd.Flags().String("snapshot-name", "", "Filename to write the snapshot to.")

	addGlobalOptions(server_upgradeCmd)

	serverCmd.AddCommand(server_upgradeCmd)

	carapace.Gen(server_upgradeCmd).FlagCompletion(carapace.ActionMap{
		"context-name": action.ActionContexts(),
	})
}
