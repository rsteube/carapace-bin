package cmd

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/helm_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/util"
	"github.com/spf13/cobra"
)

var show_allCmd = &cobra.Command{
	Use:   "all",
	Short: "show all information of the chart",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	show_allCmd.Flags().String("ca-file", "", "verify certificates of HTTPS-enabled servers using this CA bundle")
	show_allCmd.Flags().String("cert-file", "", "identify HTTPS client using this SSL certificate file")
	show_allCmd.Flags().Bool("devel", false, "use development versions, too. Equivalent to version '>0.0.0-0'. If --version is set, this is ignored")
	show_allCmd.Flags().Bool("insecure-skip-tls-verify", false, "skip tls certificate checks for the chart download")
	show_allCmd.Flags().String("key-file", "", "identify HTTPS client using this SSL key file")
	show_allCmd.Flags().String("keyring", "/home/user/.gnupg/pubring.gpg", "location of public keys used for verification")
	show_allCmd.Flags().Bool("pass-credentials", false, "pass credentials to all domains")
	show_allCmd.Flags().String("password", "", "chart repository password where to locate the requested chart")
	show_allCmd.Flags().String("repo", "", "chart repository url where to locate the requested chart")
	show_allCmd.Flags().String("username", "", "chart repository username where to locate the requested chart")
	show_allCmd.Flags().Bool("verify", false, "verify the package before using it")
	show_allCmd.Flags().String("version", "", "specify a version constraint for the chart version to use. This constraint can be a specific tag (e.g. 1.1.1) or it may reference a valid range (e.g. ^2.0.0). If this is not specified, the latest version is used")
	showCmd.AddCommand(show_allCmd)

	carapace.Gen(show_allCmd).FlagCompletion(carapace.ActionMap{
		"ca-file":   carapace.ActionFiles(),
		"cert-file": carapace.ActionFiles(),
		"key-file":  carapace.ActionFiles(),
		"keyring":   carapace.ActionFiles(),
		"version": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				if splitted := strings.Split(c.Args[0], "/"); len(splitted) == 2 {
					return action.ActionChartVersions(splitted[0], splitted[1])
				}
			}
			return carapace.ActionValues()
		}),
	})

	carapace.Gen(show_allCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if util.HasPathPrefix(c.CallbackValue) {
				return carapace.ActionFiles()
			}
			return action.ActionRepositoryCharts()
		}),
	)
}
