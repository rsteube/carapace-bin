package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var generate_appShellCmd = &cobra.Command{
	Use:   "app-shell",
	Short: "Generates an app shell for running a server-side version of an app",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_appShellCmd).Standalone()

	generate_appShellCmd.Flags().String("app-dir", "", "The name of the application directory.")
	generate_appShellCmd.Flags().String("app-id", "", "The app ID to use in withServerTransition().")
	generate_appShellCmd.Flags().String("main", "", "The name of the main entry-point file.")
	generate_appShellCmd.Flags().String("project", "", "The name of the related client app.")
	generate_appShellCmd.Flags().String("root-module-class-name", "", "The name of the root module class.")
	generate_appShellCmd.Flags().String("root-module-file-name", "", "The name of the root module file")
	generate_appShellCmd.Flags().String("route", "", "Route path used to produce the app shell.")
	generateCmd.AddCommand(generate_appShellCmd)
}
