package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var stack_importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import a deployment from standard in into an existing stack",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stack_importCmd).Standalone()
	stack_importCmd.PersistentFlags().String("file", "", "A filename to read stack input from")
	stack_importCmd.PersistentFlags().BoolP("force", "f", false, "Force the import to occur, even if apparent errors are discovered beforehand (not recommended)")
	stackCmd.AddCommand(stack_importCmd)

	carapace.Gen(stack_importCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})
}
