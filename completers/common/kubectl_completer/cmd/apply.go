package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply a configuration to a resource by filename or stdin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applyCmd).Standalone()

	applyCmd.Flags().Bool("all", false, "Select all resources in the namespace of the specified resource types.")
	applyCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only app")
	applyCmd.Flags().String("cascade", "", "Must be \"background\", \"orphan\", or \"foreground\". Selects the deletion cascading strategy for the dep")
	applyCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent,")
	applyCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	applyCmd.Flags().StringP("filename", "f", "", "that contains the configuration to apply")
	applyCmd.Flags().Bool("force", false, "If true, immediately remove resources from API and bypass graceful deletion. Note that immediate del")
	applyCmd.Flags().Bool("force-conflicts", false, "If true, server-side apply will force the changes against conflicts.")
	applyCmd.Flags().String("grace-period", "", "Period of time in seconds given to the resource to terminate gracefully. Ignored if negative. Set to")
	applyCmd.Flags().StringP("kustomize", "k", "", "Process a kustomization directory. This flag can't be used together with -f or -R.")
	applyCmd.Flags().Bool("openapi-patch", false, "If true, use openapi to calculate diff when the openapi presents and the resource can be found in th")
	applyCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|js")
	applyCmd.Flags().Bool("overwrite", false, "Automatically resolve conflicts between the modified and live configuration by using values from the")
	applyCmd.Flags().Bool("prune", false, "Automatically delete resource objects, including the uninitialized ones, that do not appear in the c")
	applyCmd.Flags().String("prune-whitelist", "", "Overwrite the default whitelist with <group/version/kind> for --prune")
	applyCmd.Flags().Bool("record", false, "Record current kubectl command in the resource annotation. If set to false, do not record the comman")
	applyCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related man")
	applyCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)")
	applyCmd.Flags().Bool("server-side", false, "If true, apply runs in the server instead of the client.")
	applyCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The templa")
	applyCmd.Flags().String("timeout", "", "The length of time to wait before giving up on a delete, zero means determine a timeout from the siz")
	applyCmd.Flags().Bool("validate", false, "If true, use a schema to validate the input before sending it")
	applyCmd.Flags().Bool("wait", false, "If true, wait for resources to be gone before returning. This waits for finalizers.")
	rootCmd.AddCommand(applyCmd)

	carapace.Gen(applyCmd).FlagCompletion(carapace.ActionMap{
		"cascade":   carapace.ActionValues("background", "orphan", "foreground"),
		"dry-run":   action.ActionDryRunModes(),
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
		"output":    action.ActionOutputFormats(),
		"template":  carapace.ActionFiles(),
	})
}
