package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Dart project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createCmd).Standalone()

	createCmd.Flags().Bool("force", false, "Force project generation, even if the target directory already exists.")
	createCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	createCmd.Flags().Bool("no-pub", false, "Do not run 'pub get' after the project has been created.")
	createCmd.Flags().Bool("pub", false, "Run 'pub get' after the project has been created.")
	createCmd.Flags().StringP("template", "t", "", "The project template to use.")
	rootCmd.AddCommand(createCmd)

	carapace.Gen(createCmd).FlagCompletion(carapace.ActionMap{
		"template": carapace.ActionValuesDescribed(
			"console-simple", "A simple command-line application.",
			"console-full", "A command-line application sample.",
			"package-simple", "A starting point for Dart libraries or applications.",
			"server-simple", "A web server built using package:shelf.",
			"web-simple", "A web app that uses only core Dart libraries.",
		),
	})

	carapace.Gen(createCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
