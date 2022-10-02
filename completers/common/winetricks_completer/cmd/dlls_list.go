package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var dlls_listCmd = &cobra.Command{
	Use:   "list",
	Short: "list verbs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dlls_listCmd).Standalone()

	dllsCmd.AddCommand(dlls_listCmd)
}
