package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var clusterInfo_dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Dump lots of relevant info for debugging and diagnosis",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clusterInfo_dumpCmd).Standalone()

	clusterInfo_dumpCmd.Flags().BoolP("all-namespaces", "A", false, "If true, dump all namespaces.  If true, --namespaces is ignored.")
	clusterInfo_dumpCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only app")
	clusterInfo_dumpCmd.Flags().String("namespaces", "", "A comma separated list of namespaces to dump.")
	clusterInfo_dumpCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|js")
	clusterInfo_dumpCmd.Flags().String("output-directory", "", "Where to output the files.  If empty or '-' uses stdout, otherwise creates a directory hierarchy in ")
	clusterInfo_dumpCmd.Flags().String("pod-running-timeout", "", "The length of time (like 5s, 2m, or 3h, higher than zero) to wait until at least one pod is running")
	clusterInfo_dumpCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The templa")
	clusterInfoCmd.AddCommand(clusterInfo_dumpCmd)

	carapace.Gen(clusterInfo_dumpCmd).FlagCompletion(carapace.ActionMap{
		"namespaces":       action.ActionResources("", "namespaces"),
		"output":           action.ActionOutputFormats(),
		"output-directory": carapace.ActionDirectories(),
		"template":         carapace.ActionFiles(),
	})
}
