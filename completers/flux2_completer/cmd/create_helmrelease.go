package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var create_helmreleaseCmd = &cobra.Command{
	Use:   "helmrelease",
	Short: "Create or update a HelmRelease resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_helmreleaseCmd).Standalone()
	create_helmreleaseCmd.Flags().String("chart", "", "Helm chart name or path")
	create_helmreleaseCmd.Flags().String("chart-version", "", "Helm chart version, accepts a semver range (ignored for charts from GitRepository sources)")
	create_helmreleaseCmd.Flags().String("crds", "", "upgrade CRDs policy, available options are: (Skip, Create, CreateReplace)")
	create_helmreleaseCmd.Flags().Bool("create-target-namespace", false, "create the target namespace if it does not exist")
	create_helmreleaseCmd.Flags().StringSlice("depends-on", []string{}, "HelmReleases that must be ready before this release can be installed, supported formats '<name>' and '<namespace>/<name>'")
	create_helmreleaseCmd.Flags().String("release-name", "", "name used for the Helm release, defaults to a composition of '[<target-namespace>-]<HelmRelease-name>'")
	create_helmreleaseCmd.Flags().String("service-account", "", "the name of the service account to impersonate when reconciling this HelmRelease")
	create_helmreleaseCmd.Flags().String("source", "", "source that contains the chart in the format '<kind>/<name>.<namespace>', where kind must be one of: (HelmRepository, GitRepository, Bucket)")
	create_helmreleaseCmd.Flags().String("target-namespace", "", "namespace to install this release, defaults to the HelmRelease namespace")
	create_helmreleaseCmd.Flags().StringSlice("values", []string{}, "local path to values.yaml files, also accepts comma-separated values")
	create_helmreleaseCmd.Flags().String("values-from", "", "Kubernetes object reference that contains the values.yaml data key in the format '<kind>/<name>', where kind must be one of: (Secret, ConfigMap)")
	createCmd.AddCommand(create_helmreleaseCmd)
}
