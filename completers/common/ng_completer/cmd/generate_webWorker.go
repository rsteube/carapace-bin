package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/ng_completer/cmd/action"
	"github.com/spf13/cobra"
)

var generate_webWorkerCmd = &cobra.Command{
	Use:   "web-worker",
	Short: "Creates a new, generic web worker definition in the given or default project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_webWorkerCmd).Standalone()

	generate_webWorkerCmd.Flags().String("project", "", "The name of the project.")
	generate_webWorkerCmd.Flags().Bool("snippet", false, "Add a worker creation snippet in a sibling file of the same name.")
	generateCmd.AddCommand(generate_webWorkerCmd)

	carapace.Gen(generate_webWorkerCmd).FlagCompletion(carapace.ActionMap{
		"project": action.ActionProjects(),
	})
}
