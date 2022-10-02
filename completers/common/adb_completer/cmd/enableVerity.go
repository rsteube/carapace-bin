package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var enableVerityCmd = &cobra.Command{
	Use:   "enable-verity",
	Short: "re-enable dm-verity checking on userdebug builds",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(enableVerityCmd).Standalone()

	rootCmd.AddCommand(enableVerityCmd)
}
