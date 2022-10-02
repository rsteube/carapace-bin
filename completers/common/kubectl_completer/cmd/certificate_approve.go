package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var certificate_approveCmd = &cobra.Command{
	Use:   "approve",
	Short: "Approve a certificate signing request",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(certificate_approveCmd).Standalone()

	certificate_approveCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only app")
	certificate_approveCmd.Flags().StringP("filename", "f", "", "Filename, directory, or URL to files identifying the resource to update")
	certificate_approveCmd.Flags().Bool("force", false, "Update the CSR even if it is already approved.")
	certificate_approveCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	certificate_approveCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|js")
	certificate_approveCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related man")
	certificate_approveCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The templa")
	certificateCmd.AddCommand(certificate_approveCmd)

	carapace.Gen(certificate_approveCmd).FlagCompletion(carapace.ActionMap{
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
		"output":    action.ActionOutputFormats(),
		"template":  carapace.ActionFiles(),
	})
}
