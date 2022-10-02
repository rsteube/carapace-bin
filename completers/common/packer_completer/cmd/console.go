package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/packer_completer/cmd/action"
	"github.com/spf13/cobra"
)

var consoleCmd = &cobra.Command{
	Use:   "console",
	Short: "creates a console for testing variable interpolation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(consoleCmd).Standalone()

	consoleCmd.Flags().StringArrayS("var", "var", []string{}, "Variable for templates, can be used multiple times.")
	consoleCmd.Flags().StringS("var-file", "var-file", "", "JSON or HCL2 file containing user variables.")
	rootCmd.AddCommand(consoleCmd)

	carapace.Gen(consoleCmd).FlagCompletion(carapace.ActionMap{
		"var": carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			if len(c.Parts) == 0 && len(c.Args) > 0 {
				return action.ActionVariables(c.Args[0]).Invoke(c).Suffix("=").ToA()
			}
			return carapace.ActionValues()
		}),
		"var-file": carapace.ActionFiles(".json", ".hcl"),
	})

	carapace.Gen(consoleCmd).PositionalCompletion(
		carapace.ActionFiles(".json", ".hcl"),
	)
}
