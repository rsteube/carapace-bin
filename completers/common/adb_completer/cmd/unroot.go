package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var unrootCmd = &cobra.Command{
	Use:   "unroot",
	Short: "restart adbd without root permissions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unrootCmd).Standalone()

	rootCmd.AddCommand(unrootCmd)
}
