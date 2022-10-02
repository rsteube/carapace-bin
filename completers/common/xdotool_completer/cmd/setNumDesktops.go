package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var setNumDesktopsCmd = &cobra.Command{
	Use:   "set_num_desktops",
	Short: "Changes the number of desktops or workspaces",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setNumDesktopsCmd).Standalone()

	rootCmd.AddCommand(setNumDesktopsCmd)
}
