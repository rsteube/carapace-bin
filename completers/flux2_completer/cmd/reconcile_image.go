package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var reconcile_imageCmd = &cobra.Command{
	Use:   "image",
	Short: "Reconcile image automation objects",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reconcile_imageCmd).Standalone()
	reconcileCmd.AddCommand(reconcile_imageCmd)
}
