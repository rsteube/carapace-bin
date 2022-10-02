package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/ng_completer/cmd/action"
	"github.com/spf13/cobra"
)

var extractI18nCmd = &cobra.Command{
	Use:   "extract-i18n",
	Short: "Extracts i18n messages from source code",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(extractI18nCmd).Standalone()

	extractI18nCmd.Flags().String("browser-target", "", "A browser builder target to extract i18n messages")
	extractI18nCmd.Flags().StringP("configuration", "c", "", "One or more named builder configurations")
	extractI18nCmd.Flags().String("format", "", "Output format for the generated file.")
	extractI18nCmd.Flags().String("out-file", "", "Name of the file to output.")
	extractI18nCmd.Flags().String("output-path", "", "Path where output will be placed.")
	extractI18nCmd.Flags().Bool("progress", false, "Log progress to the console.")
	rootCmd.AddCommand(extractI18nCmd)

	carapace.Gen(extractI18nCmd).FlagCompletion(carapace.ActionMap{
		"configuration": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			if len(c.Args) < 1 {
				return carapace.ActionValues()
			}
			return action.ActionBuilderConfigurations(c.Args[0], "").Invoke(c).Filter(c.Parts).ToA()
		}),
		"format":      carapace.ActionValues("xmb", "xlf", "xlif", "xliff", "xlf2", "xliff2", "json", "arb", "legacy-migrate"),
		"out-file":    carapace.ActionFiles(),
		"output-path": carapace.ActionDirectories(),
	})

	carapace.Gen(extractI18nCmd).PositionalCompletion(
		action.ActionProjects(),
	)
}
