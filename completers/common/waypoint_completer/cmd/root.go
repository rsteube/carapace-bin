package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/waypoint_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/aws"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "waypoint",
	Short: "Easy application deployment for Kubernetes and Amazon ECS",
	Long:  "https://www.waypointproject.io/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	rootCmd.Flags().Bool("autocomplete-install", false, "show version")
	rootCmd.Flags().Bool("autocomplete-uninstall", false, "show version")
	rootCmd.Flags().Bool("version", false, "show version")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "show help")

	carapace.Gen(rootCmd).Standalone()
}

func addGlobalOptions(cmd *cobra.Command) {
	cmd.Flags().StringP("app", "a", "", "App to target.")
	cmd.Flags().Bool("plain", false, "Plain output")
	cmd.Flags().StringP("project", "p", "", "Project to target.")
	cmd.Flags().StringP("workspace", "w", "", "Workspace to operate in.")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"project":   action.ActionProjects(),
		"workspace": action.ActionWorkspaces(),
	})
}

func addConnectionOptions(cmd *cobra.Command) {
	cmd.Flags().String("server-addr", "", "Address for the server.")
	cmd.Flags().Bool("server-tls", false, "Connected to via TLS.")
	cmd.Flags().Bool("server-tls-skip-verify", false, "Skip verification of the TLS certificate advertised by the server.")

}

func addOperationOptions(cmd *cobra.Command) {
	cmd.Flags().String("label", "", "Labels to set for this operation.")
	cmd.Flags().Bool("remote", false, "True to use a remote runner to execute.")
	cmd.Flags().String("remote-source", "", "Override configurations for how remote runners source data.")
	cmd.Flags().String("var", "", "Variable value to set for this operation.")
	cmd.Flags().String("var-file", "", "HCL or JSON file containing variable values to set for this operation.")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"var-file": carapace.ActionFiles(".hcl", ".json"),
	})
}

func addFilterOptions(cmd *cobra.Command) {
	cmd.Flags().Bool("desc", false, "Sort the values in descending order.")
	cmd.Flags().String("limit", "", "How many values to show.")
	cmd.Flags().String("order-by", "", "Order the values by which field.")
	cmd.Flags().String("physical-state", "", "Show values in the given physical states.")
	cmd.Flags().String("state", "", "Filter values to have the given status.")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"order-by":       carapace.ActionValues("start-time", "complete-time"),
		"physical-state": carapace.ActionValues("any", "created", "destroyed", "pending"),
		"state":          carapace.ActionValues("error", "running", "success", "unknown"),
	})
}

func addOidcAuthMethodOptions(cmd *cobra.Command) {
	cmd.Flags().String("allowed-redirect-uri", "", "Allowed URI for auth redirection.")
	cmd.Flags().String("claim-mapping", "", "Mapping of a claim to a variable value for the access selector.")
	cmd.Flags().String("claim-scope", "", "The optional claims scope requested.")
	cmd.Flags().String("client-id", "", "The OAuth 2.0 Client Identifier.")
	cmd.Flags().String("client-secret", "", "The client secret corresponding with the client ID.")
	cmd.Flags().String("issuer", "", "Discovery URL of the OIDC provider.")
	cmd.Flags().String("issuer-ca-pem", "", "PEM-encoded certificates for connecting to the issuer.")
	cmd.Flags().String("list-claim-mapping", "", "Same as claim-mapping but for list values.")
	cmd.Flags().String("signing-algorithm", "", "The allowed signing algorithm.")
}

func addDockerOptions(cmd *cobra.Command) {
	cmd.Flags().String("docker-odr-image", "", "Docker image for the Waypoint On-Demand Runners")
	cmd.Flags().String("docker-server-image", "", "Docker image for the Waypoint server.")
}

func addEcsOptions(cmd *cobra.Command) {
	cmd.Flags().String("ecs-cluster", "", "Configures the Cluster to install into.")
	cmd.Flags().String("ecs-cpu", "", "Configures the requested CPU amount for the Waypoint server task in ECS.")
	cmd.Flags().String("ecs-execution-role-name", "", "Configures the IAM Execution role name to use.")
	cmd.Flags().String("ecs-mem", "", "Configures the requested memory amount for the Waypoint server task in ECS.")
	cmd.Flags().String("ecs-odr-cpu", "", "Configures the requested CPU amount for the Waypoint On-Demand runner in ECS.")
	cmd.Flags().String("ecs-odr-image", "", "Docker image for the Waypoint On-Demand Runners.")
	cmd.Flags().String("ecs-odr-mem", "", "Configures the requested memory amount for the Waypoint On-Demand runner in ECS.")
	cmd.Flags().String("ecs-region", "", "Configures which AWS region to install into.")
	cmd.Flags().String("ecs-server-image", "", "Docker image for the Waypoint server.")
	cmd.Flags().String("ecs-subnets", "", "Subnets to install server into.")
	cmd.Flags().String("ecs-task-role-name", "", "IAM Execution Role to assign to the on-demand runner.")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"ecs-region": aws.ActionRegions(),
	})
}

