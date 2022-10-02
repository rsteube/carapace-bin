package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/faas-cli_completer/cmd/action"
	"github.com/spf13/cobra"
)

var template_store_describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Describe the template",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	template_store_describeCmd.PersistentFlags().StringP("url", "u", "https://raw.githubusercontent.com/openfaas/store/master/templates.json", "Use as alternative store for templates")
	template_storeCmd.AddCommand(template_store_describeCmd)

	carapace.Gen(template_store_describeCmd).PositionalCompletion(
		action.ActionLanguageTemplates(),
	)
}
