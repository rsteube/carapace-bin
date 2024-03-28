package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install or upgrade Flux",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()
	installCmd.Flags().String("arch", "", "cluster architecture, available options are: (amd64, arm, arm64)")
	installCmd.Flags().String("cluster-domain", "cluster.local", "internal cluster domain")
	installCmd.Flags().StringSlice("components", []string{"source-controller", "kustomize-controller", "helm-controller", "notification-controller"}, "list of components, accepts comma-separated values")
	installCmd.Flags().StringSlice("components-extra", []string{}, "list of components in addition to those supplied or defaulted, accepts comma-separated values")
	installCmd.Flags().Bool("dry-run", false, "only print the object that would be applied")
	installCmd.Flags().Bool("export", false, "write the install manifests to stdout and exit")
	installCmd.Flags().String("image-pull-secret", "", "Kubernetes secret name used for pulling the toolkit images from a private registry")
	installCmd.Flags().String("log-level", "", "log level, available options are: (debug, info, error)")
	installCmd.Flags().String("manifests", "", "path to the manifest directory")
	installCmd.Flags().Bool("network-policy", true, "deny ingress access to the toolkit controllers from other namespaces using network policies")
	installCmd.Flags().String("registry", "ghcr.io/fluxcd", "container registry where the toolkit images are published")
	installCmd.Flags().StringSlice("toleration-keys", []string{}, "list of toleration keys used to schedule the components pods onto nodes with matching taints")
	installCmd.Flags().StringP("version", "v", "", "toolkit version, when specified the manifests are downloaded from https://github.com/fluxcd/flux2/releases")
	installCmd.Flags().Bool("watch-all-namespaces", true, "watch for custom resources in all namespaces, if set to false it will only watch the namespace where the toolkit is installed")
	rootCmd.AddCommand(installCmd)
}
