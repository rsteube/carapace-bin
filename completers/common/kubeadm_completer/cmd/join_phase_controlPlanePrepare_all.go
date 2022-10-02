package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var join_phase_controlPlanePrepare_allCmd = &cobra.Command{
	Use:   "all",
	Short: "Prepare the machine for serving a control plane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(join_phase_controlPlanePrepare_allCmd).Standalone()
	join_phase_controlPlanePrepare_allCmd.Flags().String("apiserver-advertise-address", "", "If the node should host a new control plane instance, the IP address the API Server will advertise it's listening on. If not set the default network interface will be used.")
	join_phase_controlPlanePrepare_allCmd.Flags().Int32("apiserver-bind-port", 6443, "If the node should host a new control plane instance, the port for the API Server to bind to.")
	join_phase_controlPlanePrepare_allCmd.Flags().String("certificate-key", "", "Use this key to decrypt the certificate secrets uploaded by init.")
	join_phase_controlPlanePrepare_allCmd.Flags().String("config", "", "Path to kubeadm config file.")
	join_phase_controlPlanePrepare_allCmd.Flags().Bool("control-plane", false, "Create a new control plane instance on this node")
	join_phase_controlPlanePrepare_allCmd.Flags().String("discovery-file", "", "For file-based discovery, a file or URL from which to load cluster information.")
	join_phase_controlPlanePrepare_allCmd.Flags().String("discovery-token", "", "For token-based discovery, the token used to validate cluster information fetched from the API server.")
	join_phase_controlPlanePrepare_allCmd.Flags().StringSlice("discovery-token-ca-cert-hash", []string{}, "For token-based discovery, validate that the root CA public key matches this hash (format: \"<type>:<value>\").")
	join_phase_controlPlanePrepare_allCmd.Flags().Bool("discovery-token-unsafe-skip-ca-verification", false, "For token-based discovery, allow joining without --discovery-token-ca-cert-hash pinning.")
	join_phase_controlPlanePrepare_allCmd.Flags().String("node-name", "", "Specify the node name.")
	join_phase_controlPlanePrepare_allCmd.Flags().String("patches", "", "Path to a directory that contains files named \"target[suffix][+patchtype].extension\".")
	join_phase_controlPlanePrepare_allCmd.Flags().String("tls-bootstrap-token", "", "Specify the token used to temporarily authenticate with the Kubernetes Control Plane while joining the node.")
	join_phase_controlPlanePrepare_allCmd.Flags().String("token", "", "Use this token for both discovery-token and tls-bootstrap-token when those values are not provided.")
	join_phase_controlPlanePrepareCmd.AddCommand(join_phase_controlPlanePrepare_allCmd)

	carapace.Gen(join_phase_controlPlanePrepare_allCmd).FlagCompletion(carapace.ActionMap{
		"config":         carapace.ActionFiles(),
		"discovery-file": carapace.ActionFiles(),
		"patches":        carapace.ActionDirectories(),
	})
}
