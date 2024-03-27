package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edits a kustomization file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(editCmd).Standalone()
	rootCmd.AddCommand(editCmd)
}
