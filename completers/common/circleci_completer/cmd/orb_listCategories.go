package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var orb_listCategoriesCmd = &cobra.Command{
	Use:   "list-categories",
	Short: "List orb categories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(orb_listCategoriesCmd).Standalone()
	orb_listCategoriesCmd.PersistentFlags().Bool("json", false, "print output as json instead of human-readable")
	orbCmd.AddCommand(orb_listCategoriesCmd)
}
