package cmd

import (
	"strings"
	"time"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/helm_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/util"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "upgrade a release",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	upgradeCmd.Flags().Bool("atomic", false, "if set, upgrade process rolls back changes made in case of failed upgrade. The --wait flag will be set automatically if --atomic is used")
	upgradeCmd.Flags().String("ca-file", "", "verify certificates of HTTPS-enabled servers using this CA bundle")
	upgradeCmd.Flags().String("cert-file", "", "identify HTTPS client using this SSL certificate file")
	upgradeCmd.Flags().Bool("cleanup-on-fail", false, "allow deletion of new resources created in this upgrade when upgrade fails")
	upgradeCmd.Flags().Bool("create-namespace", false, "if --install is set, create the release namespace if not present")
	upgradeCmd.Flags().Bool("dependency-update", false, "update dependencies if they are missing before installing the chart")
	upgradeCmd.Flags().String("description", "", "add a custom description")
	upgradeCmd.Flags().Bool("devel", false, "use development versions, too. Equivalent to version '>0.0.0-0'. If --version is set, this is ignored")
	upgradeCmd.Flags().Bool("disable-openapi-validation", false, "if set, the upgrade process will not validate rendered templates against the Kubernetes OpenAPI Schema")
	upgradeCmd.Flags().Bool("dry-run", false, "simulate an upgrade")
	upgradeCmd.Flags().Bool("force", false, "force resource updates through a replacement strategy")
	upgradeCmd.Flags().Int("history-max", 10, "limit the maximum number of revisions saved per release. Use 0 for no limit")
	upgradeCmd.Flags().Bool("insecure-skip-tls-verify", false, "skip tls certificate checks for the chart download")
	upgradeCmd.Flags().BoolP("install", "i", false, "if a release by this name doesn't already exist, run an install")
	upgradeCmd.Flags().String("key-file", "", "identify HTTPS client using this SSL key file")
	upgradeCmd.Flags().String("keyring", "/home/user/.gnupg/pubring.gpg", "location of public keys used for verification")
	upgradeCmd.Flags().Bool("no-hooks", false, "disable pre/post upgrade hooks")
	upgradeCmd.Flags().StringP("output", "o", "table", "prints the output in the specified format. Allowed values: table, json, yaml")
	upgradeCmd.Flags().Bool("pass-credentials", false, "pass credentials to all domains")
	upgradeCmd.Flags().String("password", "", "chart repository password where to locate the requested chart")
	upgradeCmd.Flags().String("post-renderer", "", "the path to an executable to be used for post rendering. If it exists in $PATH, the binary will be used, otherwise it will try to look for the executable at the given path")
	upgradeCmd.Flags().Bool("recreate-pods", false, "performs pods restart for the resource if applicable")
	upgradeCmd.Flags().Bool("render-subchart-notes", false, "if set, render subchart notes along with the parent")
	upgradeCmd.Flags().String("repo", "", "chart repository url where to locate the requested chart")
	upgradeCmd.Flags().Bool("reset-values", false, "when upgrading, reset the values to the ones built into the chart")
	upgradeCmd.Flags().Bool("reuse-values", false, "when upgrading, reuse the last release's values and merge in any overrides from the command line via --set and -f. If '--reset-values' is specified, this is ignored")
	upgradeCmd.Flags().StringArray("set", []string{}, "set values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)")
	upgradeCmd.Flags().StringArray("set-file", []string{}, "set values from respective files specified via the command line (can specify multiple or separate values with commas: key1=path1,key2=path2)")
	upgradeCmd.Flags().StringArray("set-string", []string{}, "set STRING values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)")
	upgradeCmd.Flags().Bool("skip-crds", false, "if set, no CRDs will be installed when an upgrade is performed with install flag enabled. By default, CRDs are installed if not already present, when an upgrade is performed with install flag enabled")
	upgradeCmd.Flags().Duration("timeout", 5*time.Minute, "time to wait for any individual Kubernetes operation (like Jobs for hooks)")
	upgradeCmd.Flags().String("username", "", "chart repository username where to locate the requested chart")
	upgradeCmd.Flags().StringSliceP("values", "f", []string{}, "specify values in a YAML file or a URL (can specify multiple)")
	upgradeCmd.Flags().Bool("verify", false, "verify the package before using it")
	upgradeCmd.Flags().String("version", "", "specify a version constraint for the chart version to use. This constraint can be a specific tag (e.g. 1.1.1) or it may reference a valid range (e.g. ^2.0.0). If this is not specified, the latest version is used")
	upgradeCmd.Flags().Bool("wait", false, "if set, will wait until all Pods, PVCs, Services, and minimum number of Pods of a Deployment, StatefulSet, or ReplicaSet are in a ready state before marking the release as successful. It will wait for as long as --timeout")
	upgradeCmd.Flags().Bool("wait-for-jobs", false, "if set and --wait enabled, will wait until all Jobs have been completed before marking the release as successful. It will wait for as long as --timeout")
	rootCmd.AddCommand(upgradeCmd)

	carapace.Gen(upgradeCmd).FlagCompletion(carapace.ActionMap{
		"ca-file":       carapace.ActionFiles(),
		"cert-file":     carapace.ActionFiles(),
		"key-file":      carapace.ActionFiles(),
		"keyring":       carapace.ActionFiles(),
		"output":        carapace.ActionValues("table", "json", "yaml"),
		"post-renderer": carapace.ActionFiles(),
		"set-file": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 1:
					return carapace.ActionFiles()
				default:
					return carapace.ActionValues()
				}
			})
		}),
		"version": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 1 {
				if splitted := strings.Split(c.Args[1], "/"); len(splitted) == 2 {
					return action.ActionChartVersions(splitted[0], splitted[1])
				}
			}
			return carapace.ActionValues()
		}),
	})

	carapace.Gen(upgradeCmd).PositionalCompletion(
		action.ActionReleases(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if util.HasPathPrefix(c.CallbackValue) {
				return carapace.ActionFiles()
			}
			return action.ActionRepositoryCharts()
		}),
	)
}
