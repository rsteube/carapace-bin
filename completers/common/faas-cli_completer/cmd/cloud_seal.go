package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/faas-cli_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cloud_sealCmd = &cobra.Command{
	Use:   "seal",
	Short: "Seal a secret for usage with OpenFaaS Cloud",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	cloud_sealCmd.Flags().StringP("cert", "c", "pub-cert.pem", "Filename of public certificate")
	cloud_sealCmd.Flags().Bool("download", false, "Download the kubeseal binary required for this command, see also --download-version")
	cloud_sealCmd.Flags().String("download-to", "", "Specify download path for kubeseal, leave empty for a temp dir")
	cloud_sealCmd.Flags().String("download-version", "", "Specify a kubeseal version to download")
	cloud_sealCmd.Flags().StringArrayP("from-file", "i", []string{}, "Read a secret from a from file")
	cloud_sealCmd.Flags().StringArrayP("literal", "l", []string{}, "Secret literal key-value data")
	cloud_sealCmd.Flags().String("name", "", "Secret name")
	cloud_sealCmd.Flags().StringP("namespace", "n", "openfaas-fn", "Secret name")
	cloud_sealCmd.Flags().StringP("output-file", "o", "secrets.yml", "Output file for secrets")
	cloudCmd.AddCommand(cloud_sealCmd)

	carapace.Gen(cloud_sealCmd).FlagCompletion(carapace.ActionMap{
		"cert":        carapace.ActionFiles(".pem"),
		"download-to": carapace.ActionDirectories(),
		"from-file": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionFiles()
		}),
		"namespace":   action.ActionNamespaces(),
		"output-file": carapace.ActionFiles(),
	})
}
