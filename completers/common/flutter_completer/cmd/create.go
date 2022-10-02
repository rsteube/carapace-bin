package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/flutter_completer/cmd/action"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Flutter project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createCmd).Standalone()

	createCmd.Flags().StringP("android-language", "a", "", "The language to use for Android-specific code.")
	createCmd.Flags().String("description", "", "The description to use for your new Flutter project.")
	createCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	createCmd.Flags().StringP("ios-language", "i", "", "The language to use for iOS-specific code.")
	createCmd.Flags().String("list-samples", "", "Specifies a JSON output file for a listing of Flutter code samples that can be created with \"--sample\".")
	createCmd.Flags().Bool("no-offline", false, "When \"flutter pub get\" is run by the create command, does not run it in offline mode.")
	createCmd.Flags().Bool("no-overwrite", false, "When performing operations, do not overwrite existing files.")
	createCmd.Flags().Bool("no-pub", false, "Do not run \"flutter pub get\" after the project has been created.")
	createCmd.Flags().Bool("offline", false, "When \"flutter pub get\" is run by the create command, this runs it in offline mode.")
	createCmd.Flags().String("org", "", "The organization responsible for your new Flutter project, in reverse domain name notation.")
	createCmd.Flags().Bool("overwrite", false, "When performing operations, overwrite existing files.")
	createCmd.Flags().String("platforms", "", "The platforms supported by this project.")
	createCmd.Flags().String("project-name", "", "The project name for this new Flutter project.")
	createCmd.Flags().Bool("pub", false, "Run \"flutter pub get\" after the project has been created.")
	createCmd.Flags().StringP("sample", "s", "", "Specifies the Flutter code sample to use as the \"main.dart\" for an application.")
	createCmd.Flags().StringP("template", "t", "", "Specify the type of project to create.")
	rootCmd.AddCommand(createCmd)

	carapace.Gen(createCmd).FlagCompletion(carapace.ActionMap{
		"android-language": carapace.ActionValues("java", "kotlin"),
		"ios-language":     carapace.ActionValues("objc", "swift"),
		"platforms": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ios", "android", "windows", "linux", "macos", "web").Invoke(c).Filter(c.Parts).ToA()
		}),
		"sample": action.ActionSamples(),
		"template": carapace.ActionValuesDescribed(
			"app", "Generate a Flutter application.",
			"module", "Generate a project to add a Flutter module to an existing Android or iOS application.",
			"package", "Generate a shareable Flutter project containing modular Dart code.",
			"plugin", "Generate a shareable Flutter project containing an API in Dart code.",
			"skeleton", "Generate a List View / Detail View Flutter application that follows community best practices.",
		),
	})

	carapace.Gen(createCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
