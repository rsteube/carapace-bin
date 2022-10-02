package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var autoscaleCmd = &cobra.Command{
	Use:   "autoscale",
	Short: "Auto-scale a Deployment, ReplicaSet, or ReplicationController",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaleCmd).Standalone()

	autoscaleCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only app")
	autoscaleCmd.Flags().String("cpu-percent", "", "The target average CPU utilization (represented as a percent of requested CPU) over all the pods. If")
	autoscaleCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent,")
	autoscaleCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	autoscaleCmd.Flags().StringP("filename", "f", "", "Filename, directory, or URL to files identifying the resource to autoscale.")
	autoscaleCmd.Flags().String("generator", "", "The name of the API generator to use. Currently there is only 1 generator.")
	autoscaleCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	autoscaleCmd.Flags().String("max", "", "The upper limit for the number of pods that can be set by the autoscaler. Required.")
	autoscaleCmd.Flags().String("min", "", "The lower limit for the number of pods that can be set by the autoscaler. If it's not specified or n")
	autoscaleCmd.Flags().String("name", "", "The name for the newly created object. If not specified, the name of the input resource will be used")
	autoscaleCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|js")
	autoscaleCmd.Flags().Bool("record", false, "Record current kubectl command in the resource annotation. If set to false, do not record the comman")
	autoscaleCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related man")
	autoscaleCmd.Flags().Bool("save-config", false, "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotat")
	autoscaleCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The templa")
	rootCmd.AddCommand(autoscaleCmd)

	carapace.Gen(autoscaleCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":   action.ActionDryRunModes(),
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
		"output":    action.ActionOutputFormats(),
		"template":  carapace.ActionFiles(),
	})

	carapace.Gen(autoscaleCmd).PositionalCompletion(
		carapace.ActionValues("deployments", "replicasets", "replicationcontrollers"),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionResources("", c.Args[0])
		}),
	)
}
