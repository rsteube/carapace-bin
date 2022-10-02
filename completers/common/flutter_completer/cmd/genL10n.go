package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var genL10nCmd = &cobra.Command{
	Use:   "gen-l10n",
	Short: "Generate localizations for the current project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(genL10nCmd).Standalone()

	genL10nCmd.Flags().String("arb-dir", "", "The directory where the template and translated arb files are located.")
	genL10nCmd.Flags().String("gen-inputs-and-outputs-list", "", "When specified, the tool generates a JSON file containing the tool's inputs and outputs named gen_l10n_inputs_and_outputs.json.")
	genL10nCmd.Flags().String("header", "", "The header to prepend to the generated Dart localizations files.")
	genL10nCmd.Flags().String("header-file", "", "The header to prepend to the generated Dart localizations files.")
	genL10nCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	genL10nCmd.Flags().Bool("no-nullable-getter", false, "Localizations class getter is not nullable.")
	genL10nCmd.Flags().Bool("no-required-resource-attributes", false, "Does not equire all resource ids to contain a corresponding resource attribute.")
	genL10nCmd.Flags().Bool("no-synthetic-package", false, "Disable generated output files. ")
	genL10nCmd.Flags().Bool("no-use-deferred-loading", false, "Do not generate the Dart localization file with locales imported as deferred.")
	genL10nCmd.Flags().Bool("nullable-getter", false, "Localizations class getter is nullable.")
	genL10nCmd.Flags().String("output-class", "", "The Dart class name to use for the output localization and localizations delegate classes.")
	genL10nCmd.Flags().String("output-dir", "", "The directory where the generated localization classes will be written if the synthetic-package flag is set to false.")
	genL10nCmd.Flags().String("output-localization-file", "", "The filename for the output localization and localizations delegate classes.")
	genL10nCmd.Flags().String("preferred-supported-locales", "", "The list of preferred supported locales for the application.")
	genL10nCmd.Flags().String("project-dir", "", "Directory of the root Flutter project.")
	genL10nCmd.Flags().Bool("required-resource-attributes", false, "Requires all resource ids to contain a corresponding resource attribute.")
	genL10nCmd.Flags().Bool("synthetic-package", false, "Enable generated output files.")
	genL10nCmd.Flags().String("template-arb-file", "", "The template arb file that will be used as the basis for generating the Dart localization and messages files.")
	genL10nCmd.Flags().String("untranslated-messages-file", "", "The location of a file that describes the localization messages have not been translated yet.")
	genL10nCmd.Flags().Bool("use-deferred-loading", false, "Generate the Dart localization file with locales imported as deferred.")
	rootCmd.AddCommand(genL10nCmd)
}
