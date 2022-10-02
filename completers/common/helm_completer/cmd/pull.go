package cmd

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/helm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "download a chart from a repository and (optionally) unpack it in local directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	pullCmd.Flags().String("ca-file", "", "verify certificates of HTTPS-enabled servers using this CA bundle")
	pullCmd.Flags().String("cert-file", "", "identify HTTPS client using this SSL certificate file")
	pullCmd.Flags().StringP("destination", "d", ".", "location to write the chart. If this and tardir are specified, tardir is appended to this")
	pullCmd.Flags().Bool("devel", false, "use development versions, too. Equivalent to version '>0.0.0-0'. If --version is set, this is ignored.")
	pullCmd.Flags().Bool("insecure-skip-tls-verify", false, "skip tls certificate checks for the chart download")
	pullCmd.Flags().String("key-file", "", "identify HTTPS client using this SSL key file")
	pullCmd.Flags().String("keyring", "/home/user/.gnupg/pubring.gpg", "location of public keys used for verification")
	pullCmd.Flags().Bool("pass-credentials", false, "pass credentials to all domains")
	pullCmd.Flags().String("password", "", "chart repository password where to locate the requested chart")
	pullCmd.Flags().Bool("prov", false, "fetch the provenance file, but don't perform verification")
	pullCmd.Flags().String("repo", "", "chart repository url where to locate the requested chart")
	pullCmd.Flags().Bool("untar", false, "if set to true, will untar the chart after downloading it")
	pullCmd.Flags().String("untardir", ".", "if untar is specified, this flag specifies the name of the directory into which the chart is expanded")
	pullCmd.Flags().String("username", "", "chart repository username where to locate the requested chart")
	pullCmd.Flags().Bool("verify", false, "verify the package before using it")
	pullCmd.Flags().String("version", "", "specify a version constraint for the chart version to use. This constraint can be a specific tag (e.g. 1.1.1) or it may reference a valid range (e.g. ^2.0.0). If this is not specified, the latest version is used")
	rootCmd.AddCommand(pullCmd)

	carapace.Gen(pullCmd).FlagCompletion(carapace.ActionMap{
		"ca-file":     carapace.ActionFiles(),
		"cert-file":   carapace.ActionFiles(),
		"destination": carapace.ActionDirectories(),
		"key-file":    carapace.ActionFiles(),
		"keyring":     carapace.ActionFiles(),
		"untardir":    carapace.ActionDirectories(),
		"version": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				if splitted := strings.Split(c.Args[0], "/"); len(splitted) == 2 {
					return action.ActionChartVersions(splitted[0], splitted[1])
				}
			}
			return carapace.ActionValues()
		}),
	})

	carapace.Gen(pullCmd).PositionalCompletion(
		action.ActionRepositoryCharts(),
	)
}
