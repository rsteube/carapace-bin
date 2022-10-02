package cmd

import (
	"time"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "locally render templates",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	templateCmd.Flags().StringArrayP("api-versions", "a", []string{}, "Kubernetes api versions used for Capabilities.APIVersions")
	templateCmd.Flags().Bool("atomic", false, "if set, the installation process deletes the installation on failure. The --wait flag will be set automatically if --atomic is used")
	templateCmd.Flags().String("ca-file", "", "verify certificates of HTTPS-enabled servers using this CA bundle")
	templateCmd.Flags().String("cert-file", "", "identify HTTPS client using this SSL certificate file")
	templateCmd.Flags().Bool("create-namespace", false, "create the release namespace if not present")
	templateCmd.Flags().Bool("dependency-update", false, "update dependencies if they are missing before installing the chart")
	templateCmd.Flags().String("description", "", "add a custom description")
	templateCmd.Flags().Bool("devel", false, "use development versions, too. Equivalent to version '>0.0.0-0'. If --version is set, this is ignored")
	templateCmd.Flags().Bool("disable-openapi-validation", false, "if set, the installation process will not validate rendered templates against the Kubernetes OpenAPI Schema")
	templateCmd.Flags().Bool("dry-run", false, "simulate an install")
	templateCmd.Flags().BoolP("generate-name", "g", false, "generate the name (and omit the NAME parameter)")
	templateCmd.Flags().Bool("include-crds", false, "include CRDs in the templated output")
	templateCmd.Flags().Bool("insecure-skip-tls-verify", false, "skip tls certificate checks for the chart download")
	templateCmd.Flags().Bool("is-upgrade", false, "set .Release.IsUpgrade instead of .Release.IsInstall")
	templateCmd.Flags().String("key-file", "", "identify HTTPS client using this SSL key file")
	templateCmd.Flags().String("keyring", "/home/user/.gnupg/pubring.gpg", "location of public keys used for verification")
	templateCmd.Flags().String("kube-version", "", "Kubernetes version used for Capabilities.KubeVersion")
	templateCmd.Flags().String("name-template", "", "specify template used to name the release")
	templateCmd.Flags().Bool("no-hooks", false, "prevent hooks from running during install")
	templateCmd.Flags().String("output-dir", "", "writes the executed templates to files in output-dir instead of stdout")
	templateCmd.Flags().Bool("pass-credentials", false, "pass credentials to all domains")
	templateCmd.Flags().String("password", "", "chart repository password where to locate the requested chart")
	templateCmd.Flags().String("post-renderer", "", "the path to an executable to be used for post rendering. If it exists in $PATH, the binary will be used, otherwise it will try to look for the executable at the given path")
	templateCmd.Flags().Bool("release-name", false, "use release name in the output-dir path.")
	templateCmd.Flags().Bool("render-subchart-notes", false, "if set, render subchart notes along with the parent")
	templateCmd.Flags().Bool("replace", false, "re-use the given name, only if that name is a deleted release which remains in the history. This is unsafe in production")
	templateCmd.Flags().String("repo", "", "chart repository url where to locate the requested chart")
	templateCmd.Flags().StringArray("set", []string{}, "set values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)")
	templateCmd.Flags().StringArray("set-file", []string{}, "set values from respective files specified via the command line (can specify multiple or separate values with commas: key1=path1,key2=path2)")
	templateCmd.Flags().StringArray("set-string", []string{}, "set STRING values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)")
	templateCmd.Flags().StringArrayP("show-only", "s", []string{}, "only show manifests rendered from the given templates")
	templateCmd.Flags().Bool("skip-crds", false, "if set, no CRDs will be installed. By default, CRDs are installed if not already present")
	templateCmd.Flags().Bool("skip-tests", false, "skip tests from templated output")
	templateCmd.Flags().Duration("timeout", 5*time.Minute, "time to wait for any individual Kubernetes operation (like Jobs for hooks)")
	templateCmd.Flags().String("username", "", "chart repository username where to locate the requested chart")
	templateCmd.Flags().Bool("validate", false, "validate your manifests against the Kubernetes cluster you are currently pointing at. This is the same validation performed on an install")
	templateCmd.Flags().StringSliceP("values", "f", []string{}, "specify values in a YAML file or a URL (can specify multiple)")
	templateCmd.Flags().Bool("verify", false, "verify the package before using it")
	templateCmd.Flags().String("version", "", "specify a version constraint for the chart version to use. This constraint can be a specific tag (e.g. 1.1.1) or it may reference a valid range (e.g. ^2.0.0). If this is not specified, the latest version is used")
	templateCmd.Flags().Bool("wait", false, "if set, will wait until all Pods, PVCs, Services, and minimum number of Pods of a Deployment, StatefulSet, or ReplicaSet are in a ready state before marking the release as successful. It will wait for as long as --timeout")
	templateCmd.Flags().Bool("wait-for-jobs", false, "if set and --wait enabled, will wait until all Jobs have been completed before marking the release as successful. It will wait for as long as --timeout")
	rootCmd.AddCommand(templateCmd)

	carapace.Gen(templateCmd).FlagCompletion(carapace.ActionMap{
		"ca-file":       carapace.ActionFiles(),
		"cert-file":     carapace.ActionFiles(),
		"key-file":      carapace.ActionFiles(),
		"keyring":       carapace.ActionFiles(),
		"output-dir":    carapace.ActionDirectories(),
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
	})

	// TODO positional completion
}
