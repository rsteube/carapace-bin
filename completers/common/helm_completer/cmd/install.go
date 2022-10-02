package cmd

import (
	"strings"
	"time"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/helm_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/util"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "install a chart",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	installCmd.Flags().Bool("atomic", false, "if set, the installation process deletes the installation on failure. The --wait flag will be set automatically if --atomic is used")
	installCmd.Flags().String("ca-file", "", "verify certificates of HTTPS-enabled servers using this CA bundle")
	installCmd.Flags().String("cert-file", "", "identify HTTPS client using this SSL certificate file")
	installCmd.Flags().Bool("create-namespace", false, "create the release namespace if not present")
	installCmd.Flags().Bool("dependency-update", false, "update dependencies if they are missing before installing the chart")
	installCmd.Flags().String("description", "", "add a custom description")
	installCmd.Flags().Bool("devel", false, "use development versions, too. Equivalent to version '>0.0.0-0'. If --version is set, this is ignored")
	installCmd.Flags().Bool("disable-openapi-validation", false, "if set, the installation process will not validate rendered templates against the Kubernetes OpenAPI Schema")
	installCmd.Flags().Bool("dry-run", false, "simulate an install")
	installCmd.Flags().BoolP("generate-name", "g", false, "generate the name (and omit the NAME parameter)")
	installCmd.Flags().Bool("insecure-skip-tls-verify", false, "skip tls certificate checks for the chart download")
	installCmd.Flags().String("key-file", "", "identify HTTPS client using this SSL key file")
	installCmd.Flags().String("keyring", "/home/user/.gnupg/pubring.gpg", "location of public keys used for verification")
	installCmd.Flags().String("name-template", "", "specify template used to name the release")
	installCmd.Flags().Bool("no-hooks", false, "prevent hooks from running during install")
	installCmd.Flags().StringP("output", "o", "table", "prints the output in the specified format. Allowed values: table, json, yaml")
	installCmd.Flags().Bool("pass-credentials", false, "pass credentials to all domains")
	installCmd.Flags().String("password", "", "chart repository password where to locate the requested chart")
	installCmd.Flags().String("post-renderer", "", "the path to an executable to be used for post rendering. If it exists in $PATH, the binary will be used, otherwise it will try to look for the executable at the given path")
	installCmd.Flags().Bool("render-subchart-notes", false, "if set, render subchart notes along with the parent")
	installCmd.Flags().Bool("replace", false, "re-use the given name, only if that name is a deleted release which remains in the history. This is unsafe in production")
	installCmd.Flags().String("repo", "", "chart repository url where to locate the requested chart")
	installCmd.Flags().StringArray("set", []string{}, "set values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)")
	installCmd.Flags().StringArray("set-file", []string{}, "set values from respective files specified via the command line (can specify multiple or separate values with commas: key1=path1,key2=path2)")
	installCmd.Flags().StringArray("set-string", []string{}, "set STRING values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)")
	installCmd.Flags().Bool("skip-crds", false, "if set, no CRDs will be installed. By default, CRDs are installed if not already present")
	installCmd.Flags().Duration("timeout", 5*time.Minute, "time to wait for any individual Kubernetes operation (like Jobs for hooks)")
	installCmd.Flags().String("username", "", "chart repository username where to locate the requested chart")
	installCmd.Flags().StringSliceP("values", "f", []string{}, "specify values in a YAML file or a URL (can specify multiple)")
	installCmd.Flags().Bool("verify", false, "verify the package before using it")
	installCmd.Flags().String("version", "", "specify a version constraint for the chart version to use. This constraint can be a specific tag (e.g. 1.1.1) or it may reference a valid range (e.g. ^2.0.0). If this is not specified, the latest version is used")
	installCmd.Flags().Bool("wait", false, "if set, will wait until all Pods, PVCs, Services, and minimum number of Pods of a Deployment, StatefulSet, or ReplicaSet are in a ready state before marking the release as successful. It will wait for as long as --timeout")
	installCmd.Flags().Bool("wait-for-jobs", false, "if set and --wait enabled, will wait until all Jobs have been completed before marking the release as successful. It will wait for as long as --timeout")
	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).FlagCompletion(carapace.ActionMap{
		"ca-file":   carapace.ActionFiles(),
		"cert-file": carapace.ActionFiles(),
		"key-file":  carapace.ActionFiles(),
		"keyring":   carapace.ActionFiles(),
		// TODO name-template
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

	carapace.Gen(installCmd).PositionalCompletion(
		carapace.ActionValues(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if util.HasPathPrefix(c.CallbackValue) {
				return carapace.ActionDirectories()
			}
			return action.ActionRepositoryCharts()
		}),
	)
}
