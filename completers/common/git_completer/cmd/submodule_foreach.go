package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var submodule_foreachCmd = &cobra.Command{
	Use:   "foreach",
	Short: "Evaluates an arbitrary shell command in each checked out submodule",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(submodule_foreachCmd).Standalone()
	submodule_foreachCmd.Flags().Bool("recursive", false, "")

	submoduleCmd.AddCommand(submodule_foreachCmd)
}
