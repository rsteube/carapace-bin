package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var init_phase_controlPlane_apiserverCmd = &cobra.Command{
	Use:   "apiserver",
	Short: "Generates the kube-apiserver static Pod manifest",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_controlPlane_apiserverCmd).Standalone()
	init_phase_controlPlane_apiserverCmd.Flags().String("apiserver-advertise-address", "", "The IP address the API Server will advertise it's listening on. If not set the default network interface will be used.")
	init_phase_controlPlane_apiserverCmd.Flags().Int32("apiserver-bind-port", 6443, "Port for the API Server to bind to.")
	init_phase_controlPlane_apiserverCmd.Flags().String("apiserver-extra-args", "", "A set of extra flags to pass to the API Server or override default ones in form of <flagname>=<value>")
	init_phase_controlPlane_apiserverCmd.Flags().String("cert-dir", "/etc/kubernetes/pki", "The path where to save and store the certificates.")
	init_phase_controlPlane_apiserverCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	init_phase_controlPlane_apiserverCmd.Flags().String("control-plane-endpoint", "", "Specify a stable IP address or DNS name for the control plane.")
	init_phase_controlPlane_apiserverCmd.Flags().Bool("dry-run", false, "Don't apply any changes; just output what would be done.")
	init_phase_controlPlane_apiserverCmd.Flags().String("feature-gates", "", "A set of key=value pairs that describe feature gates for various features.")
	init_phase_controlPlane_apiserverCmd.Flags().String("image-repository", "k8s.gcr.io", "Choose a container registry to pull control plane images from")
	init_phase_controlPlane_apiserverCmd.Flags().String("kubernetes-version", "stable-1", "Choose a specific Kubernetes version for the control plane.")
	init_phase_controlPlane_apiserverCmd.Flags().String("patches", "", "Path to a directory that contains files named \"target[suffix][+patchtype].extension\".")
	init_phase_controlPlane_apiserverCmd.Flags().String("service-cidr", "10.96.0.0/12", "Use alternative range of IP address for service VIPs.")
	init_phase_controlPlaneCmd.AddCommand(init_phase_controlPlane_apiserverCmd)

	carapace.Gen(init_phase_controlPlane_apiserverCmd).FlagCompletion(carapace.ActionMap{
		"cert-dir": carapace.ActionDirectories(),
		"config":   carapace.ActionFiles(),
		"patches":  carapace.ActionDirectories(),
	})
}
