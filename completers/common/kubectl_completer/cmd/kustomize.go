package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var kustomizeCmd = &cobra.Command{
	Use:   "kustomize",
	Short: "Build a kustomization target from a directory or a remote url",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kustomizeCmd).Standalone()

	rootCmd.AddCommand(kustomizeCmd)

	carapace.Gen(kustomizeCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
