package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubeadm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var joinCmd = &cobra.Command{
	Use:   "join",
	Short: "Run this on any machine you wish to join an existing cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(joinCmd).Standalone()
	joinCmd.Flags().String("apiserver-advertise-address", "", "If the node should host a new control plane instance, the IP address the API Server will advertise it's listening on. If not set the default network interface will be used.")
	joinCmd.Flags().Int32("apiserver-bind-port", 6443, "If the node should host a new control plane instance, the port for the API Server to bind to.")
	joinCmd.Flags().String("certificate-key", "", "Use this key to decrypt the certificate secrets uploaded by init.")
	joinCmd.Flags().String("config", "", "Path to kubeadm config file.")
	joinCmd.Flags().Bool("control-plane", false, "Create a new control plane instance on this node")
	joinCmd.Flags().String("cri-socket", "", "Path to the CRI socket to connect. If empty kubeadm will try to auto-detect this value; use this option only if you have more than one CRI installed or if you have non-standard CRI socket.")
	joinCmd.Flags().String("discovery-file", "", "For file-based discovery, a file or URL from which to load cluster information.")
	joinCmd.Flags().String("discovery-token", "", "For token-based discovery, the token used to validate cluster information fetched from the API server.")
	joinCmd.Flags().StringSlice("discovery-token-ca-cert-hash", []string{}, "For token-based discovery, validate that the root CA public key matches this hash (format: \"<type>:<value>\").")
	joinCmd.Flags().Bool("discovery-token-unsafe-skip-ca-verification", false, "For token-based discovery, allow joining without --discovery-token-ca-cert-hash pinning.")
	joinCmd.Flags().Bool("dry-run", false, "Don't apply any changes; just output what would be done.")
	joinCmd.Flags().StringSlice("ignore-preflight-errors", []string{}, "A list of checks whose errors will be shown as warnings. Example: 'IsPrivilegedUser,Swap'. Value 'all' ignores errors from all checks.")
	joinCmd.Flags().String("node-name", "", "Specify the node name.")
	joinCmd.Flags().String("patches", "", "Path to a directory that contains files named \"target[suffix][+patchtype].extension\".")
	joinCmd.Flags().StringSlice("skip-phases", []string{}, "List of phases to be skipped")
	joinCmd.Flags().String("tls-bootstrap-token", "", "Specify the token used to temporarily authenticate with the Kubernetes Control Plane while joining the node.")
	joinCmd.Flags().String("token", "", "Use this token for both discovery-token and tls-bootstrap-token when those values are not provided.")
	rootCmd.AddCommand(joinCmd)

	carapace.Gen(joinCmd).FlagCompletion(carapace.ActionMap{
		"config":         carapace.ActionFiles(),
		"cri-socket":     carapace.ActionFiles(),
		"discovery-file": carapace.ActionFiles(),
		"ignore-preflight-errors": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionChecks().Invoke(c).Filter(c.Parts).ToA()
		}),
		"patches":     carapace.ActionDirectories(),
		"skip-phases": action.ActionPhases(),
	})
}
