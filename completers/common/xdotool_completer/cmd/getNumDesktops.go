package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var getNumDesktopsCmd = &cobra.Command{
	Use:   "get_num_desktops",
	Short: "Outputs the current number of desktops",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getNumDesktopsCmd).Standalone()

	rootCmd.AddCommand(getNumDesktopsCmd)
}