func addK8sOptions(cmd *cobra.Command) {
	cmd.Flags().Bool("k8s-advertise-internal", false, "Advertise the internal service address rather than the external.")
	cmd.Flags().String("k8s-annotate-service", "", "Annotations for the Service generated.")
	cmd.Flags().String("k8s-context", "", "The Kubernetes context to install the Waypoint server to.")
	cmd.Flags().String("k8s-cpu-request", "", "Configures the requested CPU amount for the Waypoint server in Kubernetes.")
	cmd.Flags().String("k8s-mem-request", "", "Configures the requested memory amount for the Waypoint server in Kubernetes.")
	cmd.Flags().String("k8s-namespace", "", "Namespace to install the Waypoint server into for Kubernetes.")
	cmd.Flags().String("k8s-odr-image", "", "Docker image for the Waypoint On-Demand Runners")
	cmd.Flags().Bool("k8s-openshift", false, "Enables installing the Waypoint server on Kubernetes on Red Hat OpenShift.")
	cmd.Flags().String("k8s-pull-policy", "", "Set the pull policy for the Waypoint server image.")
	cmd.Flags().String("k8s-pull-secret", "", "Secret to use to access the Waypoint server image on Kubernetes.")
	cmd.Flags().String("k8s-runner-service-account", "", "Service account to assign to the on-demand runner.")
	cmd.Flags().Bool("k8s-runner-service-account-init", false, "Create the service account if it does not exist.")
	cmd.Flags().String("k8s-secret-file", "", "Use the Kubernetes Secret in the given path to access the Waypoint server image.")
	cmd.Flags().String("k8s-server-image", "", "Docker image for the Waypoint server.")
	cmd.Flags().String("k8s-storage-request", "", "Configures the requested persistent volume size for the Waypoint server in Kubernetes.")
	cmd.Flags().String("k8s-storageclassname", "", "Name of the StorageClass required by the volume claim to install the Waypoint server image to.")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"k8s-secret-file": carapace.ActionFiles(),
	})
}

func addNomadOptions(cmd *cobra.Command) {
	cmd.Flags().String("nomad-annotate-service", "", "Annotations for the Service generated.")
	cmd.Flags().Bool("nomad-auth-soft-fail", false, "Don't fail the Nomad task on an auth failure obtaining server image container.")
	cmd.Flags().String("nomad-consul-datacenter", "", "The datacenter where Consul is located.")
	cmd.Flags().String("nomad-consul-domain", "", "The domain where Consul is located.")
	cmd.Flags().Bool("nomad-consul-service", false, "Create service for Waypoint UI and Server in Consul.")
	cmd.Flags().String("nomad-consul-service-backend-tags", "", "Tags for the Waypoint backend service generated in Consul.")
	cmd.Flags().String("nomad-consul-service-hostname", "", "If set, will use this hostname for Consul DNS rather than the default.")
	cmd.Flags().String("nomad-consul-service-ui-tags", "", "Tags for the Waypoint UI service generated in Consul.")
	cmd.Flags().String("nomad-csi-fs", "", "Nomad CSI volume mount option file system. The default is xfs.")
	cmd.Flags().String("nomad-csi-volume-capacity-max", "", "Nomad CSI volume capacity maximum, in bytes.")
	cmd.Flags().String("nomad-csi-volume-capacity-min", "", "Nomad CSI volume capacity minimum, in bytes.")
	cmd.Flags().String("nomad-csi-volume-provider", "", "Nomad CSI volume provider, required for volume type 'csi'.")
	cmd.Flags().String("nomad-dc", "", "Datacenters to install to for Nomad. The default is dc1.")
	cmd.Flags().String("nomad-host", "", "Hostname of the Nomad server to use, like for launching on-demand tasks.")
	cmd.Flags().String("nomad-host-volume", "", "Nomad host volume name, required for volume type 'host'.")
	cmd.Flags().String("nomad-namespace", "", "Namespace to install the Waypoint server into for Nomad.")
	cmd.Flags().String("nomad-odr-image", "", "Docker image for the on-demand runners.")
	cmd.Flags().Bool("nomad-policy-override", false, "Override the Nomad sentinel policy for enterprise Nomad.")
	cmd.Flags().String("nomad-region", "", "Region to install to for Nomad.")
	cmd.Flags().String("nomad-runner-cpu", "", "CPU required to run this task in MHz.")
	cmd.Flags().String("nomad-runner-memory", "", "MB of Memory to allocate to the runner job task.")
	cmd.Flags().String("nomad-server-cpu", "", "CPU required to run this task in MHz.")
	cmd.Flags().String("nomad-server-image", "", "Docker image for the Waypoint server.")
	cmd.Flags().String("nomad-server-memory", "", "MB of Memory to allocate to the Server job task.")
}
