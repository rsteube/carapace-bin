package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var create_kustomizationCmd = &cobra.Command{
	Use:   "kustomization",
	Short: "Create or update a Kustomization resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_kustomizationCmd).Standalone()
	create_kustomizationCmd.Flags().String("decryption-provider", "", "decryption provider, available options are: (sops)")
	create_kustomizationCmd.Flags().String("decryption-secret", "", "set the Kubernetes secret name that contains the OpenPGP private keys used for sops decryption")
	create_kustomizationCmd.Flags().StringSlice("depends-on", []string{}, "Kustomization that must be ready before this Kustomization can be applied, supported formats '<name>' and '<namespace>/<name>', also accepts comma-separated values")
	create_kustomizationCmd.Flags().StringSlice("health-check", []string{}, "workload to be included in the health assessment, in the format '<kind>/<name>.<namespace>'")
	create_kustomizationCmd.Flags().String("health-check-timeout", "", "timeout of health checking operations")
	create_kustomizationCmd.Flags().String("path", "./", "path to the directory containing a kustomization.yaml file")
	create_kustomizationCmd.Flags().Bool("prune", false, "enable garbage collection")
	create_kustomizationCmd.Flags().String("service-account", "", "the name of the service account to impersonate when reconciling this Kustomization")
	create_kustomizationCmd.Flags().String("source", "", "source that contains the Kubernetes manifests in the format '[<kind>/]<name>.<namespace>', where kind must be one of: (GitRepository, Bucket), if kind is not specified it defaults to GitRepository")
	create_kustomizationCmd.Flags().String("target-namespace", "", "overrides the namespace of all Kustomization objects reconciled by this Kustomization")
	create_kustomizationCmd.Flags().String("validation", "", "validate the manifests before applying them on the cluster, can be 'client' or 'server'")
	create_kustomizationCmd.Flags().Bool("wait", false, "enable health checking of all the applied resources")
	createCmd.AddCommand(create_kustomizationCmd)
}
