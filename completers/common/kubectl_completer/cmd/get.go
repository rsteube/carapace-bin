package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "display one or many resources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getCmd).Standalone()

	getCmd.Flags().BoolP("all-namespaces", "A", false, "If present, list the requested object(s) across all namespaces. Namespace in current context is igno")
	getCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only app")
	getCmd.Flags().String("chunk-size", "", "Return large lists in chunks rather than all at once. Pass 0 to disable. This flag is beta and may c")
	getCmd.Flags().String("field-selector", "", "Selector (field query) to filter on, supports '=', '==', and '!='.(e.g. --field-selector key1=value1")
	getCmd.Flags().StringP("filename", "f", "", "Filename, directory, or URL to files identifying the resource to get from a server.")
	getCmd.Flags().Bool("ignore-not-found", false, "If the requested object does not exist the command will return exit code 0.")
	getCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	getCmd.Flags().StringP("label-columns", "L", "", "Accepts a comma separated list of labels that are going to be presented as columns. Names are case-s")
	getCmd.Flags().Bool("no-headers", false, "When using the default or custom-column output format, don't print headers (default print headers).")
	getCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=..")
	getCmd.Flags().Bool("output-watch-events", false, "Output watch event objects when --watch or --watch-only is used. Existing objects are output as init")
	getCmd.Flags().String("raw", "", "Raw URI to request from the server.  Uses the transport specified by the kubeconfig file.")
	getCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related man")
	getCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)")
	getCmd.Flags().Bool("server-print", false, "If true, have the server return the appropriate table output. Supports extension APIs and CRDs.")
	getCmd.Flags().Bool("show-kind", false, "If present, list the resource type for the requested object(s).")
	getCmd.Flags().Bool("show-labels", false, "When printing, show all labels as the last column (default hide labels column)")
	getCmd.Flags().String("sort-by", "", "If non-empty, sort list types using this field specification.  The field specification is expressed ")
	getCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The templa")
	getCmd.Flags().Bool("use-openapi-print-columns", false, "If true, use x-kubernetes-print-column metadata (if present) from the OpenAPI schema for displaying ")
	getCmd.Flags().BoolP("watch", "w", false, "After listing/getting the requested object, watch for changes. Uninitialized objects are excluded if")
	getCmd.Flags().Bool("watch-only", false, "Watch for changes to the requested object(s), without listing/getting first.")
	rootCmd.AddCommand(getCmd)

	carapace.Gen(getCmd).FlagCompletion(carapace.ActionMap{
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
		"output":    action.ActionOutputFormats(),
		"template":  carapace.ActionFiles(),
	})

	carapace.Gen(getCmd).PositionalCompletion(
		carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionApiResources().Invoke(c).Filter(c.Parts).ToA()
		}),
	)

	carapace.Gen(getCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionResources("", c.Args[0]).Invoke(c).Filter(c.Args[1:]).ToA()
		}),
	)
}
