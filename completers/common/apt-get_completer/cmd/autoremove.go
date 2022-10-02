package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var autoremoveCmd = &cobra.Command{
	Use:   "autoremove",
	Short: "autoremove is used to remove packages that were automatically installed",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoremoveCmd).Standalone()

	rootCmd.AddCommand(autoremoveCmd)
}
