package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var importEnvironmentCmd = &cobra.Command{
	Use:   "import-environment",
	Short: "Import all or some environment variables",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(importEnvironmentCmd).Standalone()

	rootCmd.AddCommand(importEnvironmentCmd)

	carapace.Gen(importEnvironmentCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return os.ActionEnvironmentVariables().Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
