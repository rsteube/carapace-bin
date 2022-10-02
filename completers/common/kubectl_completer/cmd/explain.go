package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var explainCmd = &cobra.Command{
	Use:   "explain",
	Short: "Documentation of resources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(explainCmd).Standalone()

	explainCmd.Flags().String("api-version", "", "Get different explanations for particular API version (API group/version)")
	explainCmd.Flags().Bool("recursive", false, "Print the fields of fields (Currently only 1 level deep)")
	rootCmd.AddCommand(explainCmd)

	carapace.Gen(explainCmd).PositionalCompletion(
		action.ActionApiResources(),
	)
}
