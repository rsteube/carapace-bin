package cmd

import (
	"github.com/rsteube/carapace"
	terraform "github.com/rsteube/carapace-bin/completers/terraform_completer/cmd"
	"github.com/spf13/cobra"
)

var runAllCmd = &cobra.Command{
	Use:                "run-all",
	Short:              "Run a terraform command against a 'stack' by running the specified command in each subfolder",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(runAllCmd).Standalone()

	rootCmd.AddCommand(runAllCmd)

	carapace.Gen(runAllCmd).PositionalAnyCompletion(
		terraform.ActionExecute(),
	)
}
