package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var edit_fixCmd = &cobra.Command{
	Use:   "fix",
	Short: "Fix the missing fields in kustomization file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(edit_fixCmd).Standalone()
	edit_fixCmd.Flags().Bool("vars", false, "If specified, kustomize will attempt to convert vars to replacements. ")
	editCmd.AddCommand(edit_fixCmd)
}
