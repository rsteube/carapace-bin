package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var scaleCmd = &cobra.Command{
	Use:   "scale",
	Short: "Set a new size for a Deployment, ReplicaSet or Replication Controller",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scaleCmd).Standalone()

	scaleCmd.Flags().Bool("all", false, "Select all resources in the namespace of the specified resource types")
	scaleCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only app")
	scaleCmd.Flags().String("current-replicas", "", "Precondition for current size. Requires that the current size of the resource match this value in or")
	scaleCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent,")
	scaleCmd.Flags().StringP("filename", "f", "", "Filename, directory, or URL to files identifying the resource to set a new size")
	scaleCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	scaleCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|js")
	scaleCmd.Flags().Bool("record", false, "Record current kubectl command in the resource annotation. If set to false, do not record the comman")
	scaleCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related man")
	scaleCmd.Flags().String("replicas", "", "The new desired number of replicas. Required.")
	scaleCmd.Flags().String("resource-version", "", "Precondition for resource version. Requires that the current resource version match this value in or")
	scaleCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)")
	scaleCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The templa")
	scaleCmd.Flags().String("timeout", "", "The length of time to wait before giving up on a scale operation, zero means don't wait. Any other v")
	rootCmd.AddCommand(scaleCmd)

	carapace.Gen(scaleCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":   action.ActionDryRunModes(),
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
		"output":    action.ActionOutputFormats(),
		"template":  carapace.ActionFiles(),
	})

	carapace.Gen(scaleCmd).PositionalAnyCompletion(
		carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValues("deployments/", "replicasets/", "replicationcontrollers/")
			case 1:
				return action.ActionResources("", c.Parts[0])
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
