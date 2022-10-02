package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var dllsCmd = &cobra.Command{
	Use:   "dlls",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dllsCmd).Standalone()

	rootCmd.AddCommand(dllsCmd)
}
