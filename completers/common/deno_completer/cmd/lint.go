package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/deno_completer/cmd/action"
	"github.com/spf13/cobra"
)

var lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "Lint source files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lintCmd).Standalone()

	lintCmd.Flags().StringP("config", "c", "", "Load configuration file")
	lintCmd.Flags().String("ignore", "", "Ignore linting particular source files")
	lintCmd.Flags().Bool("json", false, "Output lint result in JSON format")
	lintCmd.Flags().Bool("rules", false, "List available rules")
	lintCmd.Flags().String("rules-exclude", "", "Exclude lint rules")
	lintCmd.Flags().String("rules-include", "", "Include lint rules")
	lintCmd.Flags().String("rules-tags", "", "Use set of rules with a tag")
	rootCmd.AddCommand(lintCmd)

	lintCmd.Flag("ignore").NoOptDefVal = " "
	lintCmd.Flag("rules-exclude").NoOptDefVal = " "
	lintCmd.Flag("rules-include").NoOptDefVal = " "
	lintCmd.Flag("rules-tags").NoOptDefVal = " "

	carapace.Gen(lintCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
		"rules-exclude": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionLintRules().Invoke(c).Filter(c.Parts).ToA()
		}),
		"rules-include": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionLintRules().Invoke(c).Filter(c.Parts).ToA()
		}),
	})

	carapace.Gen(lintCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
