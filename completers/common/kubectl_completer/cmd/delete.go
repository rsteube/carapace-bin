package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete resources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deleteCmd).Standalone()

	deleteCmd.Flags().Bool("all", false, "Delete all resources, including uninitialized ones, in the namespace of the specified resource types")
	deleteCmd.Flags().BoolP("all-namespaces", "A", false, "If present, list the requested object(s) across all namespaces. Namespace in current context is igno")
	deleteCmd.Flags().String("cascade", "", "Must be \"background\", \"orphan\", or \"foreground\". Selects the deletion cascading strategy for the dep")
	deleteCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent,")
	deleteCmd.Flags().String("field-selector", "", "Selector (field query) to filter on, supports '=', '==', and '!='.(e.g. --field-selector key1=value1")
	deleteCmd.Flags().StringP("filename", "f", "", "containing the resource to delete.")
	deleteCmd.Flags().Bool("force", false, "If true, immediately remove resources from API and bypass graceful deletion. Note that immediate del")
	deleteCmd.Flags().String("grace-period", "", "Period of time in seconds given to the resource to terminate gracefully. Ignored if negative. Set to")
	deleteCmd.Flags().Bool("ignore-not-found", false, "Treat \"resource not found\" as a successful delete. Defaults to \"true\" when --all is specified.")
	deleteCmd.Flags().StringP("kustomize", "k", "", "Process a kustomization directory. This flag can't be used together with -f or -R.")
	deleteCmd.Flags().Bool("now", false, "If true, resources are signaled for immediate shutdown (same as --grace-period=1).")
	deleteCmd.Flags().StringP("output", "o", "", "Output mode. Use \"-o name\" for shorter output (resource/name).")
	deleteCmd.Flags().String("raw", "", "Raw URI to DELETE to the server.  Uses the transport specified by the kubeconfig file.")
	deleteCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related man")
	deleteCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, not including uninitialized ones.")
	deleteCmd.Flags().String("timeout", "", "The length of time to wait before giving up on a delete, zero means determine a timeout from the siz")
	deleteCmd.Flags().Bool("wait", false, "If true, wait for resources to be gone before returning. This waits for finalizers.")
	rootCmd.AddCommand(deleteCmd)

	carapace.Gen(deleteCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":   action.ActionDryRunModes(),
		"kustomize": carapace.ActionDirectories(),
		"output":    action.ActionOutputFormats(),
	})

	carapace.Gen(deleteCmd).PositionalCompletion(
		carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionApiResources().Invoke(c).Filter(c.Parts).ToA()
		}),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionResources("", c.Args[0])
		}),
	)
}
