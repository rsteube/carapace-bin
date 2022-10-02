package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/ng_completer/cmd/action"
	"github.com/spf13/cobra"
)

var generate_serviceWorkerCmd = &cobra.Command{
	Use:   "service-worker",
	Short: "Pass this schematic to the \"run\" command to create a service worker",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_serviceWorkerCmd).Standalone()

	generate_serviceWorkerCmd.Flags().String("project", "", "The name of the project.")
	generate_serviceWorkerCmd.Flags().String("target", "", "The target to apply service worker to.")
	generateCmd.AddCommand(generate_serviceWorkerCmd)

	carapace.Gen(generate_serviceWorkerCmd).FlagCompletion(carapace.ActionMap{
		"project": action.ActionProjects(),
	})
}
