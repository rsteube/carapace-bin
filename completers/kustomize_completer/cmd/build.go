package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build a kustomization target from a directory or URL.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCmd).Standalone()
	buildCmd.Flags().Bool("as-current-user", false, "use the uid and gid of the command executor to run the function in the container")
	buildCmd.Flags().Bool("enable-alpha-plugins", false, "enable kustomize plugins")
	buildCmd.Flags().Bool("enable-exec", false, "enable support for exec functions (raw executables); do not use for untrusted configs! (Alpha)")
	buildCmd.Flags().Bool("enable-helm", false, "Enable use of the Helm chart inflator generator.")
	buildCmd.Flags().Bool("enable-managedby-label", false, "enable adding app.kubernetes.io/managed-by")
	buildCmd.Flags().Bool("enable-star", false, "enable support for starlark functions. (Alpha)")
	buildCmd.Flags().StringArrayP("env", "e", []string{}, "a list of environment variables to be used by functions")
	buildCmd.Flags().String("helm-command", "helm", "helm command (path to executable)")
	buildCmd.Flags().String("load-restrictor", "LoadRestrictionsRootOnly", "if set to 'LoadRestrictionsNone', local kustomizations may load files from outside their root. This does, however, break the relocatability of the kustomization.")
	buildCmd.Flags().StringArray("mount", []string{}, "a list of storage options read from the filesystem")
	buildCmd.Flags().Bool("network", false, "enable network access for functions that declare it")
	buildCmd.Flags().String("network-name", "bridge", "the docker network to run the container in")
	buildCmd.Flags().StringP("output", "o", "", "If specified, write output to this path.")
	buildCmd.Flags().String("reorder", "legacy", "Reorder the resources just before output. Use 'legacy' to apply a legacy reordering (Namespaces first, Webhooks last, etc). Use 'none' to suppress a final reordering.")
	rootCmd.AddCommand(buildCmd)

	carapace.Gen(buildCmd).FlagCompletion(carapace.ActionMap{
		"helm-command":    carapace.ActionFiles(),
		"load-restrictor": carapace.ActionValues("LoadRestrictionsNone", "LoadRestrictionsRootOnly"),
		"network-name":    docker.ActionNetworks(),
		"output":          carapace.ActionFiles(),
		"reorder": carapace.ActionValues(
			"legacy", "Namespaces first, Webhooks last, etc",
			"none", "suppress a final reordering",
		),
	})
}
