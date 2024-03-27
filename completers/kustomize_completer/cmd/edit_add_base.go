package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var edit_add_baseCmd = &cobra.Command{
	Use:   "base",
	Short: "Adds one or more bases to the kustomization.yaml in current directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(edit_add_baseCmd).Standalone()
	edit_addCmd.AddCommand(edit_add_baseCmd)
}
