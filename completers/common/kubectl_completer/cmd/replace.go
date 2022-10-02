package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var replaceCmd = &cobra.Command{
	Use:   "replace",
	Short: "Replace a resource by filename or stdin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(replaceCmd).Standalone()

	replaceCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only app")
	replaceCmd.Flags().String("cascade", "", "Must be \"background\", \"orphan\", or \"foreground\". Selects the deletion cascading strategy for the dep")
	replaceCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent,")
	replaceCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	replaceCmd.Flags().StringP("filename", "f", "", "to use to replace the resource.")
	replaceCmd.Flags().Bool("force", false, "If true, immediately remove resources from API and bypass graceful deletion. Note that immediate del")
	replaceCmd.Flags().String("grace-period", "", "Period of time in seconds given to the resource to terminate gracefully. Ignored if negative. Set to")
	replaceCmd.Flags().StringP("kustomize", "k", "", "Process a kustomization directory. This flag can't be used together with -f or -R.")
	replaceCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|js")
	replaceCmd.Flags().String("raw", "", "Raw URI to PUT to the server.  Uses the transport specified by the kubeconfig file.")
	replaceCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related man")
	replaceCmd.Flags().Bool("save-config", false, "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotat")
	replaceCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The templa")
	replaceCmd.Flags().String("timeout", "", "The length of time to wait before giving up on a delete, zero means determine a timeout from the siz")
	replaceCmd.Flags().Bool("validate", false, "If true, use a schema to validate the input before sending it")
	replaceCmd.Flags().Bool("wait", false, "If true, wait for resources to be gone before returning. This waits for finalizers.")
	rootCmd.AddCommand(replaceCmd)

	carapace.Gen(replaceCmd).FlagCompletion(carapace.ActionMap{
		"cascade":   carapace.ActionValues("background", "orphan", "foreground"),
		"dry-run":   action.ActionDryRunModes(),
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
		"output":    action.ActionOutputFormats(),
		"template":  carapace.ActionFiles(),
	})
}
