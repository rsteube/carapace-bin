package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Initiate a Postman Collection run from a given URL or pat",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().String("bail", "", "Specify whether or not to gracefully")
	runCmd.Flags().String("color", "", "Enable/Disable colored output")
	runCmd.Flags().String("cookie-jar", "", "Specify the path to a custom cookie jar")
	runCmd.Flags().String("delay-request", "", "Specify the extent of delay between")
	runCmd.Flags().Bool("disable-unicode", false, "Forces Unicode compliant symbols to be")
	runCmd.Flags().String("env-var", "", "Allows the specification of environment")
	runCmd.Flags().StringP("environment", "e", "", "Specify a URL or path to a Postman")
	runCmd.Flags().String("export-collection", "", "Exports the executed collection to a")
	runCmd.Flags().String("export-cookie-jar", "", "Exports the cookie jar to a file after")
	runCmd.Flags().String("export-environment", "", "Exports the final environment to a file")
	runCmd.Flags().String("export-globals", "", "Exports the final globals to a file")
	runCmd.Flags().String("folder", "", "Specify the folder to run from a")
	runCmd.Flags().String("global-var", "", "Allows the specification of global")
	runCmd.Flags().StringP("globals", "g", "", "Specify a URL or path to a file")
	runCmd.Flags().BoolP("help", "h", false, "display help for command")
	runCmd.Flags().Bool("ignore-redirects", false, "Prevents Newman from automatically")
	runCmd.Flags().BoolP("insecure", "k", false, "Disables SSL validations")
	runCmd.Flags().StringP("iteration-count", "n", "", "Define the number of iterations to run")
	runCmd.Flags().StringP("iteration-data", "d", "", "Specify a data file to use for")
	runCmd.Flags().Bool("no-insecure-file-read", false, "Prevents reading the files situated")
	runCmd.Flags().String("postman-api-key", "", "API Key used to load the resources from")
	runCmd.Flags().StringP("reporters", "r", "", "Specify the reporters to use for this")
	runCmd.Flags().Bool("silent", false, "Prevents Newman from showing output to")
	runCmd.Flags().String("ssl-client-cert", "", "Specify the path to a client")
	runCmd.Flags().String("ssl-client-cert-list", "", "Specify the path to a client")
	runCmd.Flags().String("ssl-client-key", "", "Specify the path to a client")
	runCmd.Flags().String("ssl-client-passphrase", "", "Specify the client certificate")
	runCmd.Flags().String("ssl-extra-ca-certs", "", "Specify additionally trusted CA")
	runCmd.Flags().BoolP("suppress-exit-code", "x", false, "Specify whether or not to override the")
	runCmd.Flags().String("timeout", "", "Specify a timeout for collection run")
	runCmd.Flags().String("timeout-request", "", "Specify a timeout for requests")
	runCmd.Flags().String("timeout-script", "", "Specify a timeout for scripts")
	runCmd.Flags().Bool("verbose", false, "Show detailed information of collection")
	runCmd.Flags().String("working-dir", "", "Specify the path to the working")
	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
		"cookie-jar":         carapace.ActionFiles(),
		"environment":        carapace.ActionFiles(),
		"export-cookie-jar":  carapace.ActionFiles(),
		"export-environment": carapace.ActionFiles(),
		"globals":            carapace.ActionFiles(),
		"iteration-data":     carapace.ActionFiles(),
		"reporters": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("cli", "json", "junit", "progress", "emojitrain").Invoke(c).Filter(c.Parts).ToA()
		}),
		"ssl-client-cert":      carapace.ActionFiles(),
		"ssl-client-cert-list": carapace.ActionFiles(),
		"ssl-client-key":       carapace.ActionFiles(),
		"ssl-extra-ca-certs":   carapace.ActionFiles(),
		"working-dir":          carapace.ActionDirectories(),
	})

	carapace.Gen(runCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
