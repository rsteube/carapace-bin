package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/faas-cli_completer/cmd/action"
	"github.com/spf13/cobra"
)

var template_store_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List templates from OpenFaaS organizations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	template_store_listCmd.Flags().StringP("platform", "p", "x86_64", "Shows the platform if the output is verbose")
	template_store_listCmd.PersistentFlags().StringP("url", "u", "https://raw.githubusercontent.com/openfaas/store/master/templates.json", "Use as alternative store for templates")
	template_store_listCmd.Flags().BoolP("verbose", "v", false, "Shows additional language and platform")
	template_storeCmd.AddCommand(template_store_listCmd)

	carapace.Gen(template_store_listCmd).FlagCompletion(carapace.ActionMap{
		"platform": action.ActionImageArchitectures(),
	})
}
