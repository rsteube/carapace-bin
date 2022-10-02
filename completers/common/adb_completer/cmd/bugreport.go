package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var bugreportCmd = &cobra.Command{
	Use:   "bugreport",
	Short: "write bugreport go given path",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bugreportCmd).Standalone()

	rootCmd.AddCommand(bugreportCmd)

	carapace.Gen(bugreportCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
