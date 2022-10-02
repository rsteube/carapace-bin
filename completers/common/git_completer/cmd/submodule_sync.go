package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var submodule_syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Synchronizes submodules remote URL to .gitmodules",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(submodule_syncCmd).Standalone()
	submodule_syncCmd.Flags().Bool("recursive", false, "")

	submoduleCmd.AddCommand(submodule_syncCmd)

	carapace.Gen(submodule_syncCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
