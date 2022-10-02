package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var disableVerityCmd = &cobra.Command{
	Use:   "disable-verity",
	Short: "disable dm-verity checking on userdebug builds",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(disableVerityCmd).Standalone()

	rootCmd.AddCommand(disableVerityCmd)
}
