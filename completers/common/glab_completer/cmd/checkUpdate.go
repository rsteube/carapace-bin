package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var checkUpdateCmd = &cobra.Command{
	Use:   "check-update",
	Short: "Check for latest glab releases",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(checkUpdateCmd).Standalone()
	rootCmd.AddCommand(checkUpdateCmd)
}
